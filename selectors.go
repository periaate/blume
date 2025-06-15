package blume

import (
	"regexp"
	"slices"
	"strings"
)

type Selector[A any] func(A) [][]int

func Pre(prefixes ...string) Selector[string] {
	return func(s string) (res [][]int) {
		for _, prefix := range prefixes {
			if strings.HasPrefix(s, prefix) {
				return [][]int{{0, len(prefix)}}
			}
		}
		return
	}
}

func Suf(suffixes ...string) Selector[string] {
	return func(s string) (res [][]int) {
		for _, suffix := range suffixes {
			if strings.HasSuffix(s, suffix) { return [][]int{ {len(s) - len(suffix), len(s) }} }
		}
		return
	}
}

func SelToPred[A any](selector Selector[A]) Pred[A] { return func(input A) bool { return len(selector(input)) > 0 } }

func Rgx(pattern string) Selector[string] {
	re := regexp.MustCompile(string(pattern))
	return func(s string) (res [][]int) { return re.FindAllStringIndex(string(s), -1) }
}

func Has[A any](selectors ...Selector[A]) func(input A) bool {
	pred := PredAnd(Map[Selector[A], Pred[A]](SelToPred[A])(selectors)...)
	return func(input A) bool { return slices.ContainsFunc(input, pred) }
}

func Del[A ~string](selectors ...Selector[A]) func(input A) A {
	return func(input A) A {
		for _, fn := range selectors {
			ranges := fn(input)
			if len(ranges) == 0 { continue }
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
	if len(pairs)%2 != 0 { panic("typeless generic function [Rep] needs to be given pairs of [`Selector[A]`, `A`].") }
	for i := 0; i < len(pairs); i += 2 {
		sel, ok := pairs[i].(Selector[A])
		if !ok { panic("typeless generic function [Rep] given invalid selector type.") }

		rep, ok := pairs[i+1].(A)
		if !ok { panic("typeless generic function [Rep] given invalid replacement type.") }

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
	if len(ranges) == 0 { return tar }
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

