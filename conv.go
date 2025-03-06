package blume

import "strconv"

func StoS[A, B ~string](value A) B   { return B(value) }
func ItoI[A, B Numeric](value A) B   { return B(value) }
func StoD[A ~string](value A) string { return string(value) }

func Parse[N Integer | Float](s string, args ...any) Option[N] {
	var a N
	var (
		bitSize int
		base    = Cast[int](ToArray(args).Get(0).Or(10)).Or(10)
	)

	switch any(a).(type) {
	case int8, uint8:
		bitSize = 8
	case int16, uint16:
		bitSize = 16
	case int32, uint32, float32:
		bitSize = 32
	case int64, uint64, float64:
		bitSize = 64
	default:
		return None[N]()
	}
	var value any
	var err error

	switch any(a).(type) {
	case int, int8, int16, int32, int64:
		value, err = strconv.ParseInt(s, base, bitSize)
	case uint, uint8, uint16, uint32, uint64:
		value, err = strconv.ParseUint(s, base, bitSize)
	case float32, float64:
		value, err = strconv.ParseFloat(s, bitSize)
	}
	if err == nil {
		return Cast[N](value)
	}
	return None[N]()
}

func ToInt[S ~string](s S) Option[int] {
	i, err := strconv.Atoi(string(s))
	return Either[int, bool]{Value: int(i), Other: err == nil}
}
func ToInt8[S ~string](s S) Option[int8] {
	i, err := strconv.ParseInt(string(s), 10, 8)
	return Either[int8, bool]{Value: int8(i), Other: err == nil}
}
func ToInt16[S ~string](s S) Option[int16] {
	i, err := strconv.ParseInt(string(s), 10, 16)
	return Either[int16, bool]{Value: int16(i), Other: err == nil}
}
func ToInt32[S ~string](s S) Option[int32] {
	i, err := strconv.ParseInt(string(s), 10, 32)
	return Either[int32, bool]{Value: int32(i), Other: err == nil}
}
func ToInt64[S ~string](s S) Option[int64] {
	i, err := strconv.ParseInt(string(s), 10, 64)
	return Either[int64, bool]{Value: int64(i), Other: err == nil}
}
func ToUint[S ~string](s S) Option[uint] {
	i, err := strconv.ParseUint(string(s), 10, 0)
	return Either[uint, bool]{Value: uint(i), Other: err == nil}
}
func ToUint8[S ~string](s S) Option[uint8] {
	i, err := strconv.ParseUint(string(s), 10, 8)
	return Either[uint8, bool]{Value: uint8(i), Other: err == nil}
}
func ToUint16[S ~string](s S) Option[uint16] {
	i, err := strconv.ParseUint(string(s), 10, 16)
	return Either[uint16, bool]{Value: uint16(i), Other: err == nil}
}
func ToUint32[S ~string](s S) Option[uint32] {
	i, err := strconv.ParseUint(string(s), 10, 32)
	return Either[uint32, bool]{Value: uint32(i), Other: err == nil}
}
func ToUint64[S ~string](s S) Option[uint64] {
	i, err := strconv.ParseUint(string(s), 10, 64)
	return Either[uint64, bool]{Value: uint64(i), Other: err == nil}
}
func ToFloat32[S ~string](s S) Option[float32] {
	i, err := strconv.ParseFloat(string(s), 32)
	return Either[float32, bool]{Value: float32(i), Other: err == nil}
}
func ToFloat64[S ~string](s S) Option[float64] {
	i, err := strconv.ParseFloat(string(s), 64)
	return Either[float64, bool]{Value: float64(i), Other: err == nil}
}
