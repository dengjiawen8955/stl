package set

import (
	"stl/ds/rbtree"
	"stl/utils/iterator"
)

// SetIterator is an iterator implementation of set
type SetIterator struct {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *SetIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Next moves the pointer of the iterator to the next node and returns itself
func (iter *SetIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Prev moves the pointer of the iterator to the previous node and returns itself
func (iter *SetIterator) Prev() iterator.ConstBidIterator {
	 //TODO: Complete me!
}

// Value returns the element of the iterator point to
func (iter *SetIterator) Value() interface{} {
	 //TODO: Complete me!
}

// Clone clones the iterator into a new SetIterator
func (iter *SetIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SetIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
