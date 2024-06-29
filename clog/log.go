package clog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"blume/core/gen"
	"blume/core/val"
)

type ClogHandler struct {
	St    *Styles
	Level slog.Leveler
	Lvl   int

	Mut *sync.Mutex
	Out io.Writer

	MaxLen int

	indLens map[int]int
}

var defLog = DefaultClog()

func GetDefaultClog() *slog.Logger { return defLog }

func Error(msg string, args ...any) { defLog.Error(msg, args...) }
func Info(msg string, args ...any)  { defLog.Info(msg, args...) }
func Warn(msg string, args ...any)  { defLog.Warn(msg, args...) }
func Debug(msg string, args ...any) { defLog.Debug(msg, args...) }

func SetLogLoggerLevel(lvl slog.Level) { defLog.Handler().(*ClogHandler).SetLogLoggerLevel(lvl) }

type Logger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}

type DummyLogger struct{}

func (DummyLogger) Error(_ string, _ ...any) {}
func (DummyLogger) Info(_ string, _ ...any)  {}
func (DummyLogger) Warn(_ string, _ ...any)  {}
func (DummyLogger) Debug(_ string, _ ...any) {}

func DefaultClog() *slog.Logger { return NewClog(os.Stdout, slog.LevelInfo, MaxLen(50)) }

func NewClog(w io.Writer, lvl slog.Level, opts ...gen.Option[*ClogHandler]) *slog.Logger {
	h := New(w, lvl, nil)

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

func Style(st *Styles) gen.Option[*ClogHandler] { return func(h *ClogHandler) { h.St = st } }
func MaxLen(l int) gen.Option[*ClogHandler]     { return func(h *ClogHandler) { h.MaxLen = l } }

func (l *ClogHandler) SetLogLoggerLevel(lvl slog.Level) { l.Lvl = int(lvl) }

type Styles struct {
	TimeStamp bool
	Delim     [2]string
}

func NewStyles(st *Styles) *Styles {
	if st == nil {
		st = &Styles{}
	}
	if len(st.Delim[0]) == 0 && len(st.Delim[1]) == 0 {
		if st.Delim == [2]string{} {
			st.Delim = [2]string{
				":" + val.Color(val.LightYellow, "<"),
				val.EndColor(">") + "; ",
			}
		}
	}

	return st
}

func (h *ClogHandler) DefGetV(vall slog.Value) string {
	switch vall.Kind() {
	case slog.KindTime:
		return vall.Time().Format(time.RFC3339Nano)
	case slog.KindInt64, slog.KindUint64, slog.KindFloat64:
		return val.HumanNumber(vall.String())
	case slog.KindGroup:
		arr := []string{}
		for _, v := range vall.Group() {
			arr = append(arr, h.DefGetV(v.Value))
		}
		return fmt.Sprintf("[%s]", strings.Join(arr, ", "))
	default:
		v := vall.String()

		if h.MaxLen > 0 {
			v = MaxStrLen(v, h.MaxLen)
		}

		return v
	}
}

func New(out io.Writer, lvl slog.Level, st *Styles) *ClogHandler {
	h := &ClogHandler{
		St:      st,
		Level:   lvl,
		Mut:     &sync.Mutex{},
		Out:     out,
		indLens: make(map[int]int),
	}
	if h.Level == nil {
		h.Level = slog.LevelInfo
	}
	return h
}

func (h *ClogHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return int(lvl) >= int(h.Lvl)
}

func (h *ClogHandler) WithGroup(name string) slog.Handler       { return h }
func (h *ClogHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }

func (h *ClogHandler) Handle(ctx context.Context, r slog.Record) error {
	var nbuf []byte
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		nbuf = h.appendAttr(slog.Time(slog.TimeKey, r.Time))
		buf = append(buf, nbuf...)
	}

	nbuf = h.appendAttr(slog.Any(slog.LevelKey, r.Level))
	buf = append(buf, nbuf...)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		nbuf = h.appendAttr(slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)))
		buf = append(buf, nbuf...)
	}

	nbuf = h.appendAttr(slog.String(slog.MessageKey, r.Message))
	buf = append(buf, nbuf...)

	h.Mut.Lock()
	defer h.Mut.Unlock()

	var i int
	r.Attrs(func(attr slog.Attr) bool {
		bl := utf8.RuneCount(buf)

		l, ok := h.indLens[i]
		switch {
		case !ok:
			h.indLens[i] = bl
		case l > bl:
			buf = append(buf, strings.Repeat(" ", l-bl)...)
		case l < bl:
			h.indLens[i] = bl
		}
		nbuf = h.appendAttr(attr)
		buf = append(buf, nbuf...)

		i++
		return true
	})
	str := string(buf)

	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\\n", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\\r", "")
	_, err := h.Out.Write([]byte(fmt.Sprintln(str)))

	return err
}

func (h *ClogHandler) appendAttr(attr slog.Attr) (buf []byte) {
	attr.Value = attr.Value.Resolve()
	if attr.Equal(slog.Attr{}) {
		return buf
	}

	switch attr.Key {
	case slog.LevelKey:
		nv := attr.Value.String()
		var sv string
		switch nv {
		case "DEBUG":
			sv = val.Colorize(val.White, nv)
		case "INFO":
			sv = val.Colorize(val.Cyan, nv+" ")
		case "WARN":
			sv = val.Colorize(val.Yellow, nv+" ")
		case "ERROR":
			sv = val.Colorize(val.Red, nv)
		}
		buf = fmt.Appendf(buf, "%s", sv)
		return buf
	case slog.TimeKey:
		if h.St.TimeStamp {
			buf = fmt.Appendf(buf, "%s", h.DefGetV(attr.Value))
		}
		return buf
	case slog.SourceKey:
		ind := strings.LastIndex(attr.Value.String(), "/")
		if ind != -1 {
			f := fmt.Sprintf(" @ %s", attr.Value.String()[ind+1:])
			buf = fmt.Appendf(buf, "%s%s", f, strings.Repeat(" ", 20-len(f)))
			return buf
		}
	}

	if res, ok := h.fmtKV(attr); ok {
		buf = fmt.Appendf(buf, "%s", res)
	}
	return buf
}

func (h *ClogHandler) fmtKV(val slog.Attr) (res string, ok bool) {
	return fmt.Sprintf("%s%s%s%s", strings.ToUpper(val.Key), h.St.Delim[0], h.DefGetV(val.Value), h.St.Delim[1]), true
}

func MaxStrLen(str string, max int) string {
	if len(str) > max {
		// cut the beginning
		str = "..." + str[len(str)-(max-3):]
	}
	return str
}
