package gen

import (
	"strconv"
	"strings"

	. "github.com/periaate/blume/core"
)

func ToInt(s string) Option[int] {
	i, err := strconv.Atoi(s)
	if err != nil { return None[int]() }
	return Some(i)

}
func ToInt8(s string) Option[int8] {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil { return None[int8]() }
	return Some(int8(i))
}
func ToInt16(s string) Option[int16] {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil { return None[int16]() }
	return Some(int16(i))
}
func ToInt32(s string) Option[int32] {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil { return None[int32]() }
	return Some(int32(i))
}
func ToInt64(s string) Option[int64] {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil { return None[int64]() }
	return Some(i)
}
func ToUint(s string) Option[uint] {
	i, err := strconv.ParseUint(s, 10, 0)
	if err != nil { return None[uint]() }
	return Some(uint(i))
}
func ToUint8(s string) Option[uint8] {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil { return None[uint8]() }
	return Some(uint8(i))
}
func ToUint16(s string) Option[uint16] {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil { return None[uint16]() }
	return Some(uint16(i))
}
func ToUint32(s string) Option[uint32] {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil { return None[uint32]() }
	return Some(uint32(i))
}
func ToUint64(s string) Option[uint64] {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil { return None[uint64]() }
	return Some(i)
}
func ToFloat32(s string) Option[float32] {
	i, err := strconv.ParseFloat(s, 32)
	if err != nil { return None[float32]() }
	return Some(float32(i))
}
func ToFloat64(s string) Option[float64] {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil { return None[float64]() }
	return Some(i)
}

func ToUpper(s string) string { return strings.ToUpper(s) }
func ToLower(s string) string { return strings.ToLower(s) }

func Trim(s string) string { return strings.Trim(s, " ") }
func TrimPrefix(prefix string, s string) string { return strings.TrimPrefix(s, prefix) }
func TrimSuffix(suffix string, s string) string { return strings.TrimSuffix(s, suffix) }
func TrimSpace(s string) string { return strings.TrimSpace(s) }


const (
	reset = "\033[0m"

	Black        = 30
	Red          = 31
	Green        = 32
	Yellow       = 33
	Blue         = 34
	Magenta      = 35
	Cyan         = 36
	LightGray    = 37
	DarkGray     = 90
	LightRed     = 91
	LightGreen   = 92
	LightYellow  = 93
	LightBlue    = 94
	LightMagenta = 95
	LightCyan    = 96
	White        = 97
)

func Colorize[S ~string](colorCode int, s S) S {
	return "\033[" + S(strconv.Itoa(colorCode)) + "m" + s + "\033[0m"
}

func Dim[S ~string](s S) S { return Colorize(2, s) }
func Bold[S ~string](s S) S { return "\033[1m" + s + "\033[0m" }
