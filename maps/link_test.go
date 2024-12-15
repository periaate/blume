package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLink(t *testing.T) {
	h := head[int]{}

	assert.Equal(t, 2, h.Push(2, Top))
	assert.Equal(t, 1, h.Push(1, Top))
	assert.Equal(t, 2, h.Pop(Bot).Value)
	assert.Equal(t, 1, h.Pop(Bot).Value)
	assert.False(t, h.Pop(Bot).Ok)
	assert.Equal(t, 0, h.Len)

	assert.Equal(t, 0, h.Push(0, Top))
	assert.Equal(t, 1, h.Push(1, Top))
	assert.Equal(t, 2, h.Push(2, Top))

	assert.Equal(t, 0, h.Push(0, Bot))
	assert.Equal(t, 1, h.Push(1, Bot))
	assert.Equal(t, 2, h.Push(2, Bot))

	assert.Equal(t, 6, h.Len)
	assert.Equal(t, 2, h.Pop(Bot).Value)
	assert.Equal(t, 2, h.Pop(Top).Value)
	assert.Equal(t, 4, h.Len)
	assert.Equal(t, 1, h.Pop(Bot).Value)
	assert.Equal(t, 1, h.Pop(Top).Value)
	assert.Equal(t, 2, h.Len)
	assert.Equal(t, 0, h.Pop(Bot).Value)
	assert.Equal(t, 0, h.Pop(Top).Value)
	assert.Equal(t, 0, h.Len)

	assert.False(t, h.Pop(Bot).Ok)
	assert.False(t, h.Pop(Top).Ok)
	assert.Equal(t, 0, h.Len)
}
