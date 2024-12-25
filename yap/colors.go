package yap

import "strconv"

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

func Colorize(s string, colorCode int) string {
	return "\033[" + strconv.Itoa(colorCode) + "m" + s + "\033[0m"
}
func Dim(s string) string  { return Colorize(s, 2) }
func Bold(s string) string { return "\033[1m" + s + "\033[0m" }
