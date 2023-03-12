package bidlist

import (
	"fmt"
	"stl/ds/container"
	"stl/utils/visitor"
)

// List is an implementation of Container
var _ container.Container = (*List)(nil)

// Node is a list node
type Node struct {
	 //TODO: Complete me!
}

// Next returns the next list node or nil.
func (n *Node) Next() *Node {
	 //TODO: Complete me!
}

// Prev returns the previous list node or nil.
func (n *Node) Prev() *Node {
	 //TODO: Complete me!
}

// List represents a bidirectional list:
//
//   head -> node1 -- node2 --  node3
//            |                   |
//           node6 -- node5 --  node4
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

// Size returns the amount of list nodes.
func (l *List) Size() int {
	 //TODO: Complete me!
}

// Empty returns true if the list is empty
func (l *List) Empty() bool {
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

// Front returns the value of the front node
func (l *List) Front() interface{} {
	 //TODO: Complete me!
}

// Back returns the value of the last node
func (l *List) Back() interface{} {
	 //TODO: Complete me!
}

// PushBack inserts a new node n with value v at the back of the list
func (l *List) PushBack(v interface{}) {
	 //TODO: Complete me!
}

// PushBack inserts a new node n with value v at the back of the list and returns n.
func (l *List) pushBack(v interface{}) *Node {
	 //TODO: Complete me!
}

// PushFront inserts a new node n with value v at the front of the list.
func (l *List) PushFront(v interface{}) {
	 //TODO: Complete me!
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	 //TODO: Complete me!
}

// InsertBefore inserts a new node n with value v immediately before mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
	 //TODO: Complete me!
}

func (l *List) insertAfter(n, at *Node) *Node {
	 //TODO: Complete me!
}

// Remove removes n from l list if n is a node of l list.
// It returns the n value n.Value.
// The node must not be nil.
func (l *List) Remove(n *Node) interface{} {
	 //TODO: Complete me!
}

func (l *List) remove(n *Node) *Node {
	 //TODO: Complete me!
}

// Clear removes all nodes
func (l *List) Clear() {
	 //TODO: Complete me!
}

// PopBack removes the last node in the list and returns its value
func (l *List) PopBack() interface{} {
	 //TODO: Complete me!
}

// PopFront removes the first node in the list and returns its value
func (l *List) PopFront() interface{} {
	 //TODO: Complete me!
}

// MoveToFront moves node n to the front of the list.
// If n is not a node of the list, the list is not modified.
// The n must not be nil.
func (l *List) MoveToFront(n *Node) {
	 //TODO: Complete me!
}

// MoveToBack moves node  n to the back of the list.
// If e is not a node of the list, the list is not modified.
// The node must not be nil.
func (l *List) MoveToBack(n *Node) {
	 //TODO: Complete me!
}

// MoveAfter moves node n to its new position after mark.
// If n or mark is not a node of the list, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (l *List) MoveAfter(n, mark *Node) {
	 //TODO: Complete me!
}

func (l *List) moveToAfter(n, at *Node) {
	 //TODO: Complete me!
}

// PushBackList inserts a copy of an other list at the back of the list.
// The list and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List) {
	 //TODO: Complete me!
}

// PushFrontList inserts a copy of an other list at the front of the list.
// The list and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
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
