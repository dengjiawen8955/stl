package rbtree

import (
	"stl/utils/iterator"
)

// RbTreeIterator is an iterator implementation of RbTree
type RbTreeIterator struct {
	 //TODO: Complete me!
}

// NewIterator creates a RbTreeIterator from the passed node
func NewIterator(node *Node) *RbTreeIterator {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *RbTreeIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *RbTreeIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Prev moves the pointer of the iterator to the previous node, and returns itself
func (iter *RbTreeIterator) Prev() iterator.ConstBidIterator {
	 //TODO: Complete me!
}

// Key returns the node's key of the iterator point to
func (iter *RbTreeIterator) Key() interface{} {
	 //TODO: Complete me!
}

// Value returns the node's value of the iterator point to
func (iter *RbTreeIterator) Value() interface{} {
	 //TODO: Complete me!
}

//SetValue sets the node's value of the iterator point to
func (iter *RbTreeIterator) SetValue(val interface{}) error {
	 //TODO: Complete me!
}

// Clone clones the iterator into a new RbTreeIterator
func (iter *RbTreeIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *RbTreeIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
