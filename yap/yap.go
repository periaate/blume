package yap

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/periaate/blume/fsio"
	. "github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
	. "github.com/periaate/blume/typ"
)

var (
	includeY = false
	includeM = false
	includeD = false
	includeh = true
	includem = true
	includes = true
)

func IncludeTimes(Y, M, D, h, m, s bool) {
	includeY = Y
	includeM = M
	includeD = D
	includeh = h
	includem = m
	includes = s
}

const webSafeBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func Encode(n int) string { return string(webSafeBase64[n%64]) }

// kinoFormat returns a 6 character timestamp which encodes a full timestamp.
// Each character is a base64 encoded digit.
func kinoFormat(tim time.Time) (res string) {
	year, month, day := tim.Date()
	hour, min, sec := tim.Clock()
	if includeY {
		res += Encode(year - 1970)
	}
	if includeM {
		res += Encode(int(month) - 1)
	}
	if includeD {
		res += Encode(day - 1)
	}
	if includeh {
		res += Encode(hour)
	}
	if includem {
		res += Encode(min)
	}
	if includes {
		res += Encode(sec)
	}
	return
}

func Time() string { return kinoFormat(time.Now()) }

type Level int

const (
	L_Error Level = -1
	L_Info  Level = 0
	L_Debug Level = 1
)

// PairReducer constructs a reducer to generate pairs from the array
func PairReducer[A any]() T.Reducer[A, [][]A] {
	return func(arr []A) [][]A {
		pairs := [][]A{}
		for i := 0; i < len(arr); i += 2 {
			if i+1 < len(arr) {
				pairs = append(pairs, []A{arr[i], arr[i+1]})
			}
		}
		return pairs
	}
}

var l = L_Info

func SetLevel(level Level) { l = level }

func (l Level) String() string {
	switch l {
	case L_Error:
		return String("E").Colorize(LightRed).String()
	case L_Info:
		return String("I").Colorize(Cyan).String()
	case L_Debug:
		return String("D").Colorize(LightYellow).String()
	default:
		return String("-").String()
	}
}

func Log(out io.Writer, format string, src string, level Level, msg string, args ...any) {
	res := PairReducer[any]()(args)
	strs := Map[[]any, string](
		func(a []any) string {
			a1 := String(fmt.Sprintf("%s", a[0])).ToUpper().Colorize(LightYellow).String()
			a = a[1:]
			res := Map[any, string](func(a any) string {
				switch v := a.(type) {
				case string:
					return String(fmt.Sprintf("%q", v)).Colorize(Yellow).String()
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					return String(fmt.Sprint(v)).Colorize(Cyan).String()
				case bool:
					return String(fmt.Sprint(v)).Colorize(LightGreen).String()
				default:
					return String(fmt.Sprint(a)).Colorize(LightGreen).String()
				}
			})(a)
			return fmt.Sprintf("%s:<%s>;", a1, strings.Join(res, ", "))
		},
	)(res)

	fmt.Fprintf(
		out,
		Or(format, "%s %s %s [%s] %s\n"),
		level.String(),
		src,
		String(Time()).Colorize(Yellow).String(),
		String(msg),
		strings.Join(strs, " "),
	)
}

func caller(file string, line int) string {
	return fmt.Sprintf("%s:%d", fsio.Name(file), line)
}

func Info(msg string, args ...any) {
	if l < L_Info {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	Assert(ok, "Failed to get caller")
	Log(os.Stdout, "", caller(file, line), L_Info, msg, args...)
}

func Error(msg string, args ...any) {
	if l < L_Error {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	Assert(ok, "Failed to get caller")
	Log(os.Stdout, "", caller(file, line), L_Error, msg, args...)
}

func Debug(msg string, args ...any) {
	if l < L_Debug {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	Assert(ok, "Failed to get caller")
	Log(os.Stdout, "", caller(file, line), L_Debug, msg, args...)
}
