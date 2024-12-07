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
	String() string
	ToInt() Result[int]
	ToInt8() Result[int8]
	ToInt16() Result[int16]
	ToInt32() Result[int32]
	ToInt64() Result[int64]
	ToUint() Result[uint]
	ToUint8() Result[uint8]
	ToUint16() Result[uint16]
	ToUint32() Result[uint32]
	ToUint64() Result[uint64]
	ToFloat32() Result[float32]
	ToFloat64() Result[float64]
	Colorize(colorCode int) S
	ToUpper() S
	ToLower() S
	Trim() S
	TrimPrefix(prefix string) S
	TrimSuffix(suffix string) S
	TrimSpace() S

	Green() S
	Yellow() S
	Red() S
	Blue() S
	LightGreen() S
	LightYellow() S
	LightRed() S
	LightBlue() S
	Cyan() S
	LightCyan() S
	Magenta() S
	LightMagenta() S
	Gray() S
	LightGray() S
	White() S
	Black() S
	Dim() S
}
