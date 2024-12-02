package hnet

import (
	"io"
	"net/http"
	"net/url"

	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/str"
)

type Response struct {
	*http.Response
	NetErr
}

func UseBody(f func(io.Reader) bool) gen.Transformer[Response] {
	return func(r Response) Response {
		defer r.Body.Close()
		if r.NetErr != nil {
			return r
		}

		r.NetErr = Def(r.StatusCode)
		if r.NetErr != nil {
			return r
		}

		if !f(r.Body) {
			r.NetErr = Status(r.StatusCode).Def()
		}
		return r
	}
}

func (r Response) UseBody(f func(io.Reader) bool) Response { return UseBody(f)(r) }

type Request struct {
	*http.Request
	NetErr
}

func (r Request) Call() Response {
	if r.NetErr != nil {
		return Response{NetErr: r.NetErr}
	}
	client := &http.Client{}
	resp, err := client.Do(r.Request)
	if err != nil {
		return Response{NetErr: Free(500, err.Error())}
	}
	return Response{Response: resp}
}

type URL string

func (u URL) ToURL() (*url.URL, error) {
	return url.Parse(string(u))
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

func WithHeaders(tuples ...[2]string) gen.Transformer[Request] {
	return func(r Request) Request {
		for _, tuple := range tuples {
			r.Header.Set(tuple[0], tuple[1])
		}
		return r
	}
}

func WithBody(body io.ReadCloser) gen.Transformer[Request] {
	return func(r Request) Request {
		r.Body = body
		return r
	}
}

func (u URL) ToRequest(opts ...gen.Transformer[Request]) (req Request, err error) {
	req = Request{Request: &http.Request{}}
	url, err := u.ToURL()
	if err != nil {
		req.NetErr = Free(400, "Invalid URL", "url", string(u))
		return req, err
	}
	req.URL = url
	req = gen.Pipe[Request](opts...)(req)
	if req.Method == "" {
		req.NetErr = Free(400, "Method not set", "err", req.NetErr.Error())
	}
	return
}

func NewURL[S ~string](inp S, options ...gen.Transformer[URL]) URL {
	options = gen.ArrayOrDefault(options, Opt_HTTP)
	return gen.Pipe[URL](options...)(URL(inp))
}

func Opt_HTTP(s URL) URL {
	if len(s) == 0 {
		return "http://"
	}
	if HasProtocol(s) {
		return str.ReplaceRegex[URL](".*://", "http://")(s)
	}
	return fsio.Join("http://", s)
}

func Opt_HTTPS(s URL) URL {
	if len(s) == 0 {
		return "https://"
	}
	if HasProtocol(s) {
		return str.ReplaceRegex[URL](".*://", "https://")(s)
	}
	return fsio.Join("https://", s)
}
