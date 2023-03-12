package slice

import "stl/utils/iterator"

// SliceIterator is a implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator = (*SliceIterator)(nil)

// SliceIterator represents a slice iterator
type SliceIterator struct {
	 //TODO: Complete me!
}

// IsValid returns trus if the iterator is valid, othterwise return false
func (iter *SliceIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Value returns the value of the iterator point to
func (iter *SliceIterator) Value() interface{} {
	 //TODO: Complete me!
}

// SetValue sets the value of the iterator point to
func (iter *SliceIterator) SetValue(val interface{}) {
	 //TODO: Complete me!
}

// Next moves the iterator's position to the next position, and returns itself
func (iter *SliceIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Prev move the iterator's position to the previous position, and return itself
func (iter *SliceIterator) Prev() iterator.ConstBidIterator {
	 //TODO: Complete me!
}

// Clone clones the iterator into a new one
func (iter *SliceIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// IteratorAt creates an iterator with the passed position
func (iter *SliceIterator) IteratorAt(position int) iterator.RandomAccessIterator {
	 //TODO: Complete me!
}

// Position returns the position of the iterator
func (iter *SliceIterator) Position() int {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SliceIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
