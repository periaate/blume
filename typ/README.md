<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# typ

```go
import "github.com/periaate/blume/typ"
```

## Index

- [func AND\[N Unsigned\]\(a, b N\) N](<#AND>)
- [func Abs\[N Numeric\]\(n N\) \(zero N\)](<#Abs>)
- [func Clamp\[N Numeric\]\(lower, upper N\) func\(N\) N](<#Clamp>)
- [func HAS\[N Unsigned\]\(a, b N\) bool](<#HAS>)
- [func Invert\[K, V comparable\]\(m map\[K\]V\) \(res map\[V\]\[\]K\)](<#Invert>)
- [func InvertAny\[K, C comparable, V any\]\(fn func\(V\) C\) func\(map\[K\]V\) map\[C\]\[\]K](<#InvertAny>)
- [func Keys\[K comparable, V any\]\(m map\[K\]V\) \[\]K](<#Keys>)
- [func NOT\[N Unsigned\]\(a, b N\) N](<#NOT>)
- [func OR\[N Unsigned\]\(a, b N\) N](<#OR>)
- [func SameSign\[N Numeric\]\(a, b N\) bool](<#SameSign>)
- [func Vals\[K comparable, V any\]\(m map\[K\]V\) \[\]V](<#Vals>)
- [func XOR\[N Unsigned\]\(a, b N\) N](<#XOR>)
- [type Float](<#Float>)
- [type Integer](<#Integer>)
- [type Numeric](<#Numeric>)
- [type Signed](<#Signed>)
- [type Unsigned](<#Unsigned>)


<a name="AND"></a>
## func AND

```go
func AND[N Unsigned](a, b N) N
```

AND bitwise operation.

<a name="Abs"></a>
## func Abs

```go
func Abs[N Numeric](n N) (zero N)
```

Abs returns the absolute value of x.

<a name="Clamp"></a>
## func Clamp

```go
func Clamp[N Numeric](lower, upper N) func(N) N
```

Clamp returns a function which ensures that the input value is within the specified range.

<a name="HAS"></a>
## func HAS

```go
func HAS[N Unsigned](a, b N) bool
```

HAS checks if a bit is set.

<a name="Invert"></a>
## func Invert

```go
func Invert[K, V comparable](m map[K]V) (res map[V][]K)
```

Invert creates and inverted index of the given map.

<a name="InvertAny"></a>
## func InvertAny

```go
func InvertAny[K, C comparable, V any](fn func(V) C) func(map[K]V) map[C][]K
```

InvertAny creates and inverted index of the given map using a custom function.

<a name="Keys"></a>
## func Keys

```go
func Keys[K comparable, V any](m map[K]V) []K
```

Keys returns the keys of the given map.

<a name="NOT"></a>
## func NOT

```go
func NOT[N Unsigned](a, b N) N
```

NOT bitwise operation.

<a name="OR"></a>
## func OR

```go
func OR[N Unsigned](a, b N) N
```

OR bitwise operation.

<a name="SameSign"></a>
## func SameSign

```go
func SameSign[N Numeric](a, b N) bool
```

SameSign returns true if a and b have the same sign.

<a name="Vals"></a>
## func Vals

```go
func Vals[K comparable, V any](m map[K]V) []V
```

Vals returns the values of the given map.

<a name="XOR"></a>
## func XOR

```go
func XOR[N Unsigned](a, b N) N
```

XOR bitwise operation.

<a name="Float"></a>
## type Float

Float is a type constraint that represents floating\-point types.

```go
type Float interface {
    // contains filtered or unexported methods
}
```

<a name="Integer"></a>
## type Integer

Integer is a type constraint that represents integer types.

```go
type Integer interface {
    // contains filtered or unexported methods
}
```

<a name="Numeric"></a>
## type Numeric

Numeric is a type constraint that represents numeric types. Numeric does not include complex numbers.

```go
type Numeric interface {
    // contains filtered or unexported methods
}
```

<a name="Signed"></a>
## type Signed

Signed is a type constraint that represents signed integer types.

```go
type Signed interface {
    // contains filtered or unexported methods
}
```

<a name="Unsigned"></a>
## type Unsigned

Unsigned is a type constraint that represents unsigned integer types.

```go
type Unsigned interface {
    // contains filtered or unexported methods
}
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)