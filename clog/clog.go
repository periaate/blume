/*
Package clog wraps log/slog with a normalized indent, humanized, and colorized style.

Example:

	DEBUG @ main.go:111      MSG:<a message>; KEY:<Values here>; err:<nil>;
	DEBUG @ main.go:111      MSG:<another message>; KEY:<Values here longer value>; err:<nil>;
	DEBUG @ main.go:111      MSG:<a message>;       KEY:<err will be adjusted>;     err:<nil>;

## TODO

  - [ ] Rewrite the handler for greater flexibility and customization.
  - [ ] Decide whether to use external libraries for coloring and formatting.
*/
package clog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/periaate/blume/str"
	"github.com/periaate/blume/val"
)

// ClogHandler is a log/slog handler.
type ClogHandler struct {
	St    *Styles
	Level slog.Leveler
	Lvl   int

	Mut *sync.Mutex
	Out io.Writer

	// MaxLen is the maximum length of a single value. If the value is longer, it will be cut.
	MaxLen int

	// indLens is used internally to store the length at the index of the key.
	// Updated whenever an index is larger than the current length.
	indLens map[int]int

	Getter func(slog.Value) string
}

func (h *ClogHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return int(lvl) >= int(h.Lvl)
}

func (h *ClogHandler) WithGroup(name string) slog.Handler       { return h }
func (h *ClogHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }

var pred = str.Contains("blume/clog", "slog/logger")

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
		var ok bool
		ok = true
		var f string
		var l int
		for i := 1; i < 10 && ok; i++ {
			_, f, l, ok = runtime.Caller(i)
			if pred(f) {
				continue
			}
			break
		}
		nbuf = h.appendAttr(slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f, l)))
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

	_, err := fmt.Fprintln(h.Out, string(buf))
	return err
}

func (h *ClogHandler) appendAttr(attr slog.Attr) (buf []byte) {
	attr.Value = attr.Value.Resolve()
	if attr.Equal(slog.Attr{}) {
		return buf
	}

	var res string
	switch attr.Key {
	case slog.LevelKey:
		nv := attr.Value.String()
		switch nv {
		case "DEBUG":
			res = val.Colorize(val.White, nv)
		case "INFO":
			res = val.Colorize(val.Cyan, nv+" ")
		case "WARN":
			res = val.Colorize(val.Yellow, nv+" ")
		case "ERROR":
			res = val.Colorize(val.Red, nv)
		}
	case slog.TimeKey:
		if h.St.TimeStamp {
			res = fmt.Sprintf("%s  ", dumbFormat(attr.Value.Time()))
		}
	case slog.SourceKey:
		sp := strings.Split(attr.Value.String(), "/")
		f := sp[len(sp)-1]
		f = maxStrLen(f, 15)
		res = fmt.Sprintf(" @ %s%s  ", f, strings.Repeat(" ", 15-len(f)))
	default:
		val := attr.Value
		switch val.Kind() {
		case slog.KindGroup:
			arr := []string{}
			for _, v := range val.Group() {
				arr = append(arr, h.Getter(v.Value))
			}
			res = fmt.Sprintf("[%s]", strings.Join(arr, ", "))
		default:
			res = h.Getter(val)
		}

		res = fmt.Sprintf("%s%s%s%s", strings.ToUpper(attr.Key), h.St.Delim[0], res, h.St.Delim[1])
	}

	buf = fmt.Appendf(buf, "%s", res)
	return buf
}

func DefaultGetter(vall slog.Value) (res string) {
	switch vall.Kind() {
	case slog.KindTime:
		res = vall.Time().Format(time.RFC3339Nano)
	case slog.KindInt64:
		res = val.HumanizeNumber(vall.Int64())
	case slog.KindUint64:
		res = val.HumanizeNumber(vall.Uint64())
	case slog.KindFloat64:
		res = val.HumanizeNumber(vall.Float64())
	default:
		res = vall.String()
	}

	return
}

func dumbFormat(tim time.Time) (res string) {
	a := tim.Format("06/002")

	h, m, s := tim.Clock()
	// 3 decimal places
	n := tim.Nanosecond() / 1e6
	s += m*60 + h*3600

	res = fmt.Sprintf("%s/%d.%d", a, s, n)

	sp := strings.Split(res, ".")
	if len(sp[1]) < 3 {
		sp[1] += strings.Repeat("0", 3-len(sp[1]))
	}

	res = strings.Join(sp, ".")
	return
}

func maxStrLen(str string, max int) string {
	if len(str) > max {
		// cut the beginning
		str = "..." + str[len(str)-(max-3):]
	}
	return str
}
