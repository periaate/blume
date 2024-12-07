package hnet

import (
	"io"
	"net/http"
	"os"

	"github.com/periaate/blume/gen/T"
)

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

func WithHeaders(tuples ...[2]string) T.Monadic[Request, Request] {
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

func WithBody(body io.ReadCloser) T.Monadic[Request, Request] {
	return func(r Request) Request {
		r.Body = body
		return r
	}
}
