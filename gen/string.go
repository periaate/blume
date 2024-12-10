package gen

import (
	"iter"
	"strconv"
	"strings"

	. "github.com/periaate/blume/core"
)

var _ = Zero[any]

func Sar[S ~string](args []string) (res []S) {
	res = make([]S, 0, len(args))
	for _, arg := range args {
		res = append(res, S(arg))
	}
	return
}

type SArray[S ~string] struct { value []S }

func (s SArray[S]) Values() []S { return s.value }

func (s SArray[S]) Iter() iter.Seq[S] {
	return func(yield func(S) bool) {
		for _, v := range s.value {
			if !yield(v) { break }
		}
	}
}

// func (s SArray[S]) From(arr []string) SAr[S] {
// 	res := make([]S, len(arr))
// 	for i, v := range arr {
// 		res[i] = S(v)
// 	}
// 	return SArray[S]{value: res}
// }

func (s SArray[S]) Array() []string {
	res := make([]string, len(s.value))
	for i, v := range s.value {
		res[i] = string(v)
	}
	return res
}

func (s SArray[S]) Join(sep string) S { return S(strings.Join(s.Array(), sep)) }

func ToInt(s string) Option[int] {
	i, err := strconv.Atoi(s)
	if err != nil { return None[int]() }
	return Some(i)

}
func ToInt8(s string) Option[int8] {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil { return None[int8]() }
	return Some(int8(i))
}
func ToInt16(s string) Option[int16] {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil { return None[int16]() }
	return Some(int16(i))
}
func ToInt32(s string) Option[int32] {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil { return None[int32]() }
	return Some(int32(i))
}
func ToInt64(s string) Option[int64] {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil { return None[int64]() }
	return Some(i)
}
func ToUint(s string) Option[uint] {
	i, err := strconv.ParseUint(s, 10, 0)
	if err != nil { return None[uint]() }
	return Some(uint(i))
}
func ToUint8(s string) Option[uint8] {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil { return None[uint8]() }
	return Some(uint8(i))
}
func ToUint16(s string) Option[uint16] {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil { return None[uint16]() }
	return Some(uint16(i))
}
func ToUint32(s string) Option[uint32] {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil { return None[uint32]() }
	return Some(uint32(i))
}
func ToUint64(s string) Option[uint64] {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil { return None[uint64]() }
	return Some(i)
}
func ToFloat32(s string) Option[float32] {
	i, err := strconv.ParseFloat(s, 32)
	if err != nil { return None[float32]() }
	return Some(float32(i))
}
func ToFloat64(s string) Option[float64] {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil { return None[float64]() }
	return Some(i)
}

func ToUpper(s string) string { return strings.ToUpper(s) }
func ToLower(s string) string { return strings.ToLower(s) }

func Trim(s string) string { return strings.Trim(s, " ") }
func TrimPrefix(prefix string, s string) string { return strings.TrimPrefix(s, prefix) }

func TrimSuffix(suffix string, s string) string { return strings.TrimSuffix(s, suffix) }
func TrimSpace(s string) string { return strings.TrimSpace(s) }
