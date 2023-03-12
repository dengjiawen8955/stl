package array

import (
	"stl/utils/iterator"
)

type T any

//ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator[T] = (*ArrayIterator[T])(nil)

// ArrayIterator is an implementation of Array iterator
type ArrayIterator[T any] struct {
	 //TODO: Complete me!
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *ArrayIterator[T]) IsValid() bool {
	 //TODO: Complete me!
}

// Value returns the value of array at the position of the iterator point to
func (iter *ArrayIterator[T]) Value() T {
	 //TODO: Complete me!
}

// SetValue sets the value of the array at the position of the iterator point to
func (iter *ArrayIterator[T]) SetValue(val T) {
	 //TODO: Complete me!
}

// Next moves the position of iterator to the next position and returns itself
func (iter *ArrayIterator[T]) Next() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Prev moves the position of iterator to the previous position and returns itself
func (iter *ArrayIterator[T]) Prev() iterator.ConstBidIterator[T] {
	 //TODO: Complete me!
}

// Clone clones the iterator to a new iterator
func (iter *ArrayIterator[T]) Clone() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// IteratorAt creates a new iterator with position pos
func (iter *ArrayIterator[T]) IteratorAt(pos int) iterator.RandomAccessIterator[T] {
	 //TODO: Complete me!
}

// Position returns the position of the iterator
func (iter *ArrayIterator[T]) Position() int {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ArrayIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	 //TODO: Complete me!
}
