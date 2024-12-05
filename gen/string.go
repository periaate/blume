package gen

import (
	"iter"
	"strconv"
	"strings"

	"github.com/periaate/blume/gen/T"
)

func Sar[S ~string](args []string) (res []S) {
	res = make([]S, 0, len(args))
	for _, arg := range args {
		res = append(res, S(arg))
	}
	return
}

type SArray[S ~string] struct {
	value []S
}

func (s SArray[S]) Values() []S {
	return s.value
}

func (s SArray[S]) Iter() iter.Seq[S] {
	return func(yield func(S) bool) {
		for _, v := range s.value {
			if !yield(v) {
				break
			}
		}
	}
}

func (s SArray[S]) From(arr []string) T.SAr[S] {
	res := make([]S, len(arr))
	for i, v := range arr {
		res[i] = S(v)
	}
	return SArray[S]{value: res}
}

func (s SArray[S]) Array() []string {
	res := make([]string, len(s.value))
	for i, v := range s.value {
		res[i] = string(v)
	}
	return res
}

func (s SArray[S]) Join(sep string) S {
	return S(strings.Join(s.Array(), sep))
}

func ToInt(s string) T.Result[int] { return T.Results(strconv.Atoi(s)) }
func ToInt8(s string) T.Result[int8] {
	i, err := strconv.ParseInt(s, 10, 8)
	return T.Results(int8(i), err)
}

func ToInt16(s string) T.Result[int16] {
	i, err := strconv.ParseInt(s, 10, 16)
	return T.Results(int16(i), err)
}

func ToInt32(s string) T.Result[int32] {
	i, err := strconv.ParseInt(s, 10, 32)
	return T.Results(int32(i), err)
}

func ToInt64(s string) T.Result[int64] {
	i, err := strconv.ParseInt(s, 10, 64)
	return T.Results(int64(i), err)
}

func ToUint(s string) T.Result[uint] {
	i, err := strconv.ParseUint(s, 10, 0)
	return T.Results(uint(i), err)
}

func ToUint8(s string) T.Result[uint8] {
	i, err := strconv.ParseUint(s, 10, 8)
	return T.Results(uint8(i), err)
}

func ToUint16(s string) T.Result[uint16] {
	i, err := strconv.ParseUint(s, 10, 16)
	return T.Results(uint16(i), err)
}

func ToUint32(s string) T.Result[uint32] {
	i, err := strconv.ParseUint(s, 10, 32)
	return T.Results(uint32(i), err)
}

func ToUint64(s string) T.Result[uint64] {
	i, err := strconv.ParseUint(s, 10, 64)
	return T.Results(uint64(i), err)
}

func ToFloat32(s string) T.Result[float32] {
	i, err := strconv.ParseFloat(s, 32)
	return T.Results(float32(i), err)
}

func ToFloat64(s string) T.Result[float64] {
	i, err := strconv.ParseFloat(s, 64)
	return T.Results(float64(i), err)
}

func ToUpper(s string) string { return strings.ToUpper(s) }
func ToLower(s string) string { return strings.ToLower(s) }

func Trim(s string) string { return strings.Trim(s, " ") }
func TrimPrefix(prefix string, s string) string {
	return strings.TrimPrefix(s, prefix)
}

func TrimSuffix(suffix string, s string) string {
	return strings.TrimSuffix(s, suffix)
}
func TrimSpace(s string) string { return strings.TrimSpace(s) }
