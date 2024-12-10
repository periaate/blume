package hnet

import (
	"net/http"
	"strings"

	. "github.com/periaate/blume/core"
)

var _ Error[int] = Nerr{}

type NetError interface {
	Error[int]
	Respond(http.ResponseWriter)
}

type Nerr struct { Err[int] }

func (c Nerr) Respond(w http.ResponseWriter) { http.Error(w, c.Error(), c.Val) }

func Free(status int, msg string, pairs ...string) NetError {
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

	if i < len(pairs) { sb.WriteString(pairs[i]) }

	return Nerr{
		Err[int]{
			Val: status,
			Str: sb.String(),
		},
	}
}

func FreeWrap(err error, status int, msg string, pairs ...string) NetError {
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

	if i < len(pairs) { sb.WriteString(pairs[i]) }

	return Nerr{
		Err[int]{
			Val: status,
			Str: sb.String(),
			Err: err,
		},
	}
}
