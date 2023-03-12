package array

import (
	"stl/utils/iterator"
)

//ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator = (*ArrayIterator)(nil)

// ArrayIterator is an an implementation of Array iterator
type ArrayIterator struct {
	 //TODO: Complete me!
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *ArrayIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Value returns the value of array at the position of the iterator point to
func (iter *ArrayIterator) Value() interface{} {
	 //TODO: Complete me!
}

// SetValue sets the value of the array at the position of the iterator point to
func (iter *ArrayIterator) SetValue(val interface{}) {
	 //TODO: Complete me!
}

// Next moves the position of iterator to the next position and returns itself
func (iter *ArrayIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Prev moves the position of iterator to the previous position and returns itself
func (iter *ArrayIterator) Prev() iterator.ConstBidIterator {
	 //TODO: Complete me!
}

// Clone clones the iterator to a new iterator
func (iter *ArrayIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// IteratorAt creates a new iterator with position pos
func (iter *ArrayIterator) IteratorAt(pos int) iterator.RandomAccessIterator {
	 //TODO: Complete me!
}

// Position returns the position of the iterator
func (iter *ArrayIterator) Position() int {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ArrayIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
