package blob

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/hnet"
	"github.com/periaate/blume/maps"
	"github.com/periaate/blume/yap"
)

/*
Type: {Content-Type}
Blobpath: {bucket}/{blob}
Blobname: {Type}{blob}
Filepath: {root}/{bucket}/{Blobname}


Type is two base64 letters.

Content-Types are provided during a Set operation on the Index.

*/

func SetIndex(root string, opts ...func(*Index)) {
	fsio.EnsureDir(root)
	I = &Index{
		Sync: maps.NewSync[Blob, ContentType](),
		Root: root,
	}

	for _, opt := range opts {
		opt(I)
	}

	// Load all blobs from the root directory.
	filepaths, _ := fsio.ReadDir(root)
	for _, filepath := range filepaths {
		fmt.Println(filepath)
		if !fsio.IsDir(filepath) {
			yap.Error("bucket is not a directory", "blob", filepath)
			continue
		}

		bucket := fsio.Base(filepath)
		blobpaths, _ := fsio.ReadDir(filepath)
		for _, blobpath := range blobpaths {
			fmt.Println(blobpath)
			if fsio.IsDir(blobpath) {
				yap.Error("blob is a directory", "blob", blobpath)
				continue
			}

			base := fsio.Base(blobpath)
			ct := GetCT(base[:2])
			if ct == -1 {
				yap.Error("invalid content type", "blob", blobpath)
				continue
			}

			name := base[2:]

			blob := Blob(fsio.Join(bucket, name))
			I.Set(blob, ct)
			if _, ok := I.Get(blob); !ok {
				yap.Error("blob not set in index", "blob", blob)
			}
		}
	}
}

func CloseIndex() {
	I = nil
}

var I *Index

type Index struct {
	*maps.Sync[Blob, ContentType]
	Root string
	mut  sync.RWMutex
}

type Blob string

func (b Blob) Split() (bucket, blob string, nerr hnet.NetErr) {
	sar := gen.SplitWithAll(string(b), false, "/")
	if len(sar) != 2 {
		nerr = hnet.Bad_Request.Def(
			"tried to split", "Blob",
			"with bad format", string(b),
			"blob format is", "{bucket}/{blob}",
		)
		return
	}

	bucket = sar[0]
	blob = sar[1]
	bucket = gen.Replace(":", "_")(bucket)
	return
}

func Filepath(b Blob, ct ContentType) (res string, nerr hnet.NetErr) {
	bucket, blob, nerr := b.Split()
	if nerr != nil {
		return
	}
	bucket = gen.Replace(":", "_")(bucket)
	res = fsio.Join(I.Root, bucket, ct.Fmt()+blob)
	return
}

func (b Blob) File() (res string, ct ContentType, nerr hnet.NetErr) {
	ct, nerr = b.Type()
	if nerr != nil {
		return
	}

	bucket, blob, nerr := b.Split()
	if nerr != nil {
		return
	}

	res = fsio.Join(I.Root, bucket, ct.Fmt()+blob)
	return
}

func (b Blob) Type() (res ContentType, nerr hnet.NetErr) {
	res, ok := I.Get(b)
	if !ok {
		nerr = hnet.Not_Found.Def(
			"Blob", "Index", string(b),
			"msg", "Blob not found in index",
		)
	}
	return
}

// Set attempts to set this blob.
func (b Blob) Set(r io.Reader, ct ContentType) (nerr hnet.NetErr) {
	fmt.Println("A")
	I.mut.Lock()
	fmt.Println("B")
	defer I.mut.Unlock()
	fmt.Println("C")
	file, nerr := Filepath(b, ct)
	if nerr != nil {
		return
	}

	bucket, _, nerr := b.Split()
	if nerr != nil {
		return
	}

	err := fsio.EnsureDir(fsio.Join(I.Root, bucket))
	if err != nil {
		nerr = hnet.Internal_Server_Error.Def(
			"tried to set", "Blob",
			"with value", string(b),
			"to", "Bucket",
			"failed to ensure buckets existence:", err.Error(),
		)

		return
	}

	fmt.Println("D")
	ok := I.Set(b, ct)
	fmt.Println("E")
	if !ok {
		nerr = hnet.Conflict.Def(
			"tried to set", "Blob",
			"with value", string(b),
			"to", "Index",
			"because", "blob already exists",
		)
		return
	}

	fmt.Println("F")
	err = fsio.WriteAll(file, r)
	if nerr != nil {
		return
	}

	return
}

// Get attempts to get this blob.
func (b Blob) Get() (r io.Reader, ct ContentType, nerr hnet.NetErr) {
	I.mut.RLock()
	defer I.mut.RUnlock()
	file, ct, nerr := b.File()
	if nerr != nil {
		return
	}

	bar, err := os.ReadFile(file)
	if err != nil {
		nerr = hnet.Not_Found.Def("Blob", string(b), "Bucket")
		return
	}

	r = bytes.NewBuffer(bar)
	return
}

// Del attempts to remove this blob.
func (b Blob) Del() (nerr hnet.NetErr) {
	I.mut.Lock()
	defer I.mut.Unlock()
	file, _, nerr := b.File()
	if nerr != nil {
		return
	}

	err := os.Remove(file)
	if err != nil {
		nerr = hnet.Internal_Server_Error.Def(
			"tried to delete", "Blob",
			"with value", string(b),
			"from", "Bucket",
			"failed to remove file:", err.Error(),
		)
		return
	}

	if !I.Del(b) {
		nerr = hnet.Conflict.Def(
			"encountered impossible error while deleting", "Blob",
			"with value", string(b),
			"from", "Index",
			"because", "blob was removed from filesystem before index",
		)
	}

	return
}
