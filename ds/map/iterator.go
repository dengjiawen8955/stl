package treemap

import (
	"stl/ds/rbtree"
	"stl/utils/iterator"
)

// MapIterator is a map iterator
type MapIterator struct {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *MapIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *MapIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Prev moves the pointer of the iterator to the previous node, and returns itseft
func (iter *MapIterator) Prev() iterator.ConstBidIterator {
	 //TODO: Complete me!
}

// Key returns the node's key of the iterator point to
func (iter *MapIterator) Key() interface{} {
	 //TODO: Complete me!
}

// Value returns the node's value of the iterator point to
func (iter *MapIterator) Value() interface{} {
	 //TODO: Complete me!
}

// SetValue sets the node's value of the iterator point to
func (iter *MapIterator) SetValue(val interface{}) {
	 //TODO: Complete me!
}

// Clone clones the iterator to a new MapIterator
func (iter *MapIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *MapIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
