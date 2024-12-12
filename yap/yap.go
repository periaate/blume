package yap

import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"strings"
	"time"

	. "github.com/periaate/blume"
)

type Yapfig struct {
	Years   bool
	Months  bool
	Days    bool
	Hours   bool
	Minutes bool
	Seconds bool

	ShowFile  bool
	ShowLevel bool
	ShowTime  bool

	Level Level
}

func Configure(yapfig Yapfig) {
	includeY = yapfig.Years
	includeM = yapfig.Months
	includeD = yapfig.Days
	includeh = yapfig.Hours
	includem = yapfig.Minutes
	includes = yapfig.Seconds

	showFile = yapfig.ShowFile
	showLevel = yapfig.ShowLevel
	showTime = yapfig.ShowTime

	l = yapfig.Level
}

var (
	showFile  = true
	showLevel = true
	showTime  = true
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
	if includeY { res += Encode(year - 1970) }
	if includeM { res += Encode(int(month) - 1) }
	if includeD { res += Encode(day - 1) }
	if includeh { res += Encode(hour) }
	if includem { res += Encode(min) }
	if includes { res += Encode(sec) }
	return
}

func Time() string { return kinoFormat(time.Now()) }

type Level int

const (
	L_Fatal Level = math.MinInt32
	L_Error Level = -1
	L_Info  Level = 0
	L_Debug Level = 1
)

var l = L_Info

func SetLevel(level Level) { l = level }

func (l Level) String() string {
	switch l {
	case L_Error: return Colorize(LightRed,"E")
	case L_Info: return Colorize(Cyan,"I")
	case L_Debug: return Colorize(LightYellow,"D")
	case L_Fatal: return Colorize(Red,"F")
	default: return "-"
	}
}

func Log(out io.Writer, format string, src string, level Level, msg string, args ...any) {
	res := Pair(args)
	strs := Map[[]any, string](
		func(a []any) string {
			a1 := String(fmt.Sprint(a[0]))
			if len(a) == 1 { return a1.Colorize(LightYellow).String() + ";" }
			a1 = a1.ToUpper().Colorize(LightYellow)
			a = a[1:]
			res := Map[any, string](func(a any) string {
				switch v := a.(type) {
				case string: return String(fmt.Sprintf("%q", v)).Colorize(Yellow).String()
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					return String(fmt.Sprint(v)).Colorize(Cyan).String()
				case bool: return String(fmt.Sprint(v)).Colorize(LightGreen).String()
				default: return String(fmt.Sprint(a)).Colorize(LightGreen).String()
				}
			})(a)
			return fmt.Sprintf("%s:<%s>;", a1, strings.Join(res, ", "))
		},
	)(res.Values())

	pr := ""

	if showLevel { pr += level.String() + " " }
	if showFile { pr += String(src).Dim().String() + "\t" }
	if showTime { pr += String(Time()).Colorize(Cyan).Dim().String() + " " }

	fmt.Fprintf(
		out,
		"%s%s %s\n",
		pr,
		String(msg),
		strings.Join(strs, " "),
	)
}

func caller(file string, line int) string {
	split := String(file).Split("/", "\\").Values()
	n := split[len(split)-1]
	return fmt.Sprintf("%s:%d", n, line)
}

func ErrFatal(v any, msg string, args ...any) {
	if v == nil { return }

	var errMsg string
	switch v := v.(type) {
	case error: errMsg = v.Error()
	case string: errMsg = v
	default: errMsg = fmt.Sprint(v)
	}

	_, file, line, _ := runtime.Caller(1)
	args = append([]any{"err", errMsg}, args...)
	Log(os.Stdout, "", caller(file, line), L_Error, msg, args...)
	os.Exit(1)
}

func Info(msg string, args ...any) {
	if l < L_Info { return }
	_, file, line, ok := runtime.Caller(1)
	if !ok { panic("Failed to get caller") }
	Log(os.Stdout, "", caller(file, line), L_Info, msg, args...)
}

func Error(msg string, args ...any) {
	if l < L_Error { return }
	_, file, line, ok := runtime.Caller(1)
	if !ok { panic("Failed to get caller") }
	Log(os.Stdout, "", caller(file, line), L_Error, msg, args...)
}

func Debug(msg string, args ...any) {
	if l < L_Debug { return }
	_, file, line, ok := runtime.Caller(1)
	if !ok { panic("Failed to get caller") }
	Log(os.Stdout, "", caller(file, line), L_Debug, msg, args...)
}

func Fatal(msg string, args ...any) {
	if l < L_Fatal { return }
	_, file, line, _ := runtime.Caller(1)
	Log(os.Stdout, "", caller(file, line), L_Debug, msg, args...)
	os.Exit(1)
}
