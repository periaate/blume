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

func Gt[N Numeric](arg N) Pred[N] { return func(n N) bool { return n > arg } }
func Ge[N Numeric](arg N) Pred[N] { return func(n N) bool { return n >= arg } }
func Lt[N Numeric](arg N) Pred[N] { return func(n N) bool { return n < arg } }
func Le[N Numeric](arg N) Pred[N] { return func(n N) bool { return n <= arg } }
func Eq[K comparable](arg K) Pred[K] { return func(n K) bool { return n == arg } }
func Ne[K comparable](arg K) Pred[K] { return func(n K) bool { return n != arg } }
