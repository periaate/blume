package blume


// func Type[T any]() reflect.Type {
// 	var zero T
// 	return reflect.TypeOf(zero)
// }

func Def[T any]() (def T) { return }

// func TypesMatch[A, B any]() bool { return Type[A]().AssignableTo(Type[B]()) }
// func TypesOfMatch(arg any, args ...any) bool {
// 	return pred.Every(reflect.TypeOf(arg).AssignableTo)(Map(reflect.TypeOf)(args))
// }

type Bool bool
type B = Bool
type Any any

func (b Bool) S() S { return If[S](b, "true", "false") }

type Type[T any] struct {}
// type TypePattern[T any] struct {}

// func (t Type[T]) Is(v Any) Bool {
// 	AsType()
// }
//
type TypeFn func(v any) bool

// type (
// 	slice any
// 	Slice = Type[slice]
// )
//
// func slice[T any](v T) bool {
// 	Type[T]{}
// }
//
// func Matcher[A, B any]() {}
//
// type M[A, B any] func()
//
// func Appender[I, O any](i I) (o O) {
//
// 	M[Slice, Type[O]]
// }
