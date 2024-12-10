package hnet

import (
	"net/http"
	"time"

	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/yap"
)

type CORS struct {
	Origin      string
	Methods     string
	Headers     string
	Credentials string
}

func (c CORS) Handler(next http.Handler) http.Handler {
	c.Origin = Or(c.Origin, "*")
	c.Methods = Or(c.Methods, "GET, POST, PUT, DELETE, OPTIONS")
	c.Headers = Or(c.Headers, "Content-Type, Authorization")
	c.Credentials = Or(c.Credentials, "true")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Origin.Set(w, c.Origin)
		Allow_Methods.Set(w, c.Methods)
		Allow_Headers.Set(w, c.Headers)
		Allow_Credentials.Set(w, c.Credentials)

		next.ServeHTTP(w, r)
	})
}

type LogHandler struct {
	http.ResponseWriter
	r     *http.Request
	start time.Time
}

func (h *LogHandler) WriteHeader(code int) {
	h.ResponseWriter.WriteHeader(code)
	if code >= 400 {
		yap.Error("request", "method", h.r.Method, "URL", h.r.RequestURI, "time", time.Since(h.start), "status", code)
	} else {
		yap.Info("request", "method", h.r.Method, "URL", h.r.RequestURI, "time", time.Since(h.start), "status", code)
	}
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(&LogHandler{w, r, start}, r)
	})
}

func Pre(pref string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.StripPrefix(pref, http.HandlerFunc(next)).ServeHTTP
	}
}
