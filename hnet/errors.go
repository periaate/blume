package hnet

import (
	"net/http"
	"strings"
)

type NetErr interface {
	error
	Status() int
	Respond(w http.ResponseWriter)
}

type Error struct {
	HTTPStatus int
	Msg        string
}

func (c Error) Error() string { return c.Msg }
func (c Error) Status() int   { return c.HTTPStatus }
func (c Error) Respond(w http.ResponseWriter) {
	http.Error(w, c.Msg, c.HTTPStatus)
}

func Free(status int, msg string, pairs ...string) NetErr {
	sb := strings.Builder{}
	sb.WriteString(msg)
	sb.WriteString(" ")
	var i int
	for i = 0; i+1 < len(pairs); i += 2 {
		sb.WriteString(pairs[i])
		sb.WriteString(" [")
		sb.WriteString(pairs[i+1])
		sb.WriteString("] ")
	}

	if i < len(pairs) {
		sb.WriteString(pairs[i])
	}

	return Error{
		HTTPStatus: status,
		Msg:        sb.String(),
	}
}
