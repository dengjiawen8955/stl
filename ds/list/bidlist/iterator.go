package bidlist

import (
	"stl/utils/iterator"
)

// ListIterator is an implementation of BidIterator
var _ iterator.BidIterator[T] = (*ListIterator[T])(nil)

// ListIterator is an implementation of list iterator
type ListIterator[T any] struct {
	 //TODO: Complete me!
}

// NewIterator creates a ListIterator
func NewIterator[T any](node *Node[T]) *ListIterator[T] {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *ListIterator[T]) IsValid() bool {
	 //TODO: Complete me!
}

// Next moves the pointer of iterator to the next node and returns itself
func (iter *ListIterator[T]) Next() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Prev moves the pointer of iterator to the previous node and returns itself
func (iter *ListIterator[T]) Prev() iterator.ConstBidIterator[T] {
	 //TODO: Complete me!
}

// Value returns the node's value of the iterator point to
func (iter *ListIterator[T]) Value() T {
	 //TODO: Complete me!
}

// SetValue sets the node's value of the iterator point to
func (iter *ListIterator[T]) SetValue(value T) {
	 //TODO: Complete me!
}

// Clone clones the iterator to a new iterator
func (iter *ListIterator[T]) Clone() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ListIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	 //TODO: Complete me!
}
