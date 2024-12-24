package blob

import (
	"io"
	"net/http"
	"os"

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

func NewService(root string) (*Service, error) {
	root, err := fsio.AbsPath(root)
	if err != nil {
		return Err[*Service](err)
	}
	return Ok(&Service{
		Blobs: &maps.Sync[string, Blob]{},
		Root:  root,
	})
}

type Blob struct {
	Type ft.Type
	Name string
	Path string
}

func (b Blob) Open() (*os.File, error) { return os.Open(b.Path) }
func (b Blob) delete() error           { return os.Remove(b.Path) }

func (s Service) Set(bucket, blob string, r io.Reader, ct ft.Type) (Blob, error) {
	err := os.MkdirAll(fsio.Join(s.Root, bucket), fsio.DirReadWrite)
	fp := fsio.Join(s.Root, bucket, blob) + ct.Ext()
	name := bucket + "/" + blob
	if err != nil {
		return Err[Blob](err)
	}
	if len(ct.ContentHeader()) == 0 {
		return Err[Blob]("the content type for blob [{:s}] is not valid", name)
	}
	err = fsio.WriteAll(fp, r)
	if err != nil {
		return Err[Blob](err)
	}
	return Ok(Blob{ct, name, fp})
}

func (s Service) Get(bucket, blob string) Option[Blob] { return s.Blobs.Get(bucket + "/" + blob) }
func (s Service) Del(bucket, blob string) error {
	name := bucket + "/" + blob
	blobO := s.Blobs.Get(name)
	if !blobO.Ok {
		return StrErr("couldn't find such blob")
	}
	ok := s.Blobs.Del(name)
	if !ok {
		return StrErr(Format("couldn't delet blob [{:s}]", name))
	}
	return blobO.Value.delete()
}

func Server(srv *Service) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{bucket}/{blob}", func(w http.ResponseWriter, r *http.Request) {
		blobO := srv.Get(r.PathValue("bucket"), r.PathValue("blob"))
		if !blobO.Ok {
			hnet.Not_Found.AsError(w)
			return
		}
		http.ServeFile(w, r, blobO.Value.Path)
	})

	mux.HandleFunc("POST /{bucket}/{blob}", func(w http.ResponseWriter, r *http.Request) {
		ct := ft.FromContentHeader(r.Header.Get("Content-Type"))
		if !ct.Ok {
			hnet.Bad_Request.AsErrorf(w, "the request is missing a content header")
			return
		}
		blob, err := srv.Set(r.PathValue("bucket"), r.PathValue("blob"), r.Body, ct.Value)
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
