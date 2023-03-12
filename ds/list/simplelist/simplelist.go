package simplelist

import (
	"fmt"
	"stl/utils/visitor"
)

// Node is a list node
type Node[T any] struct {
	 //TODO: Complete me!
}

// Next returns the next list node or nil.
func (n *Node[T]) Next() *Node[T] {
	 //TODO: Complete me!
}

// List represents a single direction list:
//
//   head -> node1 --> node2 --> node3 <- tail
//
type List[T any] struct {
	 //TODO: Complete me!
}

// New creates a list
func New[T any]() *List[T] {
	 //TODO: Complete me!
}

// Len returns the amount of list nodes.
func (l *List[T]) Len() int {
	 //TODO: Complete me!
}

// FrontNode returns the front node of the list or nil if the list is empty
func (l *List[T]) FrontNode() *Node[T] {
	 //TODO: Complete me!
}

// BackNode returns the last node of the list or nil if the list is empty
func (l *List[T]) BackNode() *Node[T] {
	 //TODO: Complete me!
}

// PushFront inserts a new node n with value v at the front of the list.
func (l *List[T]) PushFront(v T) {
	 //TODO: Complete me!
}

// PushBack inserts a new node n with value v at the back of the list.
func (l *List[T]) PushBack(v T) {
	 //TODO: Complete me!
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of the list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Node[T]) *Node[T] {
	 //TODO: Complete me!
}

func (l *List[T]) insertAfter(n, at *Node[T]) *Node[T] {
	 //TODO: Complete me!
}

// Remove removes node n from the list.
// The node must not be nil.
func (l *List[T]) Remove(pre, n *Node[T]) T {
	 //TODO: Complete me!
}

// MoveToFront moves node n to the front of the list.
// The n must not be nil.
func (l *List[T]) MoveToFront(pre, n *Node[T]) {
	 //TODO: Complete me!
}

// MoveToBack moves node n to the back of the list.
// The n must not be nil.
func (l *List[T]) MoveToBack(pre, n *Node[T]) {
	 //TODO: Complete me!
}

// String returns a string representation of the list
func (l *List[T]) String() string {
	 //TODO: Complete me!
}

// Traversal traversals elements in the list, it will not stop until to the end of the list or the visitor returns false
func (l *List[T]) Traversal(visitor visitor.Visitor[T]) {
	 //TODO: Complete me!
}
