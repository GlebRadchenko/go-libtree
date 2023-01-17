package avl_tree

type Tree[Key comparable, Value any] struct {
	root *Node[Key, Value]
}
