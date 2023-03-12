package rbtree

import (
	"stl/utils/iterator"
)

// RbTreeIterator is an iterator implementation of RbTree
type RbTreeIterator[K, V any] struct {
	 //TODO: Complete me!
}

// NewIterator creates a RbTreeIterator from the passed node
func NewIterator[K, V any](node *Node[K, V]) *RbTreeIterator[K, V] {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *RbTreeIterator[K, V]) IsValid() bool {
	 //TODO: Complete me!
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *RbTreeIterator[K, V]) Next() iterator.ConstIterator[V] {
	 //TODO: Complete me!
}

// Prev moves the pointer of the iterator to the previous node, and returns itself
func (iter *RbTreeIterator[K, V]) Prev() iterator.ConstBidIterator[V] {
	 //TODO: Complete me!
}

// Key returns the node's key of the iterator point to
func (iter *RbTreeIterator[K, V]) Key() K {
	 //TODO: Complete me!
}

// Value returns the node's value of the iterator point to
func (iter *RbTreeIterator[K, V]) Value() V {
	 //TODO: Complete me!
}

// SetValue sets the node's value of the iterator point to
func (iter *RbTreeIterator[K, V]) SetValue(val V) error {
	 //TODO: Complete me!
}

// Clone clones the iterator into a new RbTreeIterator
func (iter *RbTreeIterator[K, V]) Clone() iterator.ConstIterator[V] {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *RbTreeIterator[K, V]) Equal(other iterator.ConstIterator[V]) bool {
	 //TODO: Complete me!
}
