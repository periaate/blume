package typ

import (
	"fmt"
	"strconv"
	"time"

	"github.com/periaate/blume/str"
)

var (
	ns = time.Nanosecond.Nanoseconds()
	us = time.Microsecond.Nanoseconds()
	ms = time.Millisecond.Nanoseconds()
	s  = time.Second.Nanoseconds()
	m  = time.Minute.Nanoseconds()
	h  = time.Hour.Nanoseconds()
	d  = 24 * h
	M  = 30 * d
	w  = 7 * d
	y  = 365 * d
)

var timeMap = map[string]int64{
	"ns": ns,
	"us": us,
	"ms": ms,
	"s":  s,
	"m":  m,
	"h":  h,
	"d":  d,
	"M":  M,
	"w":  w,
	"y":  y,
}

func TimeDaten(inp string) (res time.Duration, err error) {
	/*
		ns, us, ms, s, m, h, d, M, y
		any order
		support - for negative values
		values are relative from current time by default
		default to 0, i.e., d = current date
		1d = 1 day from current date
		-1d = 1 day before current date
		if first part is "abs", then values are absolute
		format can be defined after "fmt" keyword
		defaults are context dependent
		time does not specify seconds unless second is mentioned, accuracy increases with more values
		if time is not mentioned, then only date is returned
		"now" is alias for relative 0 time, no date
	*/

	res = time.Duration(0)
	if inp == "" {
		return
	}

	p := str.SplitWithAll(inp, false, " ")

	var parts []string

	for i := 0; i < len(p); i++ {
		part := p[i]

		if part == "" {
			continue
		}

		prts := str.SplitWithAll(part, true, "us", "ms", "s", "m", "h", "d", "w", "M", "y")
		parts = append(parts, prts...)
	}

	for i := 0; i < len(parts); i++ {
		part := parts[i]
		if part == "" {
			continue
		}
		neg := false
		if part[0] == '-' {
			part = part[1:]
			neg = true
		}

		if len(parts) <= i+1 {
			err = fmt.Errorf("missing unit for %s", part)
			return
		}

		var mult int64
		var ok bool
		if mult, ok = timeMap[parts[i+1]]; !ok {
			err = fmt.Errorf("invalid unit %s", parts[i+1])
			return
		}

		var arg int64

		arg, err = strconv.ParseInt(part, 10, 64)
		if err != nil {
			return
		}
		if neg {
			arg = -arg
		}

		res += time.Duration(arg) * time.Duration(mult)
		i++
	}

	return
}
