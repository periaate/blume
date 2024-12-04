package gen

import (
	"strings"

	"github.com/periaate/blume/gen/T"
)

var _ T.TreeLike[string] = Tree[string]{}

type Tree[A any] struct {
	Nodes []Tree[A]
	Value A
}

func (t Tree[A]) Collect() []A {
	var res []A
	for _, node := range t.Nodes {
		res = append(res, node.Collect()...)
	}
	return append(res, t.Value)
}

func (t Tree[A]) Filter(preds ...T.Predicate[A]) []A {
	return Filter(preds...)(t.Collect())
}

func (t Tree[A]) TraverseDepth(depth int, fn func(A) T.Error[any]) T.Error[any] {
	return t.traverseDepthHelper(0, depth, fn)
}

func (t Tree[A]) traverseDepthHelper(currentDepth, maxDepth int, fn func(A) T.Error[any]) T.Error[any] {
	if currentDepth > maxDepth {
		return nil
	}
	err := fn(t.Value)
	if err != nil {
		return err
	}
	for _, node := range t.Nodes {
		childErr := node.traverseDepthHelper(currentDepth+1, maxDepth, fn)
		if childErr != nil {
			return childErr
		}
	}
	return nil
}

func (t Tree[A]) TraverseBreadth(depth int, fn func(A) T.Error[any]) T.Error[any] {
	type nodeDepth struct {
		node  Tree[A]
		depth int
	}
	queue := []nodeDepth{{t, 0}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.depth > depth {
			continue
		}
		err := fn(current.node.Value)
		if err != nil {
			return err
		}
		for _, child := range current.node.Nodes {
			queue = append(queue, nodeDepth{child, current.depth + 1})
		}
	}
	return nil
}

func (t Tree[A]) Format(f func(A) string) string {
	var sb strings.Builder
	t.formatHelper(&sb, f, 0)
	return sb.String()
}

func (t Tree[A]) formatHelper(sb *strings.Builder, f func(A) string, indent int) {
	prefix := strings.Repeat("  ", indent)
	sb.WriteString(prefix)
	sb.WriteString(f(t.Value))
	sb.WriteString("\n")
	for _, child := range t.Nodes {
		child.formatHelper(sb, f, indent+1)
	}
}
