package blume

import (
	"regexp"
	"sort"
	"strings"
)

type A[T any] = Array[T]

// Range returns a Selector which matches from the start, until the end.
// TODO: what to do with recursive ranges
// TODO: what to do with unfinished/incorrect ranges (starts, no end, end before start)
func Range(start, end Pred[S]) Selector[A[S]] {
	return func(s A[S]) [][]int {
		r := [][]int{}
		curr := []int{}
		var started bool
		for ind, value := range s {
			if started {
				if !end(value) {
					continue
				}
				curr = append(curr, ind)
				r = append(r, curr)
				curr = []int{}
				started = false
				continue
			}
			if start(value) {
				started = true
				curr = append(curr, ind)
			}
		}
		return r
	}
}

func RangeSel(start, end Selector[S]) Selector[A[S]] {
	return func(s A[S]) [][]int {
		r := [][]int{}
		curr := []int{}
		var started bool
		for ind, value := range s {
			if started {
				if len(end(value)) == 0 {
					continue
				}
				curr = append(curr, ind)
				r = append(r, curr)
				curr = []int{}
				started = false
				continue
			}
			if len(start(value)) > 1 {
				started = true
				curr = append(curr, ind)
			}
		}
		return r
	}
}

func Pattern[A any](selector Selector[A], actor func(A, [][]int) A) func(A) A {
	return func(value A) A {
		selected := selector(value)
		result := actor(value, selected)
		return result
	}
}

func JoinWith(delim string) func(arr []string, sel [][]int) []string {
	return func(arr []string, sel [][]int) []string {
		res := []string{}
		var last int
		for _, selection := range sel {
			res = append(res, arr[last:selection[0]]...)
			res = append(res, Join[string](delim)(arr[selection[0]:selection[1]+1]))
			last = selection[1] + 1
		}
		res = append(res, arr[last:]...)
		return res
	}
}

func Pre(prefixes ...string) Selector[string] {
	return func(s string) (res [][]int) {
		for _, prefix := range prefixes {
			if strings.HasPrefix(string(s), string(prefix)) {
				return [][]int{{0, len(prefix)}}
			}
		}
		return
	}
}

func Suf(suffixes ...string) Selector[string] {
	return func(s string) (res [][]int) {
		for _, suffix := range suffixes {
			if strings.HasSuffix(string(s), string(suffix)) {
				return [][]int{{len(s) - len(suffix), len(s)}}
			}
		}
		return
	}
}

func SelToPred[A any](selector Selector[A]) Pred[A] {
	return func(input A) bool { return len(selector(input)) > 0 }
}

func Rgx(pattern S) Selector[S] {
	re := regexp.MustCompile(string(pattern))
	return func(s S) (res [][]int) {
		return re.FindAllStringIndex(string(s), -1)
	}
}

func Has[A any](selectors ...Selector[A]) func(input A) bool {
	return func(input A) bool {
		for _, fn := range selectors {
			if len(fn(input)) > 0 {
				return true
			}
		}
		return false
	}
}

func Del[A ~string](selectors ...Selector[A]) func(input A) A {
	return func(input A) A {
		for _, fn := range selectors {
			ranges := fn(input)
			if len(ranges) == 0 {
				continue
			}
			input = ReplaceRanges(input, "", ranges)
		}
		return input
	}
}

func Rep[A ~string](pairs ...any) func(input A) A {
	replacers := []struct {
		sel Selector[A]
		rep A
	}{}
	if len(pairs)%2 != 0 {
		panic("typeless generic function [Rep] needs to be given pairs of [`Selector[A]`, `A`].")
	}
	for i := 0; i < len(pairs); i += 2 {
		sel, ok := pairs[i].(Selector[A])
		if !ok {
			panic("typeless generic function [Rep] given invalid selector type.")
		}
		rep, ok := pairs[i+1].(A)
		if !ok {
			panic("typeless generic function [Rep] given invalid replacement type.")
		}
		replacers = append(replacers, struct {
			sel Selector[A]
			rep A
		}{sel, rep})
	}

	return func(input A) A {
		for _, r := range replacers {
			input = ReplaceRanges(input, r.rep, r.sel(input))
		}
		return input
	}
}

func ReplaceRanges[S ~string](tar S, rep S, ranges [][]int) S {
	if len(ranges) == 0 {
		return tar
	}
	sortedRanges := make([][]int, len(ranges))
	for i, r := range ranges {
		sortedRanges[i] = make([]int, len(r))
		copy(sortedRanges[i], r)
	}

	for i := range len(sortedRanges) {
		for j := range len(sortedRanges) {
			if sortedRanges[i][0] < sortedRanges[j][0] {
				sortedRanges[i], sortedRanges[j] = sortedRanges[j], sortedRanges[i]
			}
		}
	}

	for _, r := range sortedRanges {
		start, end := r[0], r[1]
		if start < 0 || end > len(tar) || start > end {
			continue
		}
		tar = tar[:start] + rep + tar[end:]
	}
	return tar
}

// func Get[S ~string](selectors ...Selector[S]) func(input S) S {
// 	return func(input S) S {
// 		var result strings.Builder
// 		for _, fn := range selectors {
// 			ranges := fn(input)
// 			for _, r := range ranges {
// 				if len(r) >= 2 {
// 					result.WriteString(string(input[r[0]:r[1]]))
// 				}
// 			}
// 		}
// 		return result
// 	}
// }

func Nth[S ~string](n int, selectors ...Selector[S]) func(input S) S {
	return func(input S) S {
		var allRanges [][]int
		for _, fn := range selectors {
			ranges := fn(input)
			allRanges = append(allRanges, ranges...)
		}

		if len(allRanges) == 0 {
			return S("")
		}

		// Sort ranges by start position
		sort.Slice(allRanges, func(i, j int) bool {
			return allRanges[i][0] < allRanges[j][0]
		})

		// Handle negative index
		if n < 0 {
			n = len(allRanges) + n
		}

		// Check if index is in range
		if n < 0 || n >= len(allRanges) {
			return S("")
		}

		r := allRanges[n]
		if len(r) >= 2 {
			return S(input[r[0]:r[1]])
		}

		return S("")
	}
}
