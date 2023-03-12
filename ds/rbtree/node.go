package rbtree

// Color defines node color type
type Color bool

// Define node 's colors
const (
	RED   = false
	BLACK = true
)

// Node is a tree node
type Node[K, V any] struct {
	 //TODO: Complete me!
}

// Key returns node's key
func (n *Node[K, V]) Key() K {
	 //TODO: Complete me!
}

// Value returns node's value
func (n *Node[K, V]) Value() V {
	 //TODO: Complete me!
}

// SetValue sets node's value
func (n *Node[K, V]) SetValue(val V) {
	 //TODO: Complete me!
}

// Next returns the Node's successor as an iterator.
func (n *Node[K, V]) Next() *Node[K, V] {
	 //TODO: Complete me!
}

// Prev returns the Node's predecessor as an iterator.
func (n *Node[K, V]) Prev() *Node[K, V] {
	 //TODO: Complete me!
}

// successor returns the successor of the Node
func successor[K, V any](x *Node[K, V]) *Node[K, V] {
	 //TODO: Complete me!
}

// presuccessor returns the presuccessor of the Node
func presuccessor[K, V any](x *Node[K, V]) *Node[K, V] {
	 //TODO: Complete me!
}

// minimum finds the minimum Node of subtree n.
func minimum[K any, V any](n *Node[K, V]) *Node[K, V] {
	 //TODO: Complete me!
}

// maximum finds the maximum Node of subtree n.
func maximum[K any, V any](n *Node[K, V]) *Node[K, V] {
	 //TODO: Complete me!
}
