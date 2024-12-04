package clog

//
// import (
// 	"fmt"
// 	"io"
// 	"log/slog"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"
//
// 	"github.com/periaate/blume/gen"
// )
//
// var defLog = DefaultClog()
//
// func GetDefaultClog() *slog.Logger  { return defLog }
// func SetDefaultClog(l *slog.Logger) { defLog = l }
//
// const (
// 	LevelError = slog.LevelError
// 	LevelInfo  = slog.LevelInfo
// 	LevelWarn  = slog.LevelWarn
// 	LevelDebug = slog.LevelDebug
// )
//
// func SetLogLoggerLevel(lvl slog.Level) { defLog.Handler().(*ClogHandler).SetLogLoggerLevel(lvl) }
//
// type Logger interface {
// 	Error(msg string, args ...any)
// 	Info(msg string, args ...any)
// 	Warn(msg string, args ...any)
// 	Debug(msg string, args ...any)
// }
//
// // Dummy is a dummy logger that does nothing.
// type Dummy struct{}
//
// func (Dummy) Error(_ string, _ ...any) {}
// func (Dummy) Info(_ string, _ ...any)  {}
// func (Dummy) Warn(_ string, _ ...any)  {}
// func (Dummy) Debug(_ string, _ ...any) {}
//
// func DefaultClog() *slog.Logger { return NewClog(os.Stdout, LevelInfo, MaxLen(50)) }
//
// // NewClog creates a new clog logger with the given writer, level, and options.
// func NewClog(w io.Writer, lvl slog.Level, opts ...func(*ClogHandler)) *slog.Logger {
// 	h := New(w, lvl, nil, DefaultGetter)
//
// 	for _, opt := range opts {
// 		opt(h)
// 	}
//
// 	if h.Level == nil {
// 		h.Level = slog.LevelInfo
// 	}
//
// 	if h.St == nil {
// 		Style(NewStyles(nil))(h)
// 	}
//
// 	return slog.New(h)
// }
//
// func Style(st *Styles) func(*ClogHandler) { return func(h *ClogHandler) { h.St = st } }
// func MaxLen(l int) func(*ClogHandler)     { return func(h *ClogHandler) { h.MaxLen = l } }
//
// func (l *ClogHandler) SetLogLoggerLevel(lvl slog.Level) { l.Lvl = int(lvl) }
//
// type Styles struct {
// 	TimeStamp bool
// 	Delim     [2]string
// }
//
// func NewStyles(st *Styles) *Styles {
// 	if st == nil {
// 		st = &Styles{}
// 	}
// 	if st.Delim == [2]string{} {
// 		st.Delim = [2]string{
// 			":" + Color(LightYellow, "<"),
// 			EndColor(">") + "; ",
// 		}
// 	}
//
// 	return st
// }
//
// func New(out io.Writer, lvl slog.Level, st *Styles, getter func(slog.Value) string) *ClogHandler {
// 	h := &ClogHandler{
// 		St:      st,
// 		Level:   lvl,
// 		Mut:     &sync.Mutex{},
// 		Out:     out,
// 		indLens: make(map[int]int),
// 		Getter:  getter,
// 	}
// 	if h.Level == nil {
// 		h.Level = slog.LevelInfo
// 	}
// 	return h
// }
//
// // HumanizeNumber converts a number into a human-readable string.
// // `.` is used as a decimal separator.
// // `,` is used as a thousands separator.
// func HumanizeNumber[N gen.Numeric](n N) string {
// 	num := fmt.Sprintf("%v", n)
// 	ind := strings.Index(num, ".")
// 	var temp string
// 	if ind != -1 {
// 		temp = num[ind:]
// 		num = num[:ind]
// 	}
//
// 	for i := len(num) - 3; i > 0; i -= 3 {
// 		num = num[:i] + "," + num[i:]
// 	}
//
// 	return num + temp
// }
//
// // HumanizeBytes converts an integer byte value into a human-readable string.
// // base of 0 means the input is in bytes, of 1 in kB, 2 in MB, ...
// func HumanizeBytes[I gen.Integer](base int, val I, decimals int, asKiB bool) string {
// 	var unit float64 = 1000 // Use 1000 as base for KB, MB, GB...
// 	suffixes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
// 	if asKiB {
// 		unit = 1024 // Use 1024 as base for KiB, MiB, GiB...
// 		suffixes = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
// 	}
// 	suffixes = suffixes[base:]
//
// 	if val == 0 {
// 		return fmt.Sprintf("%.*f %s", decimals, 0.0, suffixes[0])
// 	}
//
// 	negative := val < 0
// 	val = gen.Abs(val)
//
// 	size := float64(val)
// 	i := 0
// 	for size >= unit && i < len(suffixes)-1 {
// 		size /= unit
// 		i++
// 	}
//
// 	if negative {
// 		size = -size
// 	}
//
// 	return fmt.Sprintf("%.*f %s", decimals, size, suffixes[i])
// }
//
// // RelativeTime returns a human-readable relative time string.
// func RelativeTime(inp time.Time) string {
// 	dur := time.Since(inp)
// 	days := dur.Hours() / 24
// 	switch {
// 	case days <= 31:
// 		if days < 1 {
// 			return "Today"
// 		}
// 		return strconv.Itoa(int(days)) + " days ago"
// 	case days <= 365:
// 		months := int(days / 30)
// 		if months == 1 {
// 			return "1 month ago"
// 		}
// 		return strconv.Itoa(months) + " months ago"
// 	default:
// 		years := int(days / 365)
// 		if years == 1 {
// 			return "1 year ago"
// 		}
// 		fmt.Println(days, years)
// 		return strconv.Itoa(years) + " years ago"
// 	}
// }
//
// const (
// 	reset = "\033[0m"
//
// 	Black        = 30
// 	Red          = 31
// 	Green        = 32
// 	Yellow       = 33
// 	Blue         = 34
// 	Magenta      = 35
// 	Cyan         = 36
// 	LightGray    = 37
// 	DarkGray     = 90
// 	LightRed     = 91
// 	LightGreen   = 92
// 	LightYellow  = 93
// 	LightBlue    = 94
// 	LightMagenta = 95
// 	LightCyan    = 96
// 	White        = 97
// )
//
// func Color(colorCode int, v string) string {
// 	return fmt.Sprintf("\033[%dm%s", colorCode, v)
// }
//
// func EndColor(v string) string { return fmt.Sprintf("%s%s", v, reset) }
//
// func Colorize(colorCode int, v string) string {
// 	return fmt.Sprintf("\033[%dm%s%s", colorCode, v, reset)
// }
