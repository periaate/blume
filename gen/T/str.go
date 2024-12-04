package T

import "iter"

type SAr[S ~string] interface {
	From(arr []string) SAr[S]
	Array() []string
	Join(sep string) S
	Values() []S
	Iter() iter.Seq[S]
}

type Str[S ~string] interface {
	Contains(args ...string) bool
	HasPrefix(args ...string) bool
	HasSuffix(args ...string) bool
	ReplacePrefix(pats ...string) S
	ReplaceSuffix(pats ...string) S
	Replace(pats ...string) S
	ReplaceRegex(pat string, rep string) S
	Shift(count int) S
	Pop(count int) S
	Split(pats ...string) []S
	String() string
}
