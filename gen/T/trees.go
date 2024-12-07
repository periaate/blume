package T

type Err[A any] struct {
	err    error
	reason string
	data   A
}

type errt string

func (e errt) Error() string { return string(e) }

func Errs[A any](er any, reason string, data A) Err[A] {
	switch v := er.(type) {
	case error:
		return Err[A]{err: v, reason: reason, data: data}
	case string:
		return Err[A]{err: errt(v), reason: reason, data: data}
	default:
		panic("invalid error type")
	}
}

func (e Err[A]) Err() error     { return e.err }
func (e Err[A]) Reason() string { return e.reason }
func (e Err[A]) Data() A        { return e.data }
func (e Err[A]) Error() string  { return "error: " + e.err.Error() + "; beacuse: " + e.reason }

type Error[A any] interface {
	Error() string
	Err() error
	Reason() string
	Data() A
}

type TreeLike[A any] interface {
	// Recursive
	Collect[A]
	// Recursive
	Filters[A]

	// TraverseDepth traverses the tree in depth-first order, calling the function on each node.
	// If the function returns an error, the traversal stops and the error is returned.
	TraverseDepth(depth int, fn func(A) Error[any]) Error[any]
	// TraverseBreadth traverses the tree in breadth-first order, calling the function on each node.
	// If the function returns an error, the traversal stops and the error is returned.
	TraverseBreadth(depth int, fn func(A) Error[any]) Error[any]

	// Format returns a human readable string representation of the tree.
	// Needs a function that converts the value to a string.
	Format(func(A) string) string
}
