package gen

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

func EmbedDelims(sar []string, delims ...[2]string) Tree[string] {
	car := make([]Tree[string], len(sar))
	for i, s := range sar {
		car[i].Value = s
	}
	res, _ := embeds(car, delims)
	return res
}

func embeds(car []Tree[string], delims [][2]string) (Tree[string], int) {
	var res Tree[string]
	for i := 0; len(car) > i; i++ {
		v := car[i]
		matched := false
		for _, delim := range delims {
			switch v.Value {
			case delim[0]:
				r, k := embeds(car[i+1:], delims)
				i += k
				res.Nodes = append(res.Nodes, r)
				matched = true
			case delim[1]:
				return res, i + 1
			}
			if matched {
				break
			}
		}
		if !matched {
			res.Nodes = append(res.Nodes, v)
		}
	}

	return res, 0
}
