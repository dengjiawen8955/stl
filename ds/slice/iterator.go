package slice

import "stl/utils/iterator"

type T any

// SliceIterator is an implementation of RandomAccessIterator

var _ iterator.RandomAccessIterator[T] = (*SliceIterator[T])(nil)

// SliceIterator represents a slice iterator
type SliceIterator[T any] struct {
	 //TODO: Complete me!
}

// IsValid returns trus if the iterator is valid, othterwise return false
func (iter *SliceIterator[T]) IsValid() bool {
	 //TODO: Complete me!
}

// Value returns the value of the iterator point to
func (iter *SliceIterator[T]) Value() T {
	 //TODO: Complete me!
}

// SetValue sets the value of the iterator point to
func (iter *SliceIterator[T]) SetValue(val T) {
	 //TODO: Complete me!
}

// Next moves the iterator's position to the next position, and returns itself
func (iter *SliceIterator[T]) Next() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Prev move the iterator's position to the previous position, and return itself
func (iter *SliceIterator[T]) Prev() iterator.ConstBidIterator[T] {
	 //TODO: Complete me!
}

// Clone clones the iterator into a new one
func (iter *SliceIterator[T]) Clone() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// IteratorAt creates an iterator with the passed position
func (iter *SliceIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	 //TODO: Complete me!
}

// Position returns the position of the iterator
func (iter *SliceIterator[T]) Position() int {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SliceIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	 //TODO: Complete me!
}
