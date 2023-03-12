package vector

import (
	"stl/utils/iterator"
)

type T any

//ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator[T] = (*VectorIterator[T])(nil)

// VectorIterator represents a vector iterator
type VectorIterator[T any] struct {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *VectorIterator[T]) IsValid() bool {
	 //TODO: Complete me!
}

// Value returns the value of the iterator point to
func (iter *VectorIterator[T]) Value() T {
	 //TODO: Complete me!
}

// SetValue sets the value of the iterator point to
func (iter *VectorIterator[T]) SetValue(val T) {
	 //TODO: Complete me!
}

// Next moves the position of iterator to the next position and returns itself
func (iter *VectorIterator[T]) Next() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *VectorIterator[T]) Prev() iterator.ConstBidIterator[T] {
	 //TODO: Complete me!
}

// Clone clones the iterator into a new iterator
func (iter *VectorIterator[T]) Clone() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// IteratorAt creates an iterator with the passed position
func (iter *VectorIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	 //TODO: Complete me!
}

// Position return the position of the iterator point to
func (iter *VectorIterator[T]) Position() int {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *VectorIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	 //TODO: Complete me!
}
