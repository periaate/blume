package types

import (
	. "github.com/periaate/blume/core"
)

var _ = Zero[any]
var _ TreeLike[string] = Tree[string]{}

type Tree[A any] struct {
	Nodes []Tree[A]
	Value A
}

func (t *Tree[A]) Append(arr ...A) {
	nodes := make([]Tree[A], len(arr))
	for i, v := range arr {
		nodes[i] = Tree[A]{Value: v}
	}
	t.Nodes = append(t.Nodes, nodes...)
}

func (t Tree[A]) Leaves() Array[A] {
	if len(t.Nodes) == 0 { return ToArray([]A{t.Value}) }
	var res []A
	for _, node := range t.Nodes {
		res = append(res, node.Leaves().Values()...)
	}
	return ToArray(res)
}

func (t Tree[A]) Values() Array[A] {
	var res []A
	for _, node := range t.Nodes {
		res = append(res, node.Value)
		res = append(res, node.Values().Values()...)
	}
	return ToArray(res)
}

func (t Tree[A]) Filter(preds ...Predicate[A]) Array[A] { return t.Values().Filter(preds...) }

func (t Tree[A]) TraverseDepth(depth int, fn func(A) Error[any]) Error[any] {
	return t.traverseDepthHelper(0, depth, fn)
}

func (t Tree[A]) traverseDepthHelper(currentDepth, maxDepth int, fn func(A) Error[any]) Error[any] {
	if currentDepth > maxDepth { return nil }
	err := fn(t.Value)
	if err != nil { return err }
	for _, node := range t.Nodes {
		childErr := node.traverseDepthHelper(currentDepth+1, maxDepth, fn)
		if childErr != nil { return childErr }
	}
	return nil
}

func (t Tree[A]) TraverseBreadth(depth int, fn func(A) Error[any]) Error[any] {
	type nodeDepth struct {
		node  Tree[A]
		depth int
	}
	queue := []nodeDepth{{t, 0}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.depth > depth { continue }
		err := fn(current.node.Value)
		if err != nil { return err }
		for _, child := range current.node.Nodes {
			queue = append(queue, nodeDepth{child, current.depth + 1})
		}
	}
	return nil
}
