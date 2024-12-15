package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"auth"

	"github.com/periaate/blob"
	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
	. "github.com/periaate/blume/hnet"
	"github.com/periaate/blume/yap"
)

// TODO: validate link and session hosts
// TODO: add configurable logging

func FFrom[A any](r []byte) (a A, err error) {
	err = json.Unmarshal(r, &a)
	return
}

func main() {
	args := fsio.Args[String]().Unwrap()
	inward := args.Shift().Unwrap()
	outward := args.Shift().Unwrap()

	man := auth.NewManager()
	del := []blob.Blob{}
	blob.SetIndex(fsio.Getenv("BLOB_INDEX").Or("./data"))

	for b := range blob.I.Iter() {
		r, _, _ := b.Get()
		bar, err := io.ReadAll(r)
		if err != nil { continue }
		s, err := FFrom[auth.Session](bar)
		if err != nil {
			yap.Error("error decoding session", "err", err)
			continue
		}

		ok := man.Register(s)
		if !ok {
			yap.Error("error registering session", "session", s, "blob", b)
			del = append(del, b)
			continue
		}
		yap.Info("registering cookie", "cookie", s.Cookie, "until", time.Until(s.T).String())
	}

	for _, b := range del {
		b.Del()
	}

	go func() {
		http.HandleFunc("GET /gen/{host}/{label}/{duration}/{uses}", func(w http.ResponseWriter, r *http.Request) {
			duration := time.Duration(ToInt(r.PathValue("duration")).Or(0)) * time.Minute
			uses := ToInt(r.PathValue("uses")).Or(0)
			host := r.PathValue("host")
			label := r.PathValue("label")

			if LT(time.Minute)(duration) || LT(1)(uses) || host == "" {
				yap.Error("incorrect input values generating")
				Bad_Request.AsError(w)
				return
			}

			link, _ := man.NewLink(uses, label, host, duration)
			res := fsio.Join(HTTP.Use(host), link)
			fmt.Fprintf(w, "%s", res)
		})
		yap.Info("starting inward server", "addr", URL(inward))
		http.ListenAndServe(inward.String(), nil)
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		sessKey, _ := r.Cookie("X-Session")
		if sessKey == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		opt := man.Sessions.Get(sessKey.Value)
		if !opt.Ok() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /fw-auth/{host}/{hash...}", func(w http.ResponseWriter, r *http.Request) {
		bucket := r.PathValue("host")
		hash := r.PathValue("hash")

		b := Split(bucket, false, ".")
		if len(b) >= 3 { bucket = b[len(b)-2] + "." + b[len(b)-1] }
		yap.Debug("bucket", "bucket", bucket)
		if len(bucket) == 0 {
			yap.Error("invalid hash", "hash", hash)
			w.WriteHeader(http.StatusUnauthorized)
		}

		sessKey, _ := r.Cookie("X-Session")
		if sessKey != nil {
			opt := man.Sessions.Get(sessKey.Value)
			yap.Debug("attempting cookie", "cookie", sessKey.Value)
			if opt.Ok() {
				yap.Debug("session recognized", "label", opt.Unwrap().Label)
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		yap.Debug("authenticating", "hash", hash, "host", bucket)
		if len(hash) != 44 {
			yap.Error("invalid hash", "hash", hash, "remote", r.RemoteAddr)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		yap.Debug("using link")
		_, ok := man.UseLink(hash, w)
		if !ok {
			yap.Error("error using link")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		yap.Debug("link used")

		http.Redirect(w, r, "/", http.StatusFound)
	})

	addr := outward
	yap.Info("starting fwauth server", "addr", URL(addr).Format())
	http.ListenAndServe(addr, mux)
}
