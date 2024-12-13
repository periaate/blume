package maps

import . "github.com/periaate/blume"

type head[V any] struct {
	root_Top *node[V]
	root_Bot *node[V]
	Len    int
}

type node[V any] struct {
	Top *node[V]
	Bot *node[V]
	V V
}

type End bool

const (
	Top End = false
	Bot End = true
)

func (h *head[V]) Pop(end End) (opt Option[V]) {
	if h.Len == 0 { return None[V]() }
	switch end {
	case Top:
		if h.root_Top == nil { return None[V]() }
		opt = Some(h.root_Top.V)
		h.root_Top = h.root_Top.Bot
	case Bot:
		if h.root_Bot == nil { return None[V]() }
		opt = Some(h.root_Bot.V)
		h.root_Bot = h.root_Bot.Top
	}
	h.Len--
	return
}

func (h *head[V]) Push(val V, end End) V {
	switch end {
	case Top:
		n := &node[V]{V: val, Bot: h.root_Top}
		if h.root_Top != nil { h.root_Top.Top = n }
		h.root_Top = n
		if h.root_Bot == nil { h.root_Bot = n }
	case Bot:
		n := &node[V]{V: val, Top: h.root_Bot}
		if h.root_Bot != nil { h.root_Bot.Bot = n }
		h.root_Bot = n
		if h.root_Top == nil { h.root_Top = n }
	}
	h.Len++
	return val
}
