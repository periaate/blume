package types

import (
	. "github.com/periaate/blume/core"
)

var _ = Zero[any]
// var _ TreeLike[string] = Tree[string]{}

type Tree[A any] struct {
	Array[Tree[A]]
	Value A
}

func (t *Tree[A]) Append(arr ...A) {
	nodes := make([]Tree[A], len(arr))
	for i, v := range arr {
		nodes[i] = Tree[A]{Value: v}
	}
	nodes = append(t.Array.Values(), nodes...)
	t.Array = ToArray(nodes)
}

func (t *Tree[A]) Prepend(arr ...A) {
	nodes := make([]Tree[A], len(arr))
	for i, v := range arr {
		nodes[i] = Tree[A]{Value: v}
	}
	nodes = append(nodes, t.Array.Values()...)
	t.Array = ToArray(nodes)
}

func (t Tree[A]) Leaves() Array[A] {
	if t.Array.Len() == 0 { return ToArray([]A{t.Value}) }
	var res []A
	for _, node := range t.Array.Values() {
		res = append(res, node.Leaves().Values()...)
	}
	return ToArray(res)
}

func (t Tree[A]) Values() Array[A] {
	var res []A
	for _, node := range t.Array.Values() {
		res = append(res, node.Value)
		res = append(res, node.Values().Values()...)
	}
	return ToArray(res)
}

func (t Tree[A]) Filter(preds ...Monadic[A, bool]) Array[A] { return t.Values().Filter(preds...) }

func (t Tree[A]) TraverseDepth(depth int, fn func(int, A) Error[any]) Error[any] {
	return t.traverseDepthHelper(0, depth, fn)
}

func (t Tree[A]) traverseDepthHelper(currentDepth, maxDepth int, fn func(int, A) Error[any]) Error[any] {
	if currentDepth > maxDepth { return nil }
	err := fn(currentDepth, t.Value)
	if err != nil { return err }
	for _, node := range t.Array.Values() {
		childErr := node.traverseDepthHelper(currentDepth+1, maxDepth, fn)
		if childErr != nil { return childErr }
	}
	return nil
}

func (t Tree[A]) TraverseBreadth(depth int, fn func(int, A) Error[any]) Error[any] {
	type nodeDepth struct {
		node  Tree[A]
		depth int
	}
	queue := []nodeDepth{{t, 0}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.depth > depth { continue }
		err := fn(depth, current.node.Value)
		if err != nil { return err }
		for _, child := range current.node.Array.Values() {
			queue = append(queue, nodeDepth{child, current.depth + 1})
		}
	}
	return nil
}
