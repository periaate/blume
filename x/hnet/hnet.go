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
	ACCESSS_CONTROL_ALLOW_ORIGIN      = "Access-Control-Allow-Origin"
	ACCESSS_CONTROL_ALLOW_METHODS     = "Access-Control-Allow-Methods"
	ACCESSS_CONTROL_ALLOW_HEADERS     = "Access-Control-Allow-Headers"
	ACCESSS_CONTROL_ALLOW_CREDENTIALS = "Access-Control-Allow-Credentials"

	DEFAULT_ACCESS_CONTROL_ALLOW_ORIGIN      = "*"
	DEFAULT_ACCESS_CONTROL_ALLOW_METHODS     = "GET, POST, PUT, DELETE, OPTIONS"
	DEFAULT_ACCESS_CONTROL_ALLOW_HEADERS     = "Content-Type, Authorization"
	DEFAULT_ACCESS_CONTROL_ALLOW_CREDENTIALS = "true"
)

type CORS struct {
	ACCESS_CONTROL_ALLOW_ORIGIN      string
	ACCESS_CONTROL_ALLOW_METHODS     string
	ACCESS_CONTROL_ALLOW_HEADERS     string
	ACCESS_CONTROL_ALLOW_CREDENTIALS string
}

func (c CORS) Handler(next http.Handler) http.Handler {
	ACCESS_CONTROL_ALLOW_ORIGIN_VALUE := gen.Or(c.ACCESS_CONTROL_ALLOW_ORIGIN, DEFAULT_ACCESS_CONTROL_ALLOW_ORIGIN)
	ACCESS_CONTROL_ALLOW_METHODS_VALUE := gen.Or(c.ACCESS_CONTROL_ALLOW_METHODS, DEFAULT_ACCESS_CONTROL_ALLOW_METHODS)
	ACCESS_CONTROL_ALLOW_HEADERS_VALUE := gen.Or(c.ACCESS_CONTROL_ALLOW_HEADERS, DEFAULT_ACCESS_CONTROL_ALLOW_HEADERS)
	ACCESS_CONTROL_ALLOW_CREDENTIALS_VALUE := gen.Or(c.ACCESS_CONTROL_ALLOW_CREDENTIALS, DEFAULT_ACCESS_CONTROL_ALLOW_CREDENTIALS)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ACCESSS_CONTROL_ALLOW_ORIGIN, ACCESS_CONTROL_ALLOW_ORIGIN_VALUE)
		w.Header().Set(ACCESSS_CONTROL_ALLOW_METHODS, ACCESS_CONTROL_ALLOW_METHODS_VALUE)
		w.Header().Set(ACCESSS_CONTROL_ALLOW_HEADERS, ACCESS_CONTROL_ALLOW_HEADERS_VALUE)
		w.Header().Set(ACCESSS_CONTROL_ALLOW_CREDENTIALS, ACCESS_CONTROL_ALLOW_CREDENTIALS_VALUE)

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
