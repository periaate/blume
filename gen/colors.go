package gen

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

func Colorize[S ~string](colorCode int, s S) S {
	return "\033[" + S(strconv.Itoa(colorCode)) + "m" + s + "\033[0m"
}

func Dim[S ~string](s S) S {
	return Colorize(2, s)
}

func Bold[S ~string](s S) S {
	return "\033[1m" + s + "\033[0m"
}
