package rbtree

import (
	"errors"
	"fmt"

	"stl/utils/comparator"
	"stl/utils/visitor"
)

var ErrorNotFound = errors.New("not found")

// RbTree is a kind of self-balancing binary search tree in computer science.
// Each node of the binary tree has an extra bit, and that bit is often interpreted
// as the color (red or black) of the node. These color bits are used to ensure the tree
// remains approximately balanced during insertions and deletions.
type RbTree[K, V any] struct {
	 //TODO: Complete me!
}

// New creates a new RbTree
func New[K, V any](cmp comparator.Comparator[K]) *RbTree[K, V] {
	 //TODO: Complete me!
}

// Clear clears the RbTree
func (t *RbTree[K, V]) Clear() {
	 //TODO: Complete me!
}

// Find finds the first node that the key is equal to the passed key, and returns its value
func (t *RbTree[K, V]) Find(key K) (V, error) {
	 //TODO: Complete me!
}

// FindNode the first node that the key is equal to the passed key and return it
func (t *RbTree[K, V]) FindNode(key K) *Node[K, V] {
	 //TODO: Complete me!
}

// Begin returns the node with minimum key in the RbTree
func (t *RbTree[K, V]) Begin() *Node[K, V] {
	 //TODO: Complete me!
}

// First returns the node with minimum key in the RbTree
func (t *RbTree[K, V]) First() *Node[K, V] {
	 //TODO: Complete me!
}

// RBegin returns the Node with maximum key in the RbTree
func (t *RbTree[K, V]) RBegin() *Node[K, V] {
	 //TODO: Complete me!
}

// Last returns the Node with maximum key in the RbTree
func (t *RbTree[K, V]) Last() *Node[K, V] {
	 //TODO: Complete me!
}

// IterFirst returns the iterator of first node
func (t *RbTree[K, V]) IterFirst() *RbTreeIterator[K, V] {
	 //TODO: Complete me!
}

// IterLast returns the iterator of first node
func (t *RbTree[K, V]) IterLast() *RbTreeIterator[K, V] {
	 //TODO: Complete me!
}

// Empty returns true if Tree is empty,otherwise returns false.
func (t *RbTree[K, V]) Empty() bool {
	 //TODO: Complete me!
}

// Size returns the size of the rbtree.
func (t *RbTree[K, V]) Size() int {
	 //TODO: Complete me!
}

// Insert inserts a key-value pair into the RbTree.
func (t *RbTree[K, V]) Insert(key K, value V) {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) rbInsertFixup(z *Node[K, V]) {
	 //TODO: Complete me!
}

// Delete deletes node from the RbTree
func (t *RbTree[K, V]) Delete(node *Node[K, V]) {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) rbDeleteFixup(x, parent *Node[K, V]) {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) rbFixupLeft(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) rbFixupRight(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) leftRotate(x *Node[K, V]) {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) rightRotate(x *Node[K, V]) {
	 //TODO: Complete me!
}

// findNode finds the node that its key is equal to the passed key, and returns it.
func (t *RbTree[K, V]) findNode(key K) *Node[K, V] {
	 //TODO: Complete me!
}

// findNode finds the first node that its key is equal to the passed key, and returns it
func (t *RbTree[K, V]) findFirstNode(key K) *Node[K, V] {
	 //TODO: Complete me!
}

// FindLowerBoundNode finds the first node that its key is equal or greater than the passed key, and returns it
func (t *RbTree[K, V]) FindLowerBoundNode(key K) *Node[K, V] {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) findLowerBoundNode(x *Node[K, V], key K) *Node[K, V] {
	 //TODO: Complete me!
}

// FindUpperBoundNode finds the first node that its key is greater than the passed key, and returns it
func (t *RbTree[K, V]) FindUpperBoundNode(key K) *Node[K, V] {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) findUpperBoundNode(x *Node[K, V], key K) *Node[K, V] {
	 //TODO: Complete me!
}

// Traversal traversals elements in the RbTree, it will not stop until to the end of RbTree or the visitor returns false
func (t *RbTree[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	 //TODO: Complete me!
}

// IsRbTree is a function use to test whether t is a RbTree or not
func (t *RbTree[K, V]) IsRbTree() (bool, error) {
	 //TODO: Complete me!
}

func (t *RbTree[K, V]) test(n *Node[K, V]) (int, int, bool) {
	 //TODO: Complete me!
}

// getColor returns the node's color
func getColor[K, V any](n *Node[K, V]) Color {
	 //TODO: Complete me!
}
