package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/periaate/blume"
	"github.com/periaate/blume/types/maps"
	"github.com/periaate/blume/types/str"
	"github.com/periaate/blume/yap"
)

type Link struct {
	Hash       string
	Origin     string
	Expiration time.Time
	Uses       int
	Duration   time.Duration
}

type Sess struct {
	Hash       string
	Origin     string
	Expiration time.Time
}

type Serv struct {
	Links    maps.Map[string, *Link]
	Sessions maps.Map[string, *Sess]
}

func main() {
	s := &Serv{
		maps.New(func(_ string, link *Link) (ok bool) {
			if link == nil {
				return
			}
			if time.Until(link.Expiration) < 0 {
				return
			}
			link.Uses--
			return link.Uses > 0
		}),
		maps.New(func(_ string, sess *Sess) (ok bool) {
			if sess == nil {
				return
			}
			return time.Until(sess.Expiration) > 0
		}),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /gen/{origin}/{exp}/{dur}/{uses}", func(w http.ResponseWriter, r *http.Request) {
		origin, err := url.Parse("https://" + r.PathValue("origin"))
		if err != nil {
			http.Error(w, "origin is not a valid URL", 400)
			return
		}

		exp, err := str.ToInt64(r.PathValue("exp"))
		if err != nil {
			http.Error(w, "couldn't parse expiration to int64", 400)
			return
		}
		dur, err := str.ToInt64(r.PathValue("dur"))
		if err != nil {
			http.Error(w, "couldn't parse duration to int64", 400)
			return
		}
		uses, err := str.ToInt(r.PathValue("uses"))
		if err != nil {
			http.Error(w, "couldn't parse uses to int", 400)
			return
		}

		println(origin.Host)

		l := &Link{
			Hash:       gen(),
			Origin:     origin.Host,
			Expiration: time.Now().Add(time.Duration(exp) * time.Minute),
			Duration:   time.Duration(dur),
			Uses:       uses,
		}

		if !s.Links.Set(l.Hash, l) {
			http.Error(w, "the arguments provided do not construct a valid link", 400)
			return
		}

		w.Write([]byte(origin.String() + "/" + l.Hash))
	})

	findCookie := func(r *http.Request) (ok bool) {
		cook, err := r.Cookie("FW-Auth-Session-Token")
		if err != nil {
			return
		}

		val, ok := s.Sessions.Get(cook.Value)
		fmt.Println(r.URL.Host, val.Origin)
		if !ok {
			return
		}
		if r.URL.Host != val.Origin {
			return
		}
		return true
	}

	http.HandleFunc("GET /fw-auth/{origin}/", func(w http.ResponseWriter, r *http.Request) {
		if findCookie(r) {
			w.WriteHeader(200)
			return
		}
		w.WriteHeader(401)
	})

	http.HandleFunc("GET /fw-auth/{origin}/{hash}", func(w http.ResponseWriter, r *http.Request) {
		if findCookie(r) {
			w.WriteHeader(200)
			return
		}

		origin := r.PathValue("origin")
		hash := r.PathValue("hash")
		link, ok := s.Links.Get(hash)
		println(ok)
		switch {
		case !ok, link.Origin != origin:
			w.WriteHeader(401)
		default:
			res := gen()
			cookie := &http.Cookie{
				Name:  "FW-Auth-Session-Token",
				Value: res,
			}

			ok := s.Sessions.Set(res, &Sess{
				link.Hash,
				link.Origin,
				time.Now().Add(link.Duration * time.Minute),
			})

			if !ok {
				w.WriteHeader(401)
				return
			}

			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})

	go func() {
		addr := blume.Or("127.0.0.1:7590", os.Getenv("FW_AUTH_GEN_ADDR"))
		yap.Info("serving link gen server", "http://"+addr)
		yap.Fatal("error running gen server", http.ListenAndServe(addr, mux))
	}()

	addr := blume.Or("127.0.0.1:7595", os.Getenv("FW_AUTH_ADDR"))
	yap.Info("serving fwauth server", "http://"+addr)
	yap.Fatal("error running fwauth server", http.ListenAndServe(addr, nil))
}

func gen() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}
