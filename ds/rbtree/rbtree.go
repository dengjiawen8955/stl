package rbtree

import (
	"fmt"
	"stl/utils/comparator"
	"stl/utils/visitor"
)

var (
	defaultKeyComparator = comparator.BuiltinTypeComparator
)

// Options holds RbTree's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

//WithKeyComparator is used to set RbTree's key comparator
func WithKeyComparator(cmp comparator.Comparator) Option {
	 //TODO: Complete me!
}

// RbTree is a kind of self-balancing binary search tree in computer science.
// Each node of the binary tree has an extra bit, and that bit is often interpreted
// as the color (red or black) of the node. These color bits are used to ensure the tree
// remains approximately balanced during insertions and deletions.
type RbTree struct {
	 //TODO: Complete me!
}

//New creates a new RbTree
func New(opts ...Option) *RbTree {
	 //TODO: Complete me!
}

// Clear clears the RbTree
func (t *RbTree) Clear() {
	 //TODO: Complete me!
}

// Find finds the first node that the key is equal to the passed key, and returns its value
func (t *RbTree) Find(key interface{}) interface{} {
	 //TODO: Complete me!
}

// FindNode the first node that the key is equal to the passed key and return it
func (t *RbTree) FindNode(key interface{}) *Node {
	 //TODO: Complete me!
}

// Begin returns the node with minimum key in the RbTree
func (t *RbTree) Begin() *Node {
	 //TODO: Complete me!
}

// First returns the node with minimum key in the RbTree
func (t *RbTree) First() *Node {
	 //TODO: Complete me!
}

// RBegin returns the Node with maximum key in the RbTree
func (t *RbTree) RBegin() *Node {
	 //TODO: Complete me!
}

// Last returns the Node with maximum key in the RbTree
func (t *RbTree) Last() *Node {
	 //TODO: Complete me!
}

// IterFirst returns the iterator of first node
func (t *RbTree) IterFirst() *RbTreeIterator {
	 //TODO: Complete me!
}

// IterLast returns the iterator of first node
func (t *RbTree) IterLast() *RbTreeIterator {
	 //TODO: Complete me!
}

// Empty returns true if Tree is empty,otherwise returns false.
func (t *RbTree) Empty() bool {
	 //TODO: Complete me!
}

// Size returns the size of the rbtree.
func (t *RbTree) Size() int {
	 //TODO: Complete me!
}

// Insert inserts a key-value pair into the RbTree.
func (t *RbTree) Insert(key, value interface{}) {
	 //TODO: Complete me!
}

func (t *RbTree) rbInsertFixup(z *Node) {
	 //TODO: Complete me!
}

// Delete deletes node from the RbTree
func (t *RbTree) Delete(node *Node) {
	 //TODO: Complete me!
}

func (t *RbTree) rbDeleteFixup(x, parent *Node) {
	 //TODO: Complete me!
}

func (t *RbTree) rbFixupLeft(x, parent, w *Node) (*Node, *Node) {
	 //TODO: Complete me!
}

func (t *RbTree) rbFixupRight(x, parent, w *Node) (*Node, *Node) {
	 //TODO: Complete me!
}

func (t *RbTree) leftRotate(x *Node) {
	 //TODO: Complete me!
}

func (t *RbTree) rightRotate(x *Node) {
	 //TODO: Complete me!
}

// findNode finds the node that its key is equal to the passed key, and returns it.
func (t *RbTree) findNode(key interface{}) *Node {
	 //TODO: Complete me!
}

// findNode finds the first node that its key is equal to the passed key, and returns it
func (t *RbTree) findFirstNode(key interface{}) *Node {
	 //TODO: Complete me!
}

// FindLowerBoundNode finds the first node that its key is equal or greater than the passed key, and returns it
func (t *RbTree) FindLowerBoundNode(key interface{}) *Node {
	 //TODO: Complete me!
}

func (t *RbTree) findLowerBoundNode(x *Node, key interface{}) *Node {
	 //TODO: Complete me!
}

// Traversal traversals elements in the RbTree, it will not stop until to the end of RbTree or the visitor returns false
func (t *RbTree) Traversal(visitor visitor.KvVisitor) {
	 //TODO: Complete me!
}

// IsRbTree is a function use to test whether t is a RbTree or not
func (t *RbTree) IsRbTree() (bool, error) {
	 //TODO: Complete me!
}

func (t *RbTree) test(n *Node) (int, int, bool) {
	 //TODO: Complete me!
}

// getColor returns the node's color
func getColor(n *Node) Color {
	 //TODO: Complete me!
}
