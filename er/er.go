package er

import (
	"fmt"
	"net/http"

	"github.com/periaate/blume/gen"
)

func Is[A any](a any) (ok bool) {
	_, ok = a.(A)
	return
}

type Net interface {
	error
	Status() int
}

type Log interface {
	// Log must be a slog compliant logger.
	Log(msg string, args ...any)
}

type Unexpected struct {
	Msg string

	StatusError
}

type StatusError struct{ HTTPStatus int }

// OrStatus returns the custom status code or the default status code of a Net error.
func (e StatusError) OrStatus(def int) int { return gen.Or(e.HTTPStatus, def) }

func (e Unexpected) Error() string { return e.Msg }
func (e Unexpected) Status() int   { return e.OrStatus(http.StatusInternalServerError) }

type InvalidData struct {
	Has     string
	Expects string
	Msg     string

	StatusError
}

func (e InvalidData) Error() string {
	return fmt.Sprintf("%s: has: [%s] expects: [%s]", e.Msg, e.Has, e.Expects)
}

type BadRequest struct {
	Requested string
	From      string
	With      string
	Msg       string

	StatusError
}

type NotFound struct {
	Requested string
	From      string
	With      string
	Msg       string

	StatusError
}

func (e NotFound) Error() string {
	switch {
	case e.Requested == "", e.From == "":
		return e.Msg
	case e.With != "":
		return fmt.Sprintf("requested [%s] from [%s] with [%s]: %s", e.Requested, e.From, e.With, e.Msg)
	default:
		return fmt.Sprintf("requested [%s] from [%s]: %s", e.Requested, e.From, e.Msg)
	}
}

func (e BadRequest) Error() string {
	switch {
	case e.Requested == "", e.From == "":
		return e.Msg
	case e.With != "":
		return fmt.Sprintf("requested [%s] from [%s] with [%s]: %s", e.Requested, e.From, e.With, e.Msg)
	default:
		return fmt.Sprintf("requested [%s] from [%s]: %s", e.Requested, e.From, e.Msg)
	}
}

func (e BadRequest) Status() int { return e.OrStatus(http.StatusBadRequest) }
func (e NotFound) Status() int   { return e.OrStatus(http.StatusNotFound) }

type Custom struct {
	Msg string

	HTTPStatus int
}

func (e Custom) Error() string { return e.Msg }
func (e Custom) Status() int   { return e.HTTPStatus }

type Internal struct {
	Msg string
}

func (e Internal) Error() string { return e.Msg }
func (e Internal) Status() int   { return http.StatusInternalServerError }
