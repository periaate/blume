package gen_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
	"github.com/stretchr/testify/assert"
)

// MyError is a custom error type implementing T.Error[any]
type MyError[A any] struct {
	err    error
	reason string
	data   A
}

func (e MyError[A]) Error() error {
	return e.err
}

func (e MyError[A]) Reason() string {
	return e.reason
}

func (e MyError[A]) Data() A {
	return e.data
}

func TestTreeTraverseDepth(t *testing.T) {
	tree := gen.Tree[string]{
		Value: "A",
		Nodes: []gen.Tree[string]{
			{
				Value: "B",
				Nodes: []gen.Tree[string]{
					{Value: "D"},
					{Value: "E"},
				},
			},
			{
				Value: "C",
				Nodes: []gen.Tree[string]{
					{Value: "F"},
				},
			},
		},
	}

	var visited []string

	err := tree.TraverseDepth(10, func(value string) T.Error[any] {
		visited = append(visited, value)
		return nil
	})

	assert.Nil(t, err)
	expected := []string{"A", "B", "D", "E", "C", "F"}
	assert.Equal(t, expected, visited)
}

func TestTreeTraverseBreadth(t *testing.T) {
	tree := gen.Tree[string]{
		Value: "A",
		Nodes: []gen.Tree[string]{
			{
				Value: "B",
				Nodes: []gen.Tree[string]{
					{Value: "D"},
					{Value: "E"},
				},
			},
			{
				Value: "C",
				Nodes: []gen.Tree[string]{
					{Value: "F"},
				},
			},
		},
	}

	var visited []string

	err := tree.TraverseBreadth(10, func(value string) T.Error[any] {
		visited = append(visited, value)
		return nil
	})

	assert.Nil(t, err)
	expected := []string{"A", "B", "C", "D", "E", "F"}
	assert.Equal(t, expected, visited)
}

func TestTreeFormat(t *testing.T) {
	tree := gen.Tree[string]{
		Value: "A",
		Nodes: []gen.Tree[string]{
			{
				Value: "B",
				Nodes: []gen.Tree[string]{
					{Value: "D"},
					{Value: "E"},
				},
			},
			{
				Value: "C",
				Nodes: []gen.Tree[string]{
					{Value: "F"},
				},
			},
		},
	}

	formatted := tree.Format(func(s string) string { return s })

	expected := `A
  B
    D
    E
  C
    F
`

	assert.Equal(t, expected, formatted)
}

func TestTreeTraverseDepthWithError(t *testing.T) {
	tree := gen.Tree[string]{
		Value: "A",
		Nodes: []gen.Tree[string]{
			{
				Value: "B",
				Nodes: []gen.Tree[string]{
					{Value: "D"},
					{Value: "E"},
				},
			},
			{
				Value: "C",
				Nodes: []gen.Tree[string]{
					{Value: "F"},
				},
			},
		},
	}

	var visited []string

	err := tree.TraverseDepth(10, func(value string) T.Error[any] {
		visited = append(visited, value)
		if value == "D" {
			return MyError[any]{err: errors.New("found D"), reason: "found D", data: nil}
		}
		return nil
	})

	assert.NotNil(t, err)
	assert.Equal(t, []string{"A", "B", "D"}, visited)
}

func TestTreeTraverseBreadthWithError(t *testing.T) {
	tree := gen.Tree[string]{
		Value: "A",
		Nodes: []gen.Tree[string]{
			{
				Value: "B",
				Nodes: []gen.Tree[string]{
					{Value: "D"},
					{Value: "E"},
				},
			},
			{
				Value: "C",
				Nodes: []gen.Tree[string]{
					{Value: "F"},
				},
			},
		},
	}

	var visited []string

	err := tree.TraverseBreadth(10, func(value string) T.Error[any] {
		visited = append(visited, value)
		if value == "C" {
			return MyError[any]{err: errors.New("found C"), reason: "found C", data: nil}
		}
		return nil
	})

	assert.NotNil(t, err)
	assert.Equal(t, []string{"A", "B", "C"}, visited)
}

func TestTreeTraverseDepthMaxDepth(t *testing.T) {
	tree := gen.Tree[string]{
		Value: "A",
		Nodes: []gen.Tree[string]{
			{
				Value: "B",
				Nodes: []gen.Tree[string]{
					{Value: "D"},
					{Value: "E"},
				},
			},
			{
				Value: "C",
				Nodes: []gen.Tree[string]{
					{Value: "F"},
				},
			},
		},
	}

	var visited []string

	err := tree.TraverseDepth(1, func(value string) T.Error[any] {
		visited = append(visited, value)
		return nil
	})

	assert.Nil(t, err)
	expected := []string{"A", "B", "C"}
	assert.Equal(t, expected, visited)
}

func TestTreeFormatWithFormatter(t *testing.T) {
	tree := gen.Tree[int]{
		Value: 1,
		Nodes: []gen.Tree[int]{
			{
				Value: 2,
				Nodes: []gen.Tree[int]{
					{Value: 4},
					{Value: 5},
				},
			},
			{
				Value: 3,
				Nodes: []gen.Tree[int]{
					{Value: 6},
				},
			},
		},
	}

	formatted := tree.Format(func(i int) string { return fmt.Sprintf("Node-%d", i) })

	expected := `Node-1
  Node-2
    Node-4
    Node-5
  Node-3
    Node-6
`

	assert.Equal(t, expected, formatted)
}