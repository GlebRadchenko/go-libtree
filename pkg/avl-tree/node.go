package avl_tree

type Node[Key comparable, Value any] struct {
	Key    Key
	Value  Value
	Height int

	left  *Node[Key, Value]
	right *Node[Key, Value]
}
