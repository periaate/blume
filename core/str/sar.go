package str

import (
	"fmt"
	"sort"
	"strings"
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
		for _, pattern := range match {
			if i+len(pattern) > len(str) {
				continue
			}

			if str[i:i+len(pattern)] != pattern {
				continue
			}

			if len(str[lastI:i]) != 0 {
				res = append(res, str[lastI:i])
			}

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

	if len(res) == 0 {
		return []string{str}
	}

	return res
}

func CaptureDelims(str string, delims ...string) (res []string, err error) {
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

	start := map[rune][2]rune{}
	for i := 0; i < len(delims); i += 2 {
		start[rune(delims[i][0])] = [2]rune{rune(delims[i+1][0]), rune(delims[i+1][0])}
	}

	var capturing bool
	var end rune
	sb := strings.Builder{}

	for _, r := range str {
		if capturing {
			if r == end {
				capturing = false
				sb.WriteRune(r)
				res = append(res, sb.String())
				sb.Reset()
				continue
			}
			sb.WriteRune(r)
			continue
		}

		if _, ok := start[r]; ok {
			capturing = true
			end = start[r][1]
			res = append(res, sb.String())
			sb.Reset()
			sb.WriteRune(r)
			continue
		}

		sb.WriteRune(r)
	}

	if sb.Len() > 0 {
		res = append(res, sb.String())
	}
	return
}
