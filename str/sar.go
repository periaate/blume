package str

import (
	"sort"
)

func SplitWithAll(str string, keep bool, match ...string) (res []string) {
	if len(match) == 0 || len(str) == 0 {
		return []string{str}
	}

	sort.SliceStable(match, func(i, j int) bool {
		return len(match[i]) > len(match[j])
	})

	var lastI int
	for i := 0; i < len(str); i++ {
		for _, pattern := range match {
			switch {
			case i+len(pattern) > len(str):
				continue
			case str[i:i+len(pattern)] != pattern:
				continue
			case len(str[lastI:i]) != 0:
				res = append(res, str[lastI:i])
			}

			lastI = i + len(pattern)
			if len(pattern) != 0 {
				if keep {
					res = append(res, str[i:len(pattern)+i])
				}
				i += len(pattern) - 1
			}
			break
		}
	}

	if len(str[lastI:]) != 0 {
		res = append(res, str[lastI:])
	}

	return res
}

type Hierarchy struct {
	Arr []Hierarchy
	Str string
}

func EmbedDelims(sar []string, delims ...[2]string) (Hierarchy, error) {
	car := make([]Hierarchy, len(sar))
	for i, s := range sar {
		car[i].Str = s
	}
	res, _ := embeds(car, delims)
	return res, nil
}

func embeds(car []Hierarchy, delims [][2]string) (Hierarchy, int) {
	var res Hierarchy
	for i := 0; len(car) > i; i++ {
		v := car[i]
		matched := false
		for _, delim := range delims {
			switch v.Str {
			case delim[0]:
				r, k := embeds(car[i+1:], delims)
				i += k
				res.Arr = append(res.Arr, r)
				matched = true
			case delim[1]:
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
