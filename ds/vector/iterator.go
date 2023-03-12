package vector

import (
	"stl/utils/iterator"
)

//ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator = (*VectorIterator)(nil)

// VectorIterator represents a vector iterator
type VectorIterator struct {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *VectorIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Value returns the value of the iterator point to
func (iter *VectorIterator) Value() interface{} {
	 //TODO: Complete me!
}

// SetValue sets the value of the iterator point to
func (iter *VectorIterator) SetValue(val interface{}) {
	 //TODO: Complete me!
}

// Next moves the position of iterator to the next position and returns itself
func (iter *VectorIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *VectorIterator) Prev() iterator.ConstBidIterator {
	 //TODO: Complete me!
}

// Clone clones the iterator into a new iterator
func (iter *VectorIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// IteratorAt creates an iterator with the passed position
func (iter *VectorIterator) IteratorAt(position int) iterator.RandomAccessIterator {
	 //TODO: Complete me!
}

// Position return the position of the iterator point to
func (iter *VectorIterator) Position() int {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *VectorIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
