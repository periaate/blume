package str

import (
	"fmt"
	"sort"
	"strings"

	"github.com/periaate/blume/clog"
)

// SplitWithAll splits given string into an array, using all other `match` strings as
// delimiters. String is matched using the longest delimiter first.
// If no match strings are given, the original string is returned.
// If no matches are found, the original string is returned.
// Matched delimiters are not included in the result.
// If a found match would add a zero-length string to the result, it is ignored.
// Any consecutive matches are treated as one.
// If an empty match string is given (i.e. ""), every character is split.
// Keep argument determines whether the split part is included in the split part.
func SplitWithAll(str string, keep bool, match ...string) (res []string) {
	mult := 1
	if keep {
		mult = 0
	}

	if len(match) == 0 || len(str) == 0 {
		return []string{str}
	}

	sort.SliceStable(match, func(i, j int) bool {
		return len(match[i]) > len(match[j])
	})

	var lastI int

	for i := 0; i < len(str); i++ {
		clog.Debug("pr", "a", res, "v", string(str[i]))
		for _, pattern := range match {
			if i+len(pattern) > len(str) {
				continue
			}

			if str[i:i+len(pattern)] != pattern {
				continue
			}

			clog.Debug("diff", "_", string(str[lastI:i]))
			if len(str[lastI:i]) != 0 {
				res = append(res, str[lastI:i])
			}
			clog.Debug("pr", "b", res)

			lastI = i + len(pattern)*mult
			if len(pattern) != 0 {
				i += len(pattern) - 1
			}
			break
		}
	}

	if len(str[lastI:]) != 0 {
		res = append(res, str[lastI:])
	}

	clog.Debug("pr", "c", res)
	if len(res) == 0 {
		return []string{str}
	}

	clog.Debug("pr", "d", res)
	return res
}

func CaptureDelims(str string, keep bool, delims ...rune) (res []string, err error) {
	if len(str) == 0 {
		err = fmt.Errorf("empty string")
		return
	}

	if len(delims) == 0 {
		err = fmt.Errorf("no delimiters provided")
		return
	}

	if len(delims)%2 != 0 {
		err = fmt.Errorf("odd number of delimiters provided")
		return
	}

	start := map[rune]rune{}
	for i := 0; i < len(delims); i += 2 {
		start[delims[i]] = delims[i+1]
	}

	var capturing bool
	var end rune
	sb := strings.Builder{}

	for _, r := range str {

		if capturing {
			if r == end {
				capturing = false
				if keep {
					sb.WriteRune(r)
				}
				res = append(res, sb.String())
				sb.Reset()
				continue
			}
			sb.WriteRune(r)
			continue
		}

		if v, ok := start[r]; ok {
			capturing = true
			end = v
			res = append(res, sb.String())
			sb.Reset()
			if keep {
				sb.WriteRune(r)
			}
			continue
		}

		sb.WriteRune(r)
	}

	if sb.Len() > 0 {
		res = append(res, sb.String())
	}
	return
}

type Hierarchy struct {
	Arr []Hierarchy
	Str string
}
type Delimiter struct {
	Start string
	End   string
}

func EmbedDelims(sar []string, delims ...Delimiter) (Hierarchy, error) {
	car := make([]Hierarchy, len(sar))
	for i, s := range sar {
		car[i].Str = s
	}
	res, _ := embeds(car, delims)
	return res, nil
}

func embeds(car []Hierarchy, delims []Delimiter) (Hierarchy, int) {
	var res Hierarchy
	for i := 0; len(car) > i; i++ {
		v := car[i]
		matched := false
		for _, delim := range delims {
			switch v.Str {
			case delim.Start:
				r, k := embeds(car[i+1:], delims)
				i += k
				res.Arr = append(res.Arr, r)
				matched = true
			case delim.End:
				return res, i + 1
			}
			if matched {
				break
			}
		}
		if !matched {
			res.Arr = append(res.Arr, v)
		}
	}

	return res, 0
}
