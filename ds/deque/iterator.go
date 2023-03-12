package deque

import (
	"stl/utils/iterator"
)

// DequeIterator is an implementation of RandomAccessIterator
type T any

var _ iterator.RandomAccessIterator[T] = (*DequeIterator[T])(nil)

// DequeIterator is an implementation of Deque iterator
type DequeIterator[T any] struct {
	 //TODO: Complete me!
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *DequeIterator[T]) IsValid() bool {
	 //TODO: Complete me!
}

// Value returns the value of the deque at the position of the iterator point to
func (iter *DequeIterator[T]) Value() T {
	 //TODO: Complete me!
}

// SetValue sets the value of the deque at the position of the iterator point to
func (iter *DequeIterator[T]) SetValue(val T) {
	 //TODO: Complete me!
}

// Next moves the position of the iterator to the next position and returns itself
func (iter *DequeIterator[T]) Next() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *DequeIterator[T]) Prev() iterator.ConstBidIterator[T] {
	 //TODO: Complete me!
}

// Clone clones the iterator to a new iterator
func (iter *DequeIterator[T]) Clone() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// IteratorAt creates a new iterator with the passed position
func (iter *DequeIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	 //TODO: Complete me!
}

// Position returns the position of iterator
func (iter *DequeIterator[T]) Position() int {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *DequeIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	 //TODO: Complete me!
}
