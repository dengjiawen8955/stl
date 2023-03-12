package simplelist

import (
	"fmt"
	"stl/utils/visitor"
)

// Node is a list node
type Node struct {
	 //TODO: Complete me!
}

// Next returns the next list node or nil.
func (n *Node) Next() *Node {
	 //TODO: Complete me!
}

// List represents a single direction list:
//
//   head -> node1 --> node2 --> node3 <- tail
//
type List struct {
	 //TODO: Complete me!
}

// New creates a list
func New() *List {
	 //TODO: Complete me!
}

// Len returns the amount of list nodes.
func (l *List) Len() int {
	 //TODO: Complete me!
}

// FrontNode returns the front node of the list or nil if the list is empty
func (l *List) FrontNode() *Node {
	 //TODO: Complete me!
}

// BackNode returns the last node of the list or nil if the list is empty
func (l *List) BackNode() *Node {
	 //TODO: Complete me!
}

// PushFront inserts a new node n with value v at the front of the list.
func (l *List) PushFront(v interface{}) {
	 //TODO: Complete me!
}

// PushBack inserts a new node n with value v at the back of the list.
func (l *List) PushBack(v interface{}) {
	 //TODO: Complete me!
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of the list, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	 //TODO: Complete me!
}

func (l *List) insertAfter(n, at *Node) *Node {
	 //TODO: Complete me!
}

// Remove removes node n from the list.
// The node must not be nil.
func (l *List) Remove(pre, n *Node) interface{} {
	 //TODO: Complete me!
}

// MoveToFront moves node n to the front of the list.
// The n must not be nil.
func (l *List) MoveToFront(pre, n *Node) {
	 //TODO: Complete me!
}

// MoveToBack moves node n to the back of the list.
// The n must not be nil.
func (l *List) MoveToBack(pre, n *Node) {
	 //TODO: Complete me!
}

// String returns a string representation of the list
func (l *List) String() string {
	 //TODO: Complete me!
}

// Traversal traversals elements in the list, it will not stop until to the end of the list or the visitor returns false
func (l *List) Traversal(visitor visitor.Visitor) {
	 //TODO: Complete me!
}
