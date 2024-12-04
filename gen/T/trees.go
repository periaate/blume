package T

type Err[A any] struct {
	err    error
	reason string
	data   A
}

func (e Err[A]) Error() error {
	return e.err
}

func (e Err[A]) Reason() string {
	return e.reason
}

func (e Err[A]) Data() A {
	return e.data
}

type Error[A any] interface {
	Error() error
	Reason() string
	Data() A
}

type TreeLike[A any] interface {
	Collect[A]
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
