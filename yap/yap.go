package yap

import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/periaate/blume/str"
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
	L_Fatal Level = math.MinInt32
	L_Error Level = -1
	L_Info  Level = 0
	L_Debug Level = 1
)

var l = L_Info

func SetLevel(level Level) { l = level }

func (l Level) String() string {
	switch l {
	case L_Error:
		return Colorize("E", LightRed)
	case L_Info:
		return Colorize("I", Cyan)
	case L_Debug:
		return Colorize("D", LightYellow)
	case L_Fatal:
		return Colorize("F", Red)
	default:
		return "-"
	}
}

func Pair[A any](arr []A) [][]A {
	pairs := [][]A{}
	var i int
	for i = 0; i < len(arr); i += 2 {
		cur := []A{}
		if i+1 <= len(arr) {
			cur = append(cur, arr[i])
		}
		if i+2 <= len(arr) {
			cur = append(cur, arr[i+1])
		}
		pairs = append(pairs, cur)
	}
	return pairs
}

func Log(out io.Writer, format string, src string, level Level, msg string, args ...any) {
	pairs := Pair(args)
	strs := []string{}
	for _, pair := range pairs {

		first := fmt.Sprint(pair[0])
		if len(pair) == 1 {
			strs = append(strs, Colorize(first, LightYellow)+";")
			continue
		}
		first = Colorize(first, LightYellow)
		pair := pair[1:]
		res := []string{}
		for i, val := range pair {
			switch v := val.(type) {
			case string:
				pair[i] = Colorize(fmt.Sprintf("%q", v), Yellow)
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
				pair[i] = Colorize(fmt.Sprint(v), Cyan)
			case bool:
				pair[i] = Colorize(fmt.Sprint(v), LightGreen)
			default:
				pair[i] = Colorize(fmt.Sprint(val), LightGreen)
			}
		}
		strs = append(strs, fmt.Sprintf("%s:<%s>;", first, strings.Join(res, ", ")))
	}

	pr := ""

	if showLevel {
		pr += level.String() + " "
	}
	if showFile {
		pr += Dim(src) + "\t"
	}
	if showTime {
		pr += Dim(Time()) + " "
	}

	fmt.Fprintf(
		out,
		"%s%s %s\n",
		pr,
		msg,
		strings.Join(strs, " "),
	)
}

func caller(file string, line int) string {
	split := str.Split(file, false, "/", "\\")
	n := split[len(split)-1]
	return fmt.Sprintf("%s:%d", n, line)
}

func ErrFatal(v any, msg string, args ...any) {
	if v == nil {
		return
	}

	var errMsg string
	switch v := v.(type) {
	case error:
		errMsg = v.Error()
	case string:
		errMsg = v
	default:
		errMsg = fmt.Sprint(v)
	}

	_, file, line, _ := runtime.Caller(1)
	args = append([]any{"err", errMsg}, args...)
	Log(os.Stdout, "", caller(file, line), L_Error, msg, args...)
	os.Exit(1)
}

func Info(msg string, args ...any) {
	if l < L_Info {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller")
	}
	Log(os.Stdout, "", caller(file, line), L_Info, msg, args...)
}

func Error(msg string, args ...any) {
	if l < L_Error {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller")
	}
	Log(os.Stdout, "", caller(file, line), L_Error, msg, args...)
}

func Debug(msg string, args ...any) {
	if l < L_Debug {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller")
	}
	Log(os.Stdout, "", caller(file, line), L_Debug, msg, args...)
}

func Fatal(msg string, args ...any) {
	if l < L_Fatal {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	Log(os.Stdout, "", caller(file, line), L_Debug, msg, args...)
	os.Exit(1)
}
