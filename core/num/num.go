package num

type LNum interface {
	~uint64 | int64 | ~uintptr | float64
}

type Numeric interface {
	~float32 | ~float64 | Integer
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type UInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Abs[N Numeric](x N) (zero N) {
	if x < zero {
		return -x
	}
	return x
}

func Clamp[N Numeric](val, lower, upper N) (res N) {
	switch {
	case val >= upper:
		return upper
	case val <= lower:
		return lower
	default:
		return val
	}
}

func RangeClamp[I Integer](val, a, b I) I {
	if a == b {
		return b
	}

	if a > b {
		a = a ^ b
		b = a ^ b
		a = a ^ b
	}

	switch {
	case val >= b:
		return b
	case val <= a:
		return a
	default:
		return val
	}
}

func SameSign[N Numeric](a, b N) bool {
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}

func SmartClamp[I Integer](a, b I) I {
	switch {
	case b == 0 || a == 0:
		return 0
	case !SameSign(a, b):
		a += b
	}
	return RangeClamp(a, 0, b)
}
