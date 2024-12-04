package hnet

import (
	"fmt"
	"io"
	"net/http"

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
