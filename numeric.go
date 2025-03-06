package blume

type (
	Numeric interface{ Unsigned | Signed | Float }
	Integer interface{ Signed | Unsigned }
	Float   interface{ ~float32 | ~float64 }
	Signed  interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
)

func Abs[N Numeric](n N) (zero N) {
	if n < zero {
		return -n
	}
	return n
}

// Clamp returns a function which ensures that the input value is within the specified range.
func Clamp[N Numeric](lower, upper N) func(N) N {
	if lower > upper {
		lower, upper = upper, lower
	}

	return func(val N) N {
		switch {
		case val >= upper:
			return upper
		case val <= lower:
			return lower
		default:
			return val
		}
	}
}

// SameSign returns true if a and b have the same sign.
func SameSign[N Numeric](a, b N) bool { return (a > 0 && b > 0) || (a < 0 && b < 0) }
