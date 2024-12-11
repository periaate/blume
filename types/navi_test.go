package types

import (
	"testing"

	. "github.com/periaate/blume/core"
	"github.com/stretchr/testify/assert"
)

var _ = Zero[any]

func TestNavi(t *testing.T) {
	tree := Tree[int]{}
	tree.Append(1, 5, 8, 9)
	tree.Val[0].Append(2, 3, 4)
	tree.Val[1].Append(6, 7)

	tn := NaviFromTree(tree)
	assert.True(t, tn.IsOk())
	navi := tn.Unwrap()
	assert.Equal(t, 4, navi.Root.Len())

	assert.Equal(t, 1, navi.Current().Unwrap().Value)
	assert.Equal(t, 0, navi.Indx)

	tn = navi.Next(1)
	assert.True(t, tn.IsOk())
	navi = tn.Unwrap()
	assert.Equal(t, 5, navi.Current().Unwrap().Value)
	assert.Equal(t, 1, navi.Indx)

	tn = navi.Prev(1)
	assert.True(t, tn.IsOk())
	navi = tn.Unwrap()
	assert.Equal(t, 1, navi.Current().Unwrap().Value)
	assert.Equal(t, 0, navi.Indx)

	tn = navi.Descend()
	assert.True(t, tn.IsOk())
	navi = tn.Unwrap()
	assert.Equal(t, 1, navi.Path.Len())
	assert.Equal(t, 2, navi.Current().Unwrap().Value)
	assert.Equal(t, 0, navi.Indx)

	tn = navi.Next(1)
	assert.True(t, tn.IsOk())
	navi = tn.Unwrap()
	assert.Equal(t, 3, navi.Current().Unwrap().Value)
	assert.Equal(t, 1, navi.Indx)

	tn = navi.Ascend()
	assert.True(t, tn.IsOk())
	navi = tn.Unwrap()
	assert.Equal(t, 1, navi.Current().Unwrap().Value)
	assert.Equal(t, 0, navi.Indx)

	tn = navi.Next(2)
	assert.True(t, tn.IsOk())
	navi = tn.Unwrap()
	assert.Equal(t, 8, navi.Current().Unwrap().Value)
	assert.Equal(t, 2, navi.Indx)

	tn = navi.Descend()
	assert.False(t, tn.IsOk())

	tn = navi.Prev(1)
	assert.True(t, tn.IsOk())
	navi = tn.Unwrap()
}
