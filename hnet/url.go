package hnet

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	. "github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

type Response struct {
	*http.Response
	NetErr
}

func (r Response) Assert(conds ...T.Predicate[any]) Response {
	if r.NetErr != nil {
		return r
	}

	for _, cond := range conds {
		if !cond(r) {
			r.NetErr = Bad_Request.Def("failed assertions")
			return r
		}
	}

	return r
}

func Println(s string) { fmt.Println(s) }

func (r Response) String(f func(string)) Response {
	if r.NetErr != nil {
		return r
	}
	r.UseBody(func(r io.Reader) bool {
		bar, err := io.ReadAll(r)
		if err != nil {
			return false
		}
		f(string(bar))
		return true
	})
	return r
}

func UseBody(f func(io.Reader) bool) T.Transformer[Response] {
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
			r.NetErr = Free(500, "Body assertion failed")
		}
		return r
	}
}

func (r Response) UseBody(f func(io.Reader) bool) Response {
	if r.NetErr != nil {
		return r
	}
	return UseBody(f)(r)
}

type Request struct {
	*http.Request
	NetErr
	Err error
}

func (r Request) FileAsBody(fp string) Request {
	if r.NetErr != nil {
		return r
	}
	f, err := os.Open(fp)
	if err != nil {
		r.NetErr = Not_Found.Def(err.Error())
	}
	r.WithBody(f)
	return r
}

func (r Request) WithHeaders(tuples ...[2]string) Request {
	if r.NetErr != nil {
		return r
	}
	for _, v := range tuples {
		r.Header.Add(v[0], v[1])
	}
	return r
}

func (r Request) WithBody(rc io.ReadCloser) Request {
	if r.NetErr != nil {
		return r
	}
	r.Body = rc
	return r
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

func WithHeaders(tuples ...[2]string) T.Transformer[Request] {
	return func(r Request) Request {
		if r.Header == nil {
			r.Header = http.Header{}
		}
		for _, tuple := range tuples {
			r.Header.Set(tuple[0], tuple[1])
		}
		return r
	}
}

func WithBody(body io.ReadCloser) T.Transformer[Request] {
	return func(r Request) Request {
		r.Body = body
		return r
	}
}

func (u URL) ToRequest(opts ...T.Transformer[Request]) (req Request) {
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

func (u URL) Format(options ...T.Transformer[URL]) URL {
	return Pipe[URL](ArrayOrDefault(options, AsProtocol(HTTP))...)(u)
}

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

func AsProtocol(protocol Protocol) T.Transformer[URL] {
	return func(u URL) URL { return u.AsProtocol(protocol) }
}
