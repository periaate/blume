package clog

import (
	"log/slog"
	"os"
	"testing"

	"github.com/periaate/blume/core/val"
)

type show struct {
	a string
	b []string
	c int
	f float64
}

func TestLog(t *testing.T) {
	clog := NewClog(os.Stdout, slog.LevelDebug, Style(&Styles{
		TimeStamp: false,
		Delim: [2]string{
			":" + val.Color(val.LightYellow, "<"),
			val.EndColor(">") + "; ",
		},
	}))

	// glog.Debug("testing", "abc", "dfg")
	s := show{
		a: "TestValue",
		b: []string{"SArrOne", "SArrTwo", "SArr3"},
		c: 2313421,
		f: 2142.14159265359,
	}

	clog.Error("begin flag capture", "name", s.a, "len", len(s.b), "args", s.b, "val", s.c, "fl", s.f)
	clog.Info("begin flag capture", "name", s.a, "len", len(s.b), "args", s.b, "val", s.c, "fl", s.f)
	clog.Warn("begin flag capture", "name", s.a, "len", len(s.b), "args", s.b, "val", s.c, "fl", s.f)
	clog.Debug("begin flag capture", "name", s.a, "len", len(s.b), "args", s.b, "val", s.c, "fl", s.f)
}
