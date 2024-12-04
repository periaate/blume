package T

type Str interface {
	Contains(args ...string) bool
	HasPrefix(args ...string) bool
	HasSuffix(args ...string) bool
	ReplacePrefix(pats ...string) string
	ReplaceSuffix(pats ...string) string
	Replace(pats ...string) string
	ReplaceRegex(pat string, rep string) string
	Shift(count int) string
	Pop(count int) string
	Split(pats ...string) []string
}
