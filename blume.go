package blume

import (
	"fmt"
	"reflect"
	"slices"
)

// All returns true if all arguments pass the [Predicate].
func All[T any](fns ...Pred[T]) Pred[[]T] {
	fn := PredAnd(fns...)
	return func(args []T) bool {
		for _, arg := range args {
			if !fn(arg) {
				return false
			}
		}
		return true
	}
}

type OverFn[I, O any] func([]I) []O
type Pred[T any] = func(T) bool

func Map[
	I any,
	O any,
	Fn func(I) O |
	func(any) O |
	func(...I) O |
	func(...any) O |
	func(I, ...I) O |
	func(I, ...any) O |
	func(any, ...any) O,
](arg Fn) func([]I) []O {
	var fn func(I) O
	switch fun := any(arg).(type) {
	case func(I) O: fn = fun
	case func(any) O: fn = func(i I) O { return fun(i) }
	case func(...I) O: fn = func(i I) O { return fun(i) }
	case func(...any) O: fn = func(i I) O { return fun(i) }
	case func(I, ...I) O: fn = func(i I) O { return fun(i) }
	case func(I, ...any) O: fn = func(i I) O { return fun(i) }
	case func(any, ...any) O: fn = func(i I) O { return fun(i) }

	// this branch is only reachable via calling Map with reflection
	default : panic(fmt.Sprintf("impossible invariant passed to Map: %s", reflect.TypeOf(arg).Name()))
	}

	return func(args []I) (res []O) {
		res = make([]O, 0, len(args))
		for _, arg := range args { res = append(res, fn(arg)) }
		return res
	}
}


func FilterMap[
	I any,
	O any,

	// Filter maps
	Fn func(I) Option[O] |
	func(...I) Option[O] |
	func(any) Option[O] |
	func(...any) Option[O] |
	func(I, ...I) Option[O] |
	func(I, ...any) Option[O] |
	func(any, ...any) Option[O] |
	func(I) Result[O] |
	func(...I) Result[O] |
	func(any) Result[O] |
	func(...any) Result[O] |
	func(I, ...I) Result[O] |
	func(I, ...any) Result[O] |
	func(any, ...any) Result[O] |

	// Filter maps native
	func(I) (O, bool) |
	func(...I) (O, bool) |
	func(any) (O, bool) |
	func(...any) (O, bool) |
	func(I, ...I) (O, bool) |
	func(I, ...any) (O, bool) |
	func(any, ...any) (O, bool) |
	func(I) (O, error) |
	func(...I) (O, error) |
	func(any) (O, error) |
	func(...any) (O, error) |
	func(I, ...I) (O, error) |
	func(I, ...any) (O, error) |
	func(any, ...any) (O, error),
](arg Fn) func([]I) []O {
	var fn func(I) (res O, ok bool)
	
	switch fun := any(arg).(type) {
	// Filter maps
	case func(I) Option[O]: fn = func(i I) (res O, ok bool) { return fun(i).Unwrap() }
	case func(...I) Option[O]: fn = func(i I) (res O, ok bool) { return fun(i).Unwrap() }
	case func(any) Option[O]: fn = func(i I) (res O, ok bool) { return fun(i).Unwrap() }
	case func(...any) Option[O]: fn = func(i I) (res O, ok bool) { return fun(i).Unwrap() }
	case func(I, ...I) Option[O]: fn = func(i I) (res O, ok bool) { return fun(i).Unwrap() }
	case func(I, ...any) Option[O]: fn = func(i I) (res O, ok bool) { return fun(i).Unwrap() }
	case func(any, ...any) Option[O]: fn = func(i I) (res O, ok bool) { return fun(i).Unwrap() }
	case func(I) Result[O]: fn = func(i I) (res O, ok bool) { v, err := fun(i).Unwrap(); if err != nil { return }; return v, true }
	case func(...I) Result[O]: fn = func(i I) (res O, ok bool) { v, err := fun(i).Unwrap(); if err != nil { return }; return v, true }
	case func(any) Result[O]: fn = func(i I) (res O, ok bool) { v, err := fun(i).Unwrap(); if err != nil { return }; return v, true }
	case func(...any) Result[O]: fn = func(i I) (res O, ok bool) { v, err := fun(i).Unwrap(); if err != nil { return }; return v, true }
	case func(I, ...I) Result[O]: fn = func(i I) (res O, ok bool) { v, err := fun(i).Unwrap(); if err != nil { return }; return v, true }
	case func(I, ...any) Result[O]: fn = func(i I) (res O, ok bool) { v, err := fun(i).Unwrap(); if err != nil { return }; return v, true }
	case func(any, ...any) Result[O]: fn = func(i I) (res O, ok bool) { v, err := fun(i).Unwrap(); if err != nil { return }; return v, true }

	// Filter maps native
	case func(I) (O, bool): fn = func(i I) (res O, ok bool) { return fun(i) }
	case func(...I) (O, bool): fn = func(i I) (res O, ok bool) { return fun(i) }
	case func(any) (O, bool): fn = func(i I) (res O, ok bool) { return fun(i) }
	case func(...any) (O, bool): fn = func(i I) (res O, ok bool) { return fun(i) }
	case func(I, ...I) (O, bool): fn = func(i I) (res O, ok bool) { return fun(i) }
	case func(I, ...any) (O, bool): fn = func(i I) (res O, ok bool) { return fun(i) }
	case func(any, ...any) (O, bool): fn = func(i I) (res O, ok bool) { return fun(i) }
	case func(I) (O, error): fn = func(i I) (res O, ok bool) { v, err := fun(i); if err != nil { return }; return v, true }
	case func(...I) (O, error): fn = func(i I) (res O, ok bool) { v, err := fun(i); if err != nil { return }; return v, true }
	case func(any) (O, error): fn = func(i I) (res O, ok bool) { v, err := fun(i); if err != nil { return }; return v, true }
	case func(...any) (O, error): fn = func(i I) (res O, ok bool) { v, err := fun(i); if err != nil { return }; return v, true }
	case func(I, ...I) (O, error): fn = func(i I) (res O, ok bool) { v, err := fun(i); if err != nil { return }; return v, true }
	case func(I, ...any) (O, error): fn = func(i I) (res O, ok bool) { v, err := fun(i); if err != nil { return }; return v, true }
	case func(any, ...any) (O, error): fn = func(i I) (res O, ok bool) { v, err := fun(i); if err != nil { return }; return v, true }

	// this branch is only reachable via calling FilterMap with reflection
	default : panic(fmt.Sprintf("impossible invariant passed to FilterMap: %s", reflect.TypeOf(arg).Name()))
	}

	return func(arr []I) []O {
		res := []O{}
		for _, val := range arr {
			if val, ok := fn(val); ok {
				res = append(res, val)
			}
		}
		return res
	}
}

func Filter[
	I any,
	Fn func(I) bool |
	func(...I) bool |
	func(any) bool |
	func(...any) bool |
	func(I, ...I) bool |
	func(I, ...any) bool |
	func(any, ...any) bool,
] (args ...Fn)       func([]I) []I {
	fns := []func(I) bool{}
	for _, arg := range args {
		switch fn := any(arg).(type) {
		case func(I) bool: fns = append(fns, func(i I) bool { return fn(i) })
		case func(...I) bool: fns = append(fns, func(i I) bool { return fn(i) })
		case func(any) bool: fns = append(fns, func(i I) bool { return fn(i) })
		case func(...any) bool: fns = append(fns, func(i I) bool { return fn(i) })
		case func(I, ...I) bool: fns = append(fns, func(i I) bool { return fn(i) })
		case func(I, ...any) bool: fns = append(fns, func(i I) bool { return fn(i) })
		case func(any, ...any) bool: fns = append(fns, func(i I) bool { return fn(i) })

		// this branch is only reachable via calling Filter with reflection
		default : panic(fmt.Sprintf("impossible invariant passed to Filter: %s", reflect.TypeOf(arg).Name()))
		}
	}
	fn := PredAnd(fns...)
	return func(args []I) (res []I) {
		for _, arg := range args {
			if fn(arg) {
				res = append(res, arg)
			}
		}
		return res
	}
}

func Each[I any](fn func(I)) func([]I) []I {
	return func(arr []I) []I {
		for _, value := range arr {
			fn(value)
		}
		return arr
	}
}

// Through generalizes array operations by type safe function type
// Through uses runtime type checking during construction, meaning
// once it has been called, there is no additional overhead on
// the resulting function.
func Through[
	I any,
	Fn func(I) |
	func(...I) |
	func(any) |
	func(...any) |
	func(I, ...I) |
	func(I, ...any) |
	func(any, ...any) |
	func(...any) (int, error) | // e.g., fmt.Print... family of functions.

	// Maps
	func(I) I |
	func(...I) I |
	func(any) I |
	func(...any) I |
	func(I, ...I) I |
	func(I, ...any) I |
	func(any, ...any) I |

	// Filters
	func(I) bool |
	func(...I) bool |
	func(any) bool |
	func(...any) bool |
	func(I, ...I) bool |
	func(I, ...any) bool |
	func(any, ...any) bool |

	// Filter maps
	func(I) Option[I] |
	func(...I) Option[I] |
	func(any) Option[I] |
	func(...any) Option[I] |
	func(I, ...I) Option[I] |
	func(I, ...any) Option[I] |
	func(any, ...any) Option[I] |
	func(I) Result[I] |
	func(...I) Result[I] |
	func(any) Result[I] |
	func(...any) Result[I] |
	func(I, ...I) Result[I] |
	func(I, ...any) Result[I] |
	func(any, ...any) Result[I] |

	// Filter maps native
	func(I) (I, bool) |
	func(...I) (I, bool) |
	func(any) (I, bool) |
	func(...any) (I, bool) |
	func(I, ...I) (I, bool) |
	func(I, ...any) (I, bool) |
	func(any, ...any) (I, bool) |
	func(I) (I, error) |
	func(...I) (I, error) |
	func(any) (I, error) |
	func(...any) (I, error) |
	func(I, ...I) (I, error) |
	func(I, ...any) (I, error) |
	func(any, ...any) (I, error),
] (arg Fn) func([]I) []I {
	switch fun := any(arg).(type) {
	// loops
	case func(I): return Each(fun)
	case func(any): return Each(func(i I) { fun(i) })
	case func(...I): return Each(func(i I) { fun(i) })
	case func(...any): return Each(func(i I) { fun(i) })
	case func(I, ...I): return Each(func(i I) { fun(i) })
	case func(I, ...any): return Each(func(i I) { fun(i) })
	case func(any, ...any): return Each(func(i I) { fun(i) })
	case func(...any) (int, error): return Each(func(i I) { fun(i) })

	// Maps
	case func(I) I: return Over[I, I](fun)
	case func(...I) I: return Over[I, I](fun)
	case func(any) I: return Over[I, I](fun)
	case func(...any) I: return Over[I, I](fun)
	case func(I, ...I) I: return Over[I, I](fun)
	case func(I, ...any) I: return Over[I, I](fun)
	case func(any, ...any) I: return Over[I, I](fun)

	// Filters
	case func(I) bool: return Filter[I](fun)
	case func(...I) bool: return Filter[I](fun)
	case func(any) bool: return Filter[I](fun)
	case func(...any) bool: return Filter[I](fun)
	case func(I, ...I) bool: return Filter[I](fun)
	case func(I, ...any) bool: return Filter[I](fun)
	case func(any, ...any) bool: return Filter[I](fun)

	// Filter maps
	case func(I) Option[I]: return FilterMap[I, I](fun)
	case func(...I) Option[I]: return FilterMap[I, I](fun)
	case func(any) Option[I]: return FilterMap[I, I](fun)
	case func(...any) Option[I]: return FilterMap[I, I](fun)
	case func(I, ...I) Option[I]: return FilterMap[I, I](fun)
	case func(I, ...any) Option[I]: return FilterMap[I, I](fun)
	case func(any, ...any) Option[I]: return FilterMap[I, I](fun)
	case func(I) Result[I]: return FilterMap[I, I](fun)
	case func(...I) Result[I]: return FilterMap[I, I](fun)
	case func(any) Result[I]: return FilterMap[I, I](fun)
	case func(...any) Result[I]: return FilterMap[I, I](fun)
	case func(I, ...I) Result[I]: return FilterMap[I, I](fun)
	case func(I, ...any) Result[I]: return FilterMap[I, I](fun)
	case func(any, ...any) Result[I]: return FilterMap[I, I](fun)

	// Filter maps native
	case func(I) (I, bool): return FilterMap[I, I](fun)
	case func(...I) (I, bool): return FilterMap[I, I](fun)
	case func(any) (I, bool): return FilterMap[I, I](fun)
	case func(...any) (I, bool): return FilterMap[I, I](fun)
	case func(I, ...I) (I, bool): return FilterMap[I, I](fun)
	case func(I, ...any) (I, bool): return FilterMap[I, I](fun)
	case func(any, ...any) (I, bool): return FilterMap[I, I](fun)
	case func(I) (I, error): return FilterMap[I, I](fun)
	case func(...I) (I, error): return FilterMap[I, I](fun)
	case func(any) (I, error): return FilterMap[I, I](fun)
	case func(...any) (I, error): return FilterMap[I, I](fun)
	case func(I, ...I) (I, error): return FilterMap[I, I](fun)
	case func(I, ...any) (I, error): return FilterMap[I, I](fun)
	case func(any, ...any) (I, error): return FilterMap[I, I](fun)

	// this branch is only reachable via calling Through with reflection
	default : panic(fmt.Sprintf("impossible invariant passed to Through: %s", reflect.TypeOf(arg).Name()))
	}
}


type FilterMaps[I, O any] interface{
	func(I) O |
	func(...I) O |
	func(any) O |
	func(...any) O |
	func(I, ...I) O |
	func(I, ...any) O |
	func(any, ...any) O |

	// Filter maps
	func(I) Option[O] |
	func(...I) Option[O] |
	func(any) Option[O] |
	func(...any) Option[O] |
	func(I, ...I) Option[O] |
	func(I, ...any) Option[O] |
	func(any, ...any) Option[O] |

	func(I) Result[O] |
	func(...I) Result[O] |
	func(any) Result[O] |
	func(...any) Result[O] |
	func(I, ...I) Result[O] |
	func(I, ...any) Result[O] |
	func(any, ...any) Result[O] |

	// Filter maps native
	func(I) (O, bool) |
	func(...I) (O, bool) |
	func(any) (O, bool) |
	func(...any) (O, bool) |
	func(I, ...I) (O, bool) |
	func(I, ...any) (O, bool) |
	func(any, ...any) (O, bool) |

	func(I) (O, error) |
	func(...I) (O, error) |
	func(any) (O, error) |
	func(...any) (O, error) |
	func(I, ...I) (O, error) |
	func(I, ...any) (O, error) |
	func(any, ...any) (O, error)
}

// Over generalizes array operations by type safe function type
// Over uses runtime type checking during construction, meaning
// once it has been called, there is no additional overhead on
// the resulting function.
func Over[
	I any,
	O any,
	Fn func(I) O |
	func(...I) O |
	func(any) O |
	func(...any) O |
	func(I, ...I) O |
	func(I, ...any) O |
	func(any, ...any) O |

	// Filter maps
	func(I) Option[O] |
	func(...I) Option[O] |
	func(any) Option[O] |
	func(...any) Option[O] |
	func(I, ...I) Option[O] |
	func(I, ...any) Option[O] |
	func(any, ...any) Option[O] |
	func(I) Result[O] |
	func(...I) Result[O] |
	func(any) Result[O] |
	func(...any) Result[O] |
	func(I, ...I) Result[O] |
	func(I, ...any) Result[O] |
	func(any, ...any) Result[O] |

	// Filter maps native
	func(I) (O, bool) |
	func(...I) (O, bool) |
	func(any) (O, bool) |
	func(...any) (O, bool) |
	func(I, ...I) (O, bool) |
	func(I, ...any) (O, bool) |
	func(any, ...any) (O, bool) |
	func(I) (O, error) |
	func(...I) (O, error) |
	func(any) (O, error) |
	func(...any) (O, error) |
	func(I, ...I) (O, error) |
	func(I, ...any) (O, error) |
	func(any, ...any) (O, error),
](arg Fn) (res func([]I) []O) {
	switch fn := any(arg).(type) {
	case func(I) O: return Map[I, O](fn)
	case func(...I) O: return Map[I, O](fn)
	case func(any) O: return Map[I, O](fn)
	case func(...any) O: return Map[I, O](fn)
	case func(I, ...I) O: return Map[I, O](fn)
	case func(I, ...any) O: return Map[I, O](fn)
	case func(any, ...any) O: return Map[I, O](fn)

	// Filter maps
	case func(I) Option[O]: return FilterMap[I, O](fn)
	case func(...I) Option[O]: return FilterMap[I, O](fn)
	case func(any) Option[O]: return FilterMap[I, O](fn)
	case func(...any) Option[O]: return FilterMap[I, O](fn)
	case func(I, ...I) Option[O]: return FilterMap[I, O](fn)
	case func(I, ...any) Option[O]: return FilterMap[I, O](fn)
	case func(any, ...any) Option[O]: return FilterMap[I, O](fn)
	case func(I) Result[O]: return FilterMap[I, O](fn)
	case func(...I) Result[O]: return FilterMap[I, O](fn)
	case func(any) Result[O]: return FilterMap[I, O](fn)
	case func(...any) Result[O]: return FilterMap[I, O](fn)
	case func(I, ...I) Result[O]: return FilterMap[I, O](fn)
	case func(I, ...any) Result[O]: return FilterMap[I, O](fn)
	case func(any, ...any) Result[O]: return FilterMap[I, O](fn)

	// Filter maps native
	case func(I) (O, bool): return FilterMap[I, O](fn)
	case func(...I) (O, bool): return FilterMap[I, O](fn)
	case func(any) (O, bool): return FilterMap[I, O](fn)
	case func(...any) (O, bool): return FilterMap[I, O](fn)
	case func(I, ...I) (O, bool): return FilterMap[I, O](fn)
	case func(I, ...any) (O, bool): return FilterMap[I, O](fn)
	case func(any, ...any) (O, bool): return FilterMap[I, O](fn)
	case func(I) (O, error): return FilterMap[I, O](fn)
	case func(...I) (O, error): return FilterMap[I, O](fn)
	case func(any) (O, error): return FilterMap[I, O](fn)
	case func(...any) (O, error): return FilterMap[I, O](fn)
	case func(I, ...I) (O, error): return FilterMap[I, O](fn)
	case func(I, ...any) (O, error): return FilterMap[I, O](fn)
	case func(any, ...any) (O, error): return FilterMap[I, O](fn)


	// this branch is only reachable via calling Over with reflection
	default : panic(fmt.Sprintf("impossible invariant passed to Over: %s", reflect.TypeOf(arg).Name()))
	}
}

// ====== Through Type Variants =======
// Loops
var _ = Through[int](func(_ int)           { return }) // func(I)
var _ = Through[int](func(_ any)           { return }) // func(any)
var _ = Through[int](func(_ ...int)        { return }) // func(...I)
var _ = Through[int](func(_ ...any)        { return }) // func(...any)
var _ = Through[int](func(_ int, _ ...int) { return }) // func(I, ...I)
var _ = Through[int](func(_ int, _ ...any) { return }) // func(I, ...any)
var _ = Through[int](func(_ any, _ ...any) { return }) // func(any, ...any)
var _ = Through[int](fmt.Print)                        // func(...any) (int, error)

// Maps
var _ = Through[int](func(int) (res int) { return })
var _ = Through[int](func(...int) (res int) { return })
var _ = Through[int](func(any) (res int) { return })
var _ = Through[int](func(...any) (res int) { return })
var _ = Through[int](func(int, ...int) (res int) { return })
var _ = Through[int](func(int, ...any) (res int) { return })
var _ = Through[int](func(any, ...any) (res int) { return })

// Filters
var _ = Through[int](func(int) (res bool) { return })
var _ = Through[int](func(...int) (res bool) { return })
var _ = Through[int](func(any) (res bool) { return })
var _ = Through[int](func(...any) (res bool) { return })
var _ = Through[int](func(int, ...int) (res bool) { return })
var _ = Through[int](func(int, ...any) (res bool) { return })
var _ = Through[int](func(any, ...any) (res bool) { return })

	// Filter maps
var _ = Through[int](func(int) (res Option[int]) { return })
var _ = Through[int](func(...int) (res Option[int]) { return })
var _ = Through[int](func(any) (res Option[int]) { return })
var _ = Through[int](func(...any) (res Option[int]) { return })
var _ = Through[int](func(int, ...int) (res Option[int]) { return })
var _ = Through[int](func(int, ...any) (res Option[int]) { return })
var _ = Through[int](func(any, ...any) (res Option[int]) { return })
var _ = Through[int](func(int) (res Result[int]) { return })
var _ = Through[int](func(...int) (res Result[int]) { return })
var _ = Through[int](func(any) (res Result[int]) { return })
var _ = Through[int](func(...any) (res Result[int]) { return })
var _ = Through[int](func(int, ...int) (res Result[int]) { return })
var _ = Through[int](func(int, ...any) (res Result[int]) { return })
var _ = Through[int](func(any, ...any) (res Result[int]) { return })

// Filter maps native
var _ = Through[int](func(int) (res int, v bool) { return })
var _ = Through[int](func(...int) (res int, v bool) { return })
var _ = Through[int](func(any) (res int, v bool) { return })
var _ = Through[int](func(...any) (res int, v bool) { return })
var _ = Through[int](func(int, ...int) (res int, v bool) { return })
var _ = Through[int](func(int, ...any) (res int, v bool) { return })
var _ = Through[int](func(any, ...any) (res int, v bool) { return })
var _ = Through[int](func(int) (res int, v error) { return })
var _ = Through[int](func(...int) (res int, v error) { return })
var _ = Through[int](func(any) (res int, v error) { return })
var _ = Through[int](func(...any) (res int, v error) { return })
var _ = Through[int](func(int, ...int) (res int, v error) { return })
var _ = Through[int](func(int, ...any) (res int, v error) { return })
var _ = Through[int](func(any, ...any) (res int, v error) { return })

// ====== Over Type Variants =======
// Maps
var _ = Over[int, int64](func(int) (res int64) { return })
var _ = Over[int, int64](func(...int) (res int64) { return })
var _ = Over[int, int64](func(any) (res int64) { return })
var _ = Over[int, int64](func(...any) (res int64) { return })
var _ = Over[int, int64](func(int, ...int) (res int64) { return })
var _ = Over[int, int64](func(int, ...any) (res int64) { return })
var _ = Over[int, int64](func(any, ...any) (res int64) { return })

// Filter maps
var _ = Over[int, int64](func(int) (res Option[int64]) { return })
var _ = Over[int, int64](func(...int) (res Option[int64]) { return })
var _ = Over[int, int64](func(any) (res Option[int64]) { return })
var _ = Over[int, int64](func(...any) (res Option[int64]) { return })
var _ = Over[int, int64](func(int, ...int) (res Option[int64]) { return })
var _ = Over[int, int64](func(int, ...any) (res Option[int64]) { return })
var _ = Over[int, int64](func(any, ...any) (res Option[int64]) { return })
var _ = Over[int, int64](func(int) (res Result[int64]) { return })
var _ = Over[int, int64](func(...int) (res Result[int64]) { return })
var _ = Over[int, int64](func(any) (res Result[int64]) { return })
var _ = Over[int, int64](func(...any) (res Result[int64]) { return })
var _ = Over[int, int64](func(int, ...int) (res Result[int64]) { return })
var _ = Over[int, int64](func(int, ...any) (res Result[int64]) { return })
var _ = Over[int, int64](func(any, ...any) (res Result[int64]) { return })

// Filter maps native
var _ = Over[int, int64](func(int) (res int64, v bool) { return })
var _ = Over[int, int64](func(...int) (res int64, v bool) { return })
var _ = Over[int, int64](func(any) (res int64, v bool) { return })
var _ = Over[int, int64](func(...any) (res int64, v bool) { return })
var _ = Over[int, int64](func(int, ...int) (res int64, v bool) { return })
var _ = Over[int, int64](func(int, ...any) (res int64, v bool) { return })
var _ = Over[int, int64](func(any, ...any) (res int64, v bool) { return })
var _ = Over[int, int64](func(int) (res int64, v error) { return })
var _ = Over[int, int64](func(...int) (res int64, v error) { return })
var _ = Over[int, int64](func(any) (res int64, v error) { return })
var _ = Over[int, int64](func(...any) (res int64, v error) { return })
var _ = Over[int, int64](func(int, ...int) (res int64, v error) { return })
var _ = Over[int, int64](func(int, ...any) (res int64, v error) { return })
var _ = Over[int, int64](func(any, ...any) (res int64, v error) { return })

// Any other function signature will not compile.
// Over and Through will only ever panic if called with reflection.

func Fold[T any, B any](fn func(B, T) B, init ...B) func([]T) B {
	var in B
	if len(init) > 0 { in = init[0] }
	return func(args []T) B {
		res := in
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

// Not negates a [Predicate].
func Not[T any](fn Pred[T]) Pred[T] { return func(t T) bool { return !fn(t) } }

func IsZero[K comparable](a K) bool {
	var def K
	return a == def
}

// Is returns a [Predicate] that checks if the argument is in the list.
func Is[C comparable](T ...C) func(C) bool {
	in := make(map[C]bool, len(T))
	for _, a := range T {
		in[a] = true
	}
	return func(c C) bool {
		_, ok := in[c]
		return ok
	}
}

// FindFirst returns the first value which passes the [Predicate].
func FindFirst[T any](fns ...Pred[T]) func([]T) Option[T] {
	fn := PredOr(fns...)
	return func(args []T) Option[T] {
		for _, arg := range args {
			if fn(arg) {
				return Some(arg)
			}
		}
		return None[T]()
	}
}

// Pipe runs a value through a pipeline or composes functions.
//
// If the first argument is a value, it executes a pipeline:
// T1, (T1) -> T2, (T2) -> T3, ..., (Tn-1) -> Tn
// and returns the final value: Tn
//
// If the first argument is a function, it composes a pipeline:
// (T1) -> T2, (T2) -> T3, ..., (Tn-1) -> Tn
// into a final function: (T1) -> Tn
func Pipe[Output any](values ...any) Output {
	// If no arguments are provided, return the zero value of the output type.
	var zero Output
	if len(values) == 0 {
		return zero
	}

	first := reflect.ValueOf(values[0])

	if first.Kind() != reflect.Func {
		if len(values) == 1 {
			// Only a single value was passed, return it.
			return first.Interface().(Output)
		}

		result := first
		for i := 1; i < len(values); i++ {
			fn := reflect.ValueOf(values[i])

			if fn.Kind() != reflect.Func {
				panic("Pipe Error: For value processing, all subsequent arguments must be functions.")
			}
			outputs := fn.Call([]reflect.Value{result})

			if len(outputs) == 0 {
				if i < len(values)-1 {
					panic("Pipe Error: A function in the middle of the pipeline returned no value, breaking the chain.")
				}
				return zero
			}
			result = outputs[0]
		}
		return result.Interface().(Output)
	}

	funcs := make([]reflect.Value, len(values))
	for i, v := range values {
		fn := reflect.ValueOf(v)
		if fn.Kind() != reflect.Func {
			panic("Pipe Error: For function composition, all arguments must be functions.")
		}
		funcs[i] = fn
	}

	for i := range len(funcs)-1 {
		if funcs[i].Type().NumOut() != funcs[i+1].Type().NumIn() {
			panic(fmt.Sprintf(
				"Pipe Error: Arity mismatch between function %d (returns %d values) and function %d (expects %d values).",
				i, funcs[i].Type().NumOut(), i+1, funcs[i+1].Type().NumIn(),
			))
		}

		for j := range funcs[i].Type().NumOut() {
			if funcs[i].Type().Out(j) != funcs[i+1].Type().In(j) {
				panic(fmt.Sprintf(
					"Pipe Error: Type mismatch between output %d of function %d (%s) and input %d of function %d (%s).",
					j, i, funcs[i].Type().Out(j), j, i+1, funcs[i+1].Type().In(j),
				))
			}
		}
	}

	firstFuncType := funcs[0].Type()
	lastFuncType := funcs[len(funcs)-1].Type()

	inTypes := make([]reflect.Type, firstFuncType.NumIn())
	for i := range firstFuncType.NumIn() {
		inTypes[i] = firstFuncType.In(i)
	}

	outTypes := make([]reflect.Type, lastFuncType.NumOut())
	for i := range lastFuncType.NumOut() {
		outTypes[i] = lastFuncType.Out(i)
	}

	composedFuncType := reflect.FuncOf(inTypes, outTypes, firstFuncType.IsVariadic())

	composedFuncImpl := func(args []reflect.Value) []reflect.Value {
		var currentResult = args
		for _, fn := range funcs {
			currentResult = fn.Call(currentResult)
		}
		return currentResult
	}

	composedFunc := reflect.MakeFunc(composedFuncType, composedFuncImpl)

	return composedFunc.Interface().(Output)
}

func String[T byte | []byte | rune | []rune | string](arg T) string { return string(arg) }

func FromTo[I, O []rune | []byte | string](input I) O { return O(string(input)) }

func Cat[I, T, O any](a func(I) T, b func(T) O) func(I) O { return func(c I) O { return b(a(c)) } }


func IfCat[I, T, O any](actor func(I) (T, bool), transformer func(T) O) func(I) (O, bool) {
	return func(i I) (res O, ok bool) {
		var b T
		b, ok = actor(i)
		if ok { res = transformer(b) }
		return
	}
}

func IfCat2[I1, I2, T, O any](actor func(I1, I2) (T, bool), transformer func(T) O) func(I1, I2) (O, bool) {
	return func(i1 I1, i2 I2) (res O, ok bool) {
		var b T
		b, ok = actor(i1, i2)
		if ok { res = transformer(b) }
		return
	}
}

func PredAnd[T any](preds ...Pred[T]) Pred[T] {
	return func(a T) bool {
		for _, pred := range preds {
			if !pred(a) {
				return false
			}
		}
		return true
	}
}

func PredOr[T any](preds ...Pred[T]) Pred[T] {
	return func(a T) bool {
		for _, pred := range preds {
			if pred(a) {
				return true
			}
		}
		return false
	}
}

func limit[T ~string | ~[]any](Max int) func([]T) []T {
	return func(args []T) (res []T) {
		for _, a := range args {
			if len(a) <= Max {
				res = append(res, a)
			}
		}
		return
	}
}

func Includes[K comparable](inclusive bool) func(args ...K) func([]K) bool {
	return func(args ...K) func([]K) bool {
		var pred Pred[K]
		// if inclusive { pred = Is(args...) } else { pred = IsEvery(args...) }
		pred = Is(args...)
		return func(arr []K) bool { return slices.ContainsFunc(arr, pred) }
	}
}

func Vals[K comparable, V any](m map[K]V) (res []V) {
	if m == nil { return }
	for _, v := range m {res = append(res, v)}
	return
}

func Keys[K comparable, V any](m map[K]V) (res []K) {
	if m == nil { return }
	for k := range m {res = append(res, k)}
	return
}
