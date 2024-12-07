package hnet

import (
	"net/http"
	"net/url"

	. "github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

//blume:derive String
type URL string

func (u URL) ToURL() (*url.URL, error) {
	return url.Parse(string(u.Format()))
}

func POST(r Request) Request {
	r.Method = "POST"
	return r
}

func GET(r Request) Request {
	r.Method = "GET"
	return r
}

func PUT(r Request) Request {
	r.Method = "PUT"
	return r
}

func DELETE(r Request) Request {
	r.Method = "DELETE"
	return r
}

func HEAD(r Request) Request {
	r.Method = "HEAD"
	return r
}

func OPTIONS(r Request) Request {
	r.Method = "OPTIONS"
	return r
}

func PATCH(r Request) Request {
	r.Method = "PATCH"
	return r
}

func (u URL) ToRequest(opts ...T.Monadic[Request, Request]) (req Request) {
	if len(u) == 0 {
		req.NetErr = Free(400, "Invalid URL", "url", string(u))
		return
	}
	req = Request{Request: &http.Request{
		Header: http.Header{},
	}}
	url, err := u.ToURL()
	if err != nil {
		req.NetErr = Free(400, "Invalid URL", "url", string(u))
		req.Err = err
		return
	}
	req.URL = url
	req = Pipe[Request](opts...)(req)
	if req.Method == "" {
		req.NetErr = Free(400, "Method not set")
	}
	return
}

func (u URL) Format(options ...T.Monadic[URL, URL]) URL {
	return Pipe[URL](ArrayOrDefault(options, AsProtocol(HTTP))...)(u)
}

//blume:derive String
type Protocol string

const (
	HTTP  Protocol = "http"
	HTTPS Protocol = "https"
	WS    Protocol = "ws"
	WSS   Protocol = "wss"
)

func (u URL) AsProtocol(protocol Protocol) URL {
	if len(u) == 0 {
		return URL(protocol + "://")
	}
	if HasProtocol(u) {
		return URL(u.ReplaceRegex(".*://", string(protocol+"://")))
	}
	return URL(protocol) + "://" + u
}

func AsProtocol(protocol Protocol) T.Monadic[URL, URL] {
	return func(u URL) URL { return u.AsProtocol(protocol) }
}
