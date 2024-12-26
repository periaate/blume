package blob

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/fsio/ft"
	"github.com/periaate/blume/types/maps"
)

func New(root string, validator func(string, string) bool) (res *Store, err error) {
	root, err = filepath.Abs(root)
	if err == nil {
		res = &Store{maps.New(validator), root}
	}
	return
}

type Store struct {
	maps.Map[string, string]
	root string
}

func (s Store) Set(bucket, blob string, r io.Reader, ct ft.Type) (res string, err error) {
	err = os.MkdirAll(path.Join(s.root, bucket), 0o777)
	if err != nil {
		return
	}
	res = path.Join(s.root, bucket, blob) + ct.Ext()
	err = fsio.WriteAll(res, r)
	return
}

func (s Store) Get(bucket, name string) (val string, err error) {
	val, ok, isValid := s.GetFull(bucket + "/" + name)
	if !ok {
		return "", fmt.Errorf("couldn't find such blob")
	}
	if !isValid {
		return "", s.Del(bucket, name)
	}
	return
}

func (s Store) Del(bucket, name string) error {
	name = bucket + "/" + name
	blob, ok := s.Map.Get(name)
	if !ok {
		return fmt.Errorf("couldn't find such blob")
	}
	s.Map.Del(name)
	return os.Remove(blob)
}

func Server(srv *Store) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{bucket}/{blob}", func(w http.ResponseWriter, r *http.Request) {
		blob, err := srv.Get(r.PathValue("bucket"), r.PathValue("blob"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, blob)
	})
	mux.HandleFunc("POST /{bucket}/{blob}", func(w http.ResponseWriter, r *http.Request) {
		ct, ok := ft.FromContentHeader(r.Header.Get("Content-Type"))
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		blob, err := srv.Set(r.PathValue("bucket"), r.PathValue("blob"), r.Body, ct)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, blob)
	})
	mux.HandleFunc("DELETE /{bucket}/{blob}", func(w http.ResponseWriter, r *http.Request) {
		err := srv.Del(r.PathValue("bucket"), r.PathValue("blob"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
