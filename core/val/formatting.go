package val

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/periaate/blume/core/num"
)

func HumanNumber(num string) string {
	ind := strings.Index(num, ".")
	var temp string
	if ind != -1 {
		temp = num[ind:]
		num = num[:ind]
	}

	for i := len(num) - 3; i > 0; i -= 3 {
		num = num[:i] + "," + num[i:]
	}

	return num + temp
}

func HumanNumb[N num.Numeric](n N) string {
	num := fmt.Sprintf("%v", n)
	ind := strings.Index(num, ".")
	var temp string
	if ind != -1 {
		temp = num[ind:]
		num = num[:ind]
	}

	for i := len(num) - 3; i > 0; i -= 3 {
		num = num[:i] + "," + num[i:]
	}

	return num + temp
}

// HumanizeBytes converts an integer byte value into a human-readable string.
// base of 0 means the input is in bytes, of 1 in kB, ...
func HumanizeBytes[T num.Integer](base int, val T, decimals int, asKiB bool) string {
	var unit float64 = 1000 // Use 1000 as base for KB, MB, GB...
	suffixes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	if asKiB {
		unit = 1024 // Use 1024 as base for KiB, MiB, GiB...
		suffixes = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	}
	suffixes = suffixes[base:]

	if val == 0 {
		return fmt.Sprintf("%.*f %s", decimals, 0.0, suffixes[0])
	}

	negative := val < 0
	val = num.Abs(val)

	size := float64(val)
	i := 0
	for size >= unit && i < len(suffixes)-1 {
		size /= unit
		i++
	}

	if negative {
		size = -size
	}

	return fmt.Sprintf("%.*f %s", decimals, size, suffixes[i])
}

func RelativeTime(s string) string {
	inp, err := time.Parse(time.DateTime, s)
	if err != nil {
		return ""
	}

	dur := time.Since(inp)
	days := dur.Hours() / 24
	switch {
	case days <= 31:
		if days < 1 {
			return "Today"
		}
		return strconv.Itoa(int(days)) + " days ago"
	case days <= 365:
		months := int(days / 30)
		if months == 1 {
			return "1 month ago"
		}
		return strconv.Itoa(months) + " months ago"
	default:
		years := int(days / 365)
		if years == 1 {
			return "1 year ago"
		}
		fmt.Println(days, years)
		return strconv.Itoa(years) + " years ago"
	}
}

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

func Color(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s", strconv.Itoa(colorCode), v)
}

func EndColor(v string) string { return fmt.Sprintf("%s%s", v, reset) }

func Colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

func FmtF(f string, args ...string) (res string, err error) {
	if strings.Count(f, `%s`) != len(args) {
		err = fmt.Errorf("number of args does not match number of format specifiers")
		return
	}
	var i int
	for strings.Contains(f, `%s`) {
		f = strings.Replace(f, `%s`, args[i], 1)
		i++
	}

	res = f
	for k, v := range escapeMap {
		res = strings.ReplaceAll(res, k, string(v))
	}

	return res, nil
}

var escapeMap = map[string]rune{
	"\\n": '\n',
	"\\t": '\t',
}
