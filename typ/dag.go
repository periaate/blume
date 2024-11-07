package typ

/*
HDAG is a nested hash-map, creating a DAG.
It allows for jumping to any unique child
from any point in the tree from which it is accessible.

A
	One
		Bar
		Baz
	Two
B
	Three
C
	Four
		Foo
		Baz
	Five

"A" can be accessed from the root.
"Bar" can be accessed from the root.
"Baz" has two matches, and can not be accessed for the root.

All of the following are valid searches
	For A/One/Baz
	"A Baz"
	"One Baz"

	For C/Four/Baz
	"C Baz"
	"Four Baz"

*/
// type HDAG[K comparable, V any] struct {
// 	map[K]Node
// }
//
// type Node[K comparable, V any] struct {
// 	Value V
// }
//
