package gen

type Tree[A any] struct {
	Nodes []Tree[A]
	Value A
}
