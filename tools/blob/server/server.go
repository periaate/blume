package server

import (
	"io"
	"net/http"

	"github.com/periaate/blob"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/hnet"
	"github.com/periaate/blume/util"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{bucket}/{name}", Get)
	mux.HandleFunc("POST /{bucket}/{name}", Set)
	mux.HandleFunc("DELETE /{bucket}/{name}", Del)

	return mux
}

func GetContentType(r *http.Request) (ct blob.ContentType, nerr hnet.NetErr) {
	cth := r.Header.Get("Content-Type")
	if cth == "" {
		nerr = hnet.BadRequest.Def("missing Content-Type header")
		return
	}

	ct = blob.GetCT(cth)
	return
}

func PathValues(r *http.Request) (bucket, name string, nerr hnet.NetErr) {
	p := hnet.PathValue(r)
	bucket = p.String("bucket", util.NotZero[string]())
	name = p.String("name", util.NotZero[string]())
	for _, v := range p.Nerrors {
		nerr = v
		return
	}
	return
}

func Get(w http.ResponseWriter, r *http.Request) {
	bucket, name, nerr := PathValues(r)
	if nerr != nil {
		nerr.Respond(w)
		return
	}

	reader, ct, nerr := blob.Blob(fsio.Join(bucket, name)).Get()
	if nerr != nil {
		nerr.Respond(w)
		return
	}

	w.Header().Set("Content-Type", ct.String())
	_, err := io.Copy(w, reader)
	if err != nil {
		http.Error(w, "couldn't write blob to response", http.StatusInternalServerError)
	}
}

func Set(w http.ResponseWriter, r *http.Request) {
	bucket, name, nerr := PathValues(r)
	if nerr != nil {
		nerr.Respond(w)
		return
	}

	ct, nerr := GetContentType(r)
	if nerr != nil {
		nerr.Respond(w)
		return
	}

	nerr = blob.Blob(fsio.Join(bucket, name)).Set(r.Body, ct)
	if nerr != nil {
		nerr.Respond(w)
	}
}

func Del(w http.ResponseWriter, r *http.Request) {
	bucket, name, nerr := PathValues(r)
	if nerr != nil {
		nerr.Respond(w)
		return
	}

	nerr = blob.Blob(fsio.Join(bucket, name)).Del()
	if nerr != nil {
		nerr.Respond(w)
	}
}
