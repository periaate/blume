package maps

type Operation uint8

const (
	OP_NIL Operation = iota
	OP_GET
	OP_SET
	OP_DEL
	OP_RET
)

type Hooks[K comparable, V any] struct {
	Del func(K) (K, Operation)
	Get func(K, V) (K, V, Operation)
	Set func(K, V) (K, V, Operation)
}
