package T

// Numeric is a type constraint that represents numeric types.
// Numeric does not include complex numbers.
type Numeric interface{ Unsigned | Signed | Float }

// Integer is a type constraint that represents integer types.
type Integer interface{ Signed | Unsigned }

// Float is a type constraint that represents floating-point types.
type Float interface{ ~float32 | ~float64 }

// Signed is a type constraint that represents signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a type constraint that represents unsigned integer types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
