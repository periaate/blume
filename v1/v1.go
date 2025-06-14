package v1

type (
	A[T any] = Array[T]
	S = String
)

type Array[T any] struct { Value []T }

type String string

// Pipe attempts to pattern match between input-...-output to infer the valid resolution mechanism
// 
func Pipe[I, O any](i I, rest ...any) (res O) { return }

// func StartsWith[I, O any](i I, rest ...any) (res O)

type Type func(any) bool
func DefType[T any](a any) bool { return true }

type none any

var (
	None = DefType[none]
)


