package deque

import (
	"stl/utils/iterator"
)

// DequeIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator = (*DequeIterator)(nil)

// DequeIterator is an implementation of Deque iterator
type DequeIterator struct {
	 //TODO: Complete me!
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *DequeIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Value returns the value of the deque at the position of the iterator point to
func (iter *DequeIterator) Value() interface{} {
	 //TODO: Complete me!
}

// SetValue sets the value of the deque at the position of the iterator point to
func (iter *DequeIterator) SetValue(val interface{}) {
	 //TODO: Complete me!
}

// Next moves the position of the iterator to the next position and returns itself
func (iter *DequeIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *DequeIterator) Prev() iterator.ConstBidIterator {
	 //TODO: Complete me!
}

// Clone clones the iterator to a new iterator
func (iter *DequeIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// IteratorAt creates a new iterator with the passed position
func (iter *DequeIterator) IteratorAt(position int) iterator.RandomAccessIterator {
	 //TODO: Complete me!
}

// Position returns the position of iterator
func (iter *DequeIterator) Position() int {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *DequeIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
