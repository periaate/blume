package er

import (
	"fmt"
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
}

func (e Unexpected) Error() string { return e.Msg }
func (e Unexpected) Status() int   { return 500 }

type InvalidData struct {
	Has     string
	Expects string
	Msg     string
}

func (e InvalidData) Error() string {
	return fmt.Sprintf("%s: has: [%s] expects: [%s]", e.Msg, e.Has, e.Expects)
}

type BadRequest struct {
	Requested string
	From      string
	With      string
	Msg       string
}

type NotFound BadRequest

func (e NotFound) Error() string { return e.Error() }

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

func (e BadRequest) Status() int { return 400 }
func (e NotFound) Status() int   { return 404 }

// func (e InvalidData) Status() int { return 500 }
