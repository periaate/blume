package blume

import (
	"regexp"
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
			if strings.HasSuffix(s, suffix) {
				return [][]int{{len(s)-len(suffix), len(s)}}
			}
		}
		return
	}
}

func SelToPred[T any](selector Selector[T]) Pred[T] { return func(input T) bool { return len(selector(input)) > 0 } }

func Rgx(pattern string) Selector[string] {
	re := regexp.MustCompile(string(pattern))
	return func(s string) (res [][]int) { return re.FindAllStringIndex(string(s), -1) }
}

func Has[T any](selectors ...Selector[T]) func(input T) bool { return PredAnd(Map[Selector[T], Pred[T]](SelToPred[T])(selectors)...) }

func Del(selectors ...Selector[string]) func(input string) string {
	return func(input string) string {
		for _, fn := range selectors {
			ranges := fn(input)
			if len(ranges) == 0 { continue }
			input = ReplaceRanges(input, "", ranges)
		}
		return input
	}
}

func Rep(pairs ...any) func(input string) string {
	replacers := []struct {
		sel Selector[string]
		rep string
	}{}
	if len(pairs)%2 != 0 { panic("typeless generic function [Rep] needs to be given pairs of [`Selector[string]`, `string`].") }
	for i := 0; i < len(pairs); i += 2 {
		sel, ok := pairs[i].(Selector[string])
		if !ok { panic("typeless generic function [Rep] given invalid selector type.") }

		rep, ok := pairs[i+1].(string)
		if !ok { panic("typeless generic function [Rep] given invalid replacement type.") }

		replacers = append(replacers, struct {
			sel Selector[string]
			rep string
		}{sel, rep})
	}

	return func(input string) string {
		for _, r := range replacers {
			input = ReplaceRanges(input, r.rep, r.sel(input))
		}
		return input
	}
}

func ReplaceRanges(tar string, rep string, ranges [][]int) string {
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

