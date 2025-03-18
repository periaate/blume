package blume

import (
	"fmt"
	"github.com/periaate/blume/color"
	"github.com/periaate/blume/symbols"
)

func Log[A any](args ...any) func(A) {
	return func(a A) { fmt.Println(append(args, a)...) }
}
func Logs[A any](a A) { fmt.Println(a) }
func Logf[A any](format string, args ...any) func(A) {
	return func(arg A) { fmt.Printf(format, append(args, any(arg))...) }
}

func Through[A any](fn func(A)) func(A) A { return func(arg A) A { fn(arg); return arg } }

func HexToRGB(hex string) (int64, int64, int64) {
	hex = Del(Rgx("^#"))(hex)

	r := Parse[int64](hex[0:2], 16).Or(255)
	g := Parse[int64](hex[2:4], 16).Or(255)
	b := Parse[int64](hex[4:6], 16).Or(255)

	return r, g, b
}

func ColorFg(hex string) string {
	r, g, b := HexToRGB(hex)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func ColorBg(hex string) string {
	r, g, b := HexToRGB(hex)
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func Up(lines int) String { return String(fmt.Sprintf("\033[%dA", lines)) }
func Clean() String       { return String(fmt.Sprint("\r\033[K")) }

const P F = ""

type F string

func (f F) N() F { return f + "\n" }
func (f F) R() F { return f + "\r" }

func (f F) Up(lines ...int) F              { return f + F(Up(ToArray(lines).Get(0).Or(1))) }
func (f F) Clean() F                       { return f + F(Clean()) }
func (f F) S(args ...any) F                { return f + F(fmt.Sprint(args...)) }
func (f F) W() F                           { return f + F(" ") }
func (f F) F(format string, args ...any) F { return f + F(fmt.Sprintf(format, args...)) }

func (f F) Print(args ...any) F   { fmt.Printf("%s%s", f, fmt.Sprint(args...)); return f }
func (f F) Println(args ...any) F { fmt.Printf("%s%s", f, fmt.Sprintln(args...)); return f }
func (f F) Printf(format string, args ...any) F {
	fmt.Printf("%s%s", f, fmt.Sprintf(format, args...))
	return f
}

const Reset = "\033[0m"

func T[A any](ok bool, a A, b A) A {
	if ok {
		return a
	} else {
		return b
	}
}

func (f F) Color(hex string, args ...any) F { return f + F(ColorFg(hex)+fmt.Sprint(args...)+Reset) }

var spinChars = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func (f F) Spin(i int) F { return f + F(spinChars[i%len(spinChars)]) + " " }

func (f F) Info(args ...any) F      { return f.Color(color.Info, symbols.Info).W().S(args...) }
func (f F) Lock(args ...any) F      { return f.Color(color.Warning, symbols.Lock).W().S(args...) }
func (f F) Debug(args ...any) F     { return f.Color(color.Pending, symbols.Debug).W().S(args...) }
func (f F) Error(args ...any) F     { return f.Color(color.Error, symbols.Error).W().S(args...) }
func (f F) Success(args ...any) F   { return f.Color(color.Success, symbols.Success).W().S(args...) }
func (f F) Warning(args ...any) F   { return f.Color(color.Warning, symbols.Warning).W().S(args...) }
func (f F) Waiting(args ...any) F   { return f.Color(color.Waiting, symbols.Waiting).W().S(args...) }
func (f F) Question(args ...any) F  { return f.Color(color.Info, symbols.Question).W().S(args...) }
func (f F) Cancelled(args ...any) F { return f.Color(color.Error, symbols.Cancelled).W().S(args...) }
func (f F) InProgress(args ...any) F {
	return f.Color(color.Pending, symbols.InProgress).W().S(args...)
}

func (f F) Checkbox(done bool, args ...any) F {
	return T(done,
		f.Color(color.Success, symbols.CheckboxDone),
		f.Color(color.Warning, symbols.CheckboxEmpty),
	).S(args...)
}
