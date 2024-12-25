package blob

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/fsio/ft"
	"github.com/periaate/blume/hnet"
	"github.com/periaate/blume/maps"
)

type Service struct {
	Blobs *maps.Sync[string, Blob] // name -> Blob
	Root  string                   // root must be an absolute path.
}

func NewService(root string) (res *Service, err error) {
	root, err = filepath.Abs(root)
	if err != nil {
		return
	}
	res = &Service{
		Blobs: &maps.Sync[string, Blob]{},
		Root:  root,
	}
	return res, err
}

type Blob struct {
	Type ft.Type
	Name string
	Path string
}

func (b Blob) Open() (*os.File, error) { return os.Open(b.Path) }
func (b Blob) delete() error           { return os.Remove(b.Path) }

func (s Service) Set(bucket, blob string, r io.Reader, ct ft.Type) (res Blob, err error) {
	err = os.MkdirAll(path.Join(s.Root, bucket), 0o777)
	if err != nil {
		return
	}
	fp := path.Join(s.Root, bucket, blob) + ct.Ext()
	name := bucket + "/" + blob
	if len(ct.ContentHeader()) == 0 {
		err = fmt.Errorf("the content type for blob [%s] is not valid", name)
		return
	}
	err = fsio.WriteTo(fp, r)
	if err != nil {
		return
	}
	res = Blob{ct, name, fp}
	return
}

func (s Service) Get(bucket, name string) (Blob, bool) { return s.Blobs.Get(bucket + "/" + name) }
func (s Service) Del(bucket, name string) error {
	name = bucket + "/" + name
	blob, ok := s.Blobs.Get(name)
	if !ok {
		return StrErr("couldn't find such blob")
	}
	ok = s.Blobs.Del(name)
	if !ok {
		return fmt.Errorf("couldn't delet blob [%s]", name)
	}
	return blob.delete()
}

func Server(srv *Service) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{bucket}/{blob}", func(w http.ResponseWriter, r *http.Request) {
		blob, ok := srv.Get(r.PathValue("bucket"), r.PathValue("blob"))
		if !ok {
			hnet.Not_Found.AsError(w)
			return
		}
		http.ServeFile(w, r, blob.Path)
	})

	mux.HandleFunc("POST /{bucket}/{blob}", func(w http.ResponseWriter, r *http.Request) {
		ct, ok := ft.FromContentHeader(r.Header.Get("Content-Type"))
		if !ok {
			hnet.Bad_Request.AsErrorf(w, "the request is missing a content header")
			return
		}
		blob, err := srv.Set(r.PathValue("bucket"), r.PathValue("blob"), r.Body, ct)
		if err != nil {
			hnet.Not_Found.AsErrorf(w, "%s", err)
			return
		}
		http.ServeFile(w, r, blob.Path)
	})

	mux.HandleFunc("DELETE /{bucket}/{blob}", func(w http.ResponseWriter, r *http.Request) {
		err := srv.Del(r.PathValue("bucket"), r.PathValue("blob"))
		if err != nil {
			hnet.Internal_Server_Error.AsErrorf(w, "error deleting blob %s", err)
			return
		}
		w.WriteHeader(int(hnet.OK))
	})

	return mux
}
