package clog

import (
	"io"
	"log/slog"
	"os"
	"sync"

	"github.com/periaate/blume/val"
)

var defLog = DefaultClog()

func GetDefaultClog() *slog.Logger  { return defLog }
func SetDefaultClog(l *slog.Logger) { defLog = l }

const (
	LevelError = slog.LevelError
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelDebug = slog.LevelDebug
)

func SetLogLoggerLevel(lvl slog.Level) { defLog.Handler().(*ClogHandler).SetLogLoggerLevel(lvl) }

type Logger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}

// Dummy is a dummy logger that does nothing.
type Dummy struct{}

func (Dummy) Error(_ string, _ ...any) {}
func (Dummy) Info(_ string, _ ...any)  {}
func (Dummy) Warn(_ string, _ ...any)  {}
func (Dummy) Debug(_ string, _ ...any) {}

func DefaultClog() *slog.Logger { return NewClog(os.Stdout, LevelInfo, MaxLen(50)) }

// NewClog creates a new clog logger with the given writer, level, and options.
func NewClog(w io.Writer, lvl slog.Level, opts ...func(*ClogHandler)) *slog.Logger {
	h := New(w, lvl, nil, DefaultGetter)

	for _, opt := range opts {
		opt(h)
	}

	if h.Level == nil {
		h.Level = slog.LevelInfo
	}

	if h.St == nil {
		Style(NewStyles(nil))(h)
	}

	return slog.New(h)
}

func Style(st *Styles) func(*ClogHandler) { return func(h *ClogHandler) { h.St = st } }
func MaxLen(l int) func(*ClogHandler)     { return func(h *ClogHandler) { h.MaxLen = l } }

func (l *ClogHandler) SetLogLoggerLevel(lvl slog.Level) { l.Lvl = int(lvl) }

type Styles struct {
	TimeStamp bool
	Delim     [2]string
}

func NewStyles(st *Styles) *Styles {
	if st == nil {
		st = &Styles{}
	}
	if st.Delim == [2]string{} {
		st.Delim = [2]string{
			":" + val.Color(val.LightYellow, "<"),
			val.EndColor(">") + "; ",
		}
	}

	return st
}

func New(out io.Writer, lvl slog.Level, st *Styles, getter func(slog.Value) string) *ClogHandler {
	h := &ClogHandler{
		St:      st,
		Level:   lvl,
		Mut:     &sync.Mutex{},
		Out:     out,
		indLens: make(map[int]int),
		Getter:  getter,
	}
	if h.Level == nil {
		h.Level = slog.LevelInfo
	}
	return h
}
