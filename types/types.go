package types

import (
	. "github.com/periaate/blume/core"
)

var _ = Zero[any]


type TreeLike[A any] interface {
	Leaves() Array[A]
	Values() Array[A]
	// Recursively filter the tree with the given predicates.
	Filter(preds ...Monadic[A, bool]) Array[A]
	// TraverseDepth traverses the tree in depth-first order, calling the function on each node.
	// If the function returns an error, the traversal stops and the error is returned.
	TraverseDepth(depth int, fn func(A) Error[any]) Error[any]
	// TraverseBreadth traverses the tree in breadth-first order, calling the function on each node.
	// If the function returns an error, the traversal stops and the error is returned.
	TraverseBreadth(depth int, fn func(A) Error[any]) Error[any]
}
