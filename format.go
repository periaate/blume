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

// Turn these into consts at some point when you feel like manually transferring everything from hex to rgb
var (
	Info       String = P.Color(color.Info, symbols.Info).W()
	Lock       String = P.Color(color.Warning, symbols.Lock).W()
	Debug      String = P.Color(color.Pending, symbols.Debug).W()
	Error      String = P.Color(color.Error, symbols.Error).W()
	Success    String = P.Color(color.Success, symbols.Success).W()
	Warning    String = P.Color(color.Warning, symbols.Warning).W()
	Waiting    String = P.Color(color.Waiting, symbols.Waiting).W()
	Question   String = P.Color(color.Info, symbols.Question).W()
	Cancelled  String = P.Color(color.Error, symbols.Cancelled).W()
	InProgress String = P.Color(color.Pending, symbols.InProgress).W()
)

func HexToRGB(hex string) (int64, int64, int64) {
	hex = Del(Rgx[string]("^#"))(hex)

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

const Reset = "\033[0m"

func Up(lines int) String { return String(fmt.Sprintf("\033[%dA", lines)) }
func Clean() String       { return String(fmt.Sprint("\r\033[K")) }

const P String = ""

func (f String) N() String { return f + "\n" }
func (f String) R() String { return f + "\r" }

func (f String) Up(lines ...int) String { return f + String(Up(ToArray(lines).Get(0).Or(1))) }
func (f String) Clean() String          { return f + String(Clean()) }
func (f String) S(args ...any) String   { return f + String(fmt.Sprint(args...)) }
func (f String) F(format String, args ...any) String {
	return f + String(fmt.Sprintf(format.String(), args...))
}
func (f String) W() String { return f + String(" ") }

func (f String) Print(args ...any) String   { fmt.Printf("%s%s", f, fmt.Sprint(args...)); return f }
func (f String) Println(args ...any) String { fmt.Printf("%s%s", f, fmt.Sprintln(args...)); return f }
func (f String) Printf(format string, args ...any) String {
	fmt.Printf("%s%s", f, fmt.Sprintf(format, args...))
	return f
}

func (f String) Color(hex string, args ...any) String {
	return f + String(ColorFg(hex)+fmt.Sprint(args...)+Reset)
}

var spinChars = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func (f String) Spin(i int) String { return f + String(spinChars[i%len(spinChars)]) + " " }

func (f String) Info(args ...any) String { return f.Color(color.Info, symbols.Info).W().S(args...) }
func (f String) Lock(args ...any) String { return f.Color(color.Warning, symbols.Lock).W().S(args...) }
func (f String) Debug(args ...any) String {
	return f.Color(color.Pending, symbols.Debug).W().S(args...)
}
func (f String) Error(args ...any) String { return f.Color(color.Error, symbols.Error).W().S(args...) }
func (f String) Success(args ...any) String {
	return f.Color(color.Success, symbols.Success).W().S(args...)
}
func (f String) Warning(args ...any) String {
	return f.Color(color.Warning, symbols.Warning).W().S(args...)
}
func (f String) Waiting(args ...any) String {
	return f.Color(color.Waiting, symbols.Waiting).W().S(args...)
}
func (f String) Question(args ...any) String {
	return f.Color(color.Info, symbols.Question).W().S(args...)
}
func (f String) Cancelled(args ...any) String {
	return f.Color(color.Error, symbols.Cancelled).W().S(args...)
}
func (f String) InProgress(args ...any) String {
	return f.Color(color.Pending, symbols.InProgress).W().S(args...)
}

func (f String) Checkbox(done bool, args ...any) String {
	return T(done,
		f.Color(color.Success, symbols.CheckboxDone),
		f.Color(color.Warning, symbols.CheckboxEmpty),
	).S(args...)
}
func (s String) Green() String        { return String(color.Colorize(color.Green, string(s))) }
func (s String) LightGreen() String   { return String(color.Colorize(color.LightGreen, string(s))) }
func (s String) Yellow() String       { return String(color.Colorize(color.Yellow, string(s))) }
func (s String) LightYellow() String  { return String(color.Colorize(color.LightYellow, string(s))) }
func (s String) Red() String          { return color.Colorize(color.Red, s) }
func (s String) LightRed() String     { return color.Colorize(color.LightRed, s) }
func (s String) Blue() String         { return color.Colorize(color.Blue, s) }
func (s String) LightBlue() String    { return color.Colorize(color.LightBlue, s) }
func (s String) Cyan() String         { return color.Colorize(color.Cyan, s) }
func (s String) LightCyan() String    { return color.Colorize(color.LightCyan, s) }
func (s String) Magenta() String      { return color.Colorize(color.Magenta, s) }
func (s String) LightMagenta() String { return color.Colorize(color.LightMagenta, s) }
func (s String) White() String        { return color.Colorize(color.White, s) }
func (s String) Black() String        { return color.Colorize(color.Black, s) }
func (s String) Gray() String         { return color.Colorize(color.DarkGray, s) }
func (s String) LightGray() String    { return color.Colorize(color.LightGray, s) }
