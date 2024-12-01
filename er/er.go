package er

import (
	"net/http"
	"strings"
)

type Net interface {
	error
	Status() int
	Respond(w http.ResponseWriter)
}

type Custom struct {
	HTTPStatus int
	Msg        string
}

func (c Custom) Error() string { return c.Msg }
func (c Custom) Status() int   { return c.HTTPStatus }
func (c Custom) Respond(w http.ResponseWriter) {
	http.Error(w, c.Msg, c.HTTPStatus)
}

func Free(status int, pairs ...string) Net {
	sb := strings.Builder{}
	var i int
	for i = 0; i < len(pairs); i += 2 {
		sb.WriteString(pairs[i])
		sb.WriteString(" [")
		sb.WriteString(pairs[i+1])
		sb.WriteString("] ")
	}

	if i < len(pairs) {
		sb.WriteString(pairs[i])
	}

	return Custom{
		HTTPStatus: status,
		Msg:        sb.String(),
	}
}

// 4xx
func BadRequest(pairs ...string) Net   { return Free(http.StatusBadRequest, pairs...) }   // 400
func Unauthorized(pairs ...string) Net { return Free(http.StatusUnauthorized, pairs...) } // 401
func Forbidden(pairs ...string) Net    { return Free(http.StatusForbidden, pairs...) }    // 403

func NotFound(kind, val, from string, pairs ...string) Net {
	pairs = append([]string{"error:", kind, "with value", val, "not found in", from}, pairs...)
	return Free(http.StatusNotFound, pairs...)
}

func MethodNotAllowed(pairs ...string) Net { return Free(http.StatusMethodNotAllowed, pairs...) } // 405
func Timeout(pairs ...string) Net          { return Free(http.StatusRequestTimeout, pairs...) }
func Conflict(pairs ...string) Net         { return Free(http.StatusConflict, pairs...) } // 409
func LengthRequired(pairs ...string) Net   { return Free(http.StatusLengthRequired, pairs...) }

func UnsupportedMediaType(pairs ...string) Net {
	return Free(http.StatusUnsupportedMediaType, pairs...)
}

func TypeRequired() Net {
	return Free(http.StatusBadRequest, "Server rejected the request because the [Content-Type] header field is not defined and the server requires it.")
}

func Exists(name string) Net {
	return Free(http.StatusConflict, "The resource", name, "already exists.")
}

// 5xx
func InternalServerError(pairs ...string) Net { return Free(http.StatusInternalServerError, pairs...) } // 500
func NotImplemented(pairs ...string) Net      { return Free(http.StatusNotImplemented, pairs...) }
func InsufficientStorage(pairs ...string) Net { return Free(http.StatusInsufficientStorage, pairs...) }
