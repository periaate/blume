package types_test

import (
	. "github.com/periaate/blume/core"
)

var _ = Zero[any]

// func TestTreeTraverseDepth(t *testing.T) {
// 	tree := types.Tree[string]{
// 		Value: "A",
// 		Nodes: []types.Tree[string]{
// 			{
// 				Value: "B",
// 				Nodes: []types.Tree[string]{
// 					{Value: "D"},
// 					{Value: "E"},
// 				},
// 			},
// 			{
// 				Value: "C",
// 				Nodes: []types.Tree[string]{
// 					{Value: "F"},
// 				},
// 			},
// 		},
// 	}
//
// 	var visited []string
//
// 	err := tree.TraverseDepth(10, func(value string) Error[any] {
// 		visited = append(visited, value)
// 		return nil
// 	})
//
// 	assert.Nil(t, err)
// 	expected := []string{"A", "B", "D", "E", "C", "F"}
// 	assert.Equal(t, expected, visited)
// }
//
// func TestTreeTraverseBreadth(t *testing.T) {
// 	tree := types.Tree[string]{
// 		Value: "A",
// 		Nodes: []types.Tree[string]{
// 			{
// 				Value: "B",
// 				Nodes: []types.Tree[string]{
// 					{Value: "D"},
// 					{Value: "E"},
// 				},
// 			},
// 			{
// 				Value: "C",
// 				Nodes: []types.Tree[string]{
// 					{Value: "F"},
// 				},
// 			},
// 		},
// 	}
//
// 	var visited []string
//
// 	err := tree.TraverseBreadth(10, func(value string) Error[any] {
// 		visited = append(visited, value)
// 		return nil
// 	})
//
// 	assert.Nil(t, err)
// 	expected := []string{"A", "B", "C", "D", "E", "F"}
// 	assert.Equal(t, expected, visited)
// }
//
// func TestTreeTraverseDepthWithError(t *testing.T) {
// 	tree := types.Tree[string]{
// 		Value: "A",
// 		Nodes: []types.Tree[string]{
// 			{
// 				Value: "B",
// 				Nodes: []types.Tree[string]{
// 					{Value: "D"},
// 					{Value: "E"},
// 				},
// 			},
// 			{
// 				Value: "C",
// 				Nodes: []types.Tree[string]{
// 					{Value: "F"},
// 				},
// 			},
// 		},
// 	}
//
// 	var visited []string
//
// 	err := tree.TraverseDepth(10, func(value string) Error[any] {
// 		visited = append(visited, value)
// 		if value == "D" { return StrError("found D") }
// 		return nil
// 	})
//
// 	assert.NotNil(t, err)
// 	assert.Equal(t, []string{"A", "B", "D"}, visited)
// }
//
// func TestTreeTraverseBreadthWithError(t *testing.T) {
// 	tree := types.Tree[string]{
// 		Value: "A",
// 		Nodes: []types.Tree[string]{
// 			{
// 				Value: "B",
// 				Nodes: []types.Tree[string]{
// 					{Value: "D"},
// 					{Value: "E"},
// 				},
// 			},
// 			{
// 				Value: "C",
// 				Nodes: []types.Tree[string]{
// 					{Value: "F"},
// 				},
// 			},
// 		},
// 	}
//
// 	var visited []string
//
// 	err := tree.TraverseBreadth(10, func(value string) Error[any] {
// 		visited = append(visited, value)
// 		if value == "C" { return StrError("found C") }
// 		return nil
// 	})
//
// 	assert.NotNil(t, err)
// 	assert.Equal(t, []string{"A", "B", "C"}, visited)
// }
//
// func TestTreeTraverseDepthMaxDepth(t *testing.T) {
// 	tree := types.Tree[string]{
// 		Value: "A",
// 		Nodes: []types.Tree[string]{
// 			{
// 				Value: "B",
// 				Nodes: []types.Tree[string]{
// 					{Value: "D"},
// 					{Value: "E"},
// 				},
// 			},
// 			{
// 				Value: "C",
// 				Nodes: []types.Tree[string]{
// 					{Value: "F"},
// 				},
// 			},
// 		},
// 	}
//
// 	var visited []string
//
// 	err := tree.TraverseDepth(1, func(value string) Error[any] {
// 		visited = append(visited, value)
// 		return nil
// 	})
//
// 	assert.Nil(t, err)
// 	expected := []string{"A", "B", "C"}
// 	assert.Equal(t, expected, visited)
// }
