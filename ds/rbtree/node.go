package rbtree

// Color defines node color type
type Color bool

// Define node 's colors
const (
	RED   = false
	BLACK = true
)

// Node is a tree node
type Node struct {
	 //TODO: Complete me!
}

// Key returns node's key
func (n *Node) Key() interface{} {
	 //TODO: Complete me!
}

// Value returns node's value
func (n *Node) Value() interface{} {
	 //TODO: Complete me!
}

// SetValue sets node's value
func (n *Node) SetValue(val interface{}) {
	 //TODO: Complete me!
}

// Next returns the Node's successor as an iterator.
func (n *Node) Next() *Node {
	 //TODO: Complete me!
}

// Prev returns the Node's predecessor as an iterator.
func (n *Node) Prev() *Node {
	 //TODO: Complete me!
}

// successor returns the successor of the Node
func successor(x *Node) *Node {
	 //TODO: Complete me!
}

// presuccessor returns the presuccessor of the Node
func presuccessor(x *Node) *Node {
	 //TODO: Complete me!
}

// minimum finds the minimum Node of subtree n.
func minimum(n *Node) *Node {
	 //TODO: Complete me!
}

// maximum finds the maximum Node of subtree n.
func maximum(n *Node) *Node {
	 //TODO: Complete me!
}
