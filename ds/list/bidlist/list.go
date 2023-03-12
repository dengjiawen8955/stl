package bidlist

import (
	"errors"
	"fmt"

	"stl/ds/container"
	"stl/utils/visitor"
)

// List is an implementation of Container
type T any

var ErrorOutOfRange = errors.New("out of range")

var _ container.Container[T] = (*List[T])(nil)

// Node is a list node
type Node[T any] struct {
	 //TODO: Complete me!
}

// Next returns the next list node or nil.
func (n *Node[T]) Next() *Node[T] {
	 //TODO: Complete me!
}

// Prev returns the previous list node or nil.
func (n *Node[T]) Prev() *Node[T] {
	 //TODO: Complete me!
}

// List represents a bidirectional list:
//
//	head -> node1 -- node2 --  node3
//	         |                   |
//	        node6 -- node5 --  node4
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

// Size returns the amount of list nodes.
func (l *List[T]) Size() int {
	 //TODO: Complete me!
}

// Empty returns true if the list is empty
func (l *List[T]) Empty() bool {
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

// Front returns the value of the front node
func (l *List[T]) Front() T {
	 //TODO: Complete me!
}

// Back returns the value of the last node
func (l *List[T]) Back() T {
	 //TODO: Complete me!
}

// PushBack inserts a new node n with value v at the back of the list
func (l *List[T]) PushBack(v T) {
	 //TODO: Complete me!
}

// PushBack inserts a new node n with value v at the back of the list and returns n.
func (l *List[T]) pushBack(v T) *Node[T] {
	 //TODO: Complete me!
}

// PushFront inserts a new node n with value v at the front of the list.
func (l *List[T]) PushFront(v T) {
	 //TODO: Complete me!
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Node[T]) *Node[T] {
	 //TODO: Complete me!
}

// InsertBefore inserts a new node n with value v immediately before mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertBefore(v T, mark *Node[T]) *Node[T] {
	 //TODO: Complete me!
}

func (l *List[T]) insertAfter(n, at *Node[T]) *Node[T] {
	 //TODO: Complete me!
}

// Remove removes n from l list if n is a node of l list.
// It returns the n value n.Value.
// The node must not be nil.
func (l *List[T]) Remove(n *Node[T]) T {
	 //TODO: Complete me!
}

func (l *List[T]) remove(n *Node[T]) *Node[T] {
	 //TODO: Complete me!
}

// Clear removes all nodes
func (l *List[T]) Clear() {
	 //TODO: Complete me!
}

// PopBack removes the last node in the list and returns its value
func (l *List[T]) PopBack() T {
	 //TODO: Complete me!
}

// PopFront removes the first node in the list and returns its value
func (l *List[T]) PopFront() T {
	 //TODO: Complete me!
}

// MoveToFront moves node n to the front of the list.
// If n is not a node of the list, the list is not modified.
// The n must not be nil.
func (l *List[T]) MoveToFront(n *Node[T]) {
	 //TODO: Complete me!
}

// MoveToBack moves node  n to the back of the list.
// If e is not a node of the list, the list is not modified.
// The node must not be nil.
func (l *List[T]) MoveToBack(n *Node[T]) {
	 //TODO: Complete me!
}

// MoveAfter moves node n to its new position after mark.
// If n or mark is not a node of the list, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (l *List[T]) MoveAfter(n, mark *Node[T]) {
	 //TODO: Complete me!
}

func (l *List[T]) moveToAfter(n, at *Node[T]) {
	 //TODO: Complete me!
}

// PushBackList inserts a copy of an other list at the back of the list.
// The list and other may be the same. They must not be nil.
func (l *List[T]) PushBackList(other *List[T]) {
	 //TODO: Complete me!
}

// PushFrontList inserts a copy of another list at the front of the list.
// The list and other may be the same. They must not be nil.
func (l *List[T]) PushFrontList(other *List[T]) {
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
