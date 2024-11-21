package hnet

import (
	"net/http"
	"time"

	"github.com/periaate/blume/clog"
	"github.com/periaate/blume/gen"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		clog.Info("request", "method", r.Method, "URL", r.RequestURI, "time", time.Since(start))
	})
}

const (
	ACA_ORIGIN      = "Access-Control-Allow-Origin"
	ACA_METHODS     = "Access-Control-Allow-Methods"
	ACA_HEADERS     = "Access-Control-Allow-Headers"
	ACA_CREDENTIALS = "Access-Control-Allow-Credentials"

	DEF_ACA_ORIGIN      = "*"
	DEF_ACA_METHODS     = "GET, POST, PUT, DELETE, OPTIONS"
	DEF_ACA_HEADERS     = "Content-Type, Authorization"
	DEF_ACA_CREDENTIALS = "true"
)

type CORS struct {
	ACA_Origin      string
	ACA_Methods     string
	ACA_Headers     string
	ACA_Credentials string
}

func (c CORS) Handler(next http.Handler) http.Handler {
	acao := gen.Or(c.ACA_Origin, DEF_ACA_ORIGIN)
	acam := gen.Or(c.ACA_Methods, DEF_ACA_METHODS)
	acah := gen.Or(c.ACA_Headers, DEF_ACA_HEADERS)
	acac := gen.Or(c.ACA_Credentials, DEF_ACA_CREDENTIALS)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ACA_ORIGIN, acao)
		w.Header().Set(ACA_ORIGIN, acam)
		w.Header().Set(ACA_ORIGIN, acah)
		w.Header().Set(ACA_ORIGIN, acac)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Pre(pref string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.StripPrefix(pref, http.HandlerFunc(next)).ServeHTTP
	}
}
