package types

import (
	. "github.com/periaate/blume/core"
)

var _ = Zero[any]

func NaviFromTree[A any](t Tree[A]) Option[*Navi[A]] {
	if t.Len() == 0 { return None[*Navi[A]](nil) }
	navi := new(Navi[A])
	navi.Root = t
	return Some(navi)
}


type Navi[A any] struct {
	Root Tree[A]
	Path Array[*Navi[A]]
	Indx int
}

// Look returns a range from current 
func (navi *Navi[A]) Look(offset, count int) Option[Array[Tree[A]]] {
	return Some(navi.Root.Array.Slice(offset, count))
}

func (navi *Navi[A]) Current() Option[Tree[A]] {
	if navi.Root.Len() == 0 { return None[Tree[A]](nil) }
	return Some(navi.Root.Array.Val[navi.Indx])
}

func (navi *Navi[A]) Prev(n int) Option[*Navi[A]] {
	if navi.Root.Len() == 0 { return None[*Navi[A]](nil) }
	if navi.Root.Len() < navi.Indx+n || navi.Indx-n < 0 { return None[*Navi[A]](nil) }
	navi.Indx -= n
	return Some(navi)
}

func (navi *Navi[A]) Next(n int) Option[*Navi[A]] {
	if navi.Root.Len() == 0 { return None[*Navi[A]](nil) }
	if navi.Root.Len() < navi.Indx+n || navi.Indx+n < 0 { return None[*Navi[A]](nil) }
	navi.Indx += n
	return Some(navi)
}

func (navi *Navi[A]) Ascend () Option[*Navi[A]] {
	if navi.Path.Len() == 0 { return None[*Navi[A]](nil) }
	opt := navi.Path.Pop()
	if !opt.IsOk() { return None[*Navi[A]](nil) }
	return Some(opt.Unwrap())
}

func (navi *Navi[A]) Descend() Option[*Navi[A]] {
	if navi.Root.Len() == 0 { return None[*Navi[A]](nil) }
	tc := navi.Current()
	if !tc.IsOk() { return None[*Navi[A]](nil) }
	curr := tc.Unwrap()

	tn := NaviFromTree(curr)
	if !tn.IsOk() { return None[*Navi[A]](nil) }
	newNavi := tn.Unwrap()
	newNavi.Path = navi.Path
	newNavi.Path = newNavi.Path.Append(navi)
	return Some(newNavi)
}
