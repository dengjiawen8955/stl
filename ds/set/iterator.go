package set

import (
	"stl/ds/rbtree"
	"stl/utils/iterator"
)

// SetIterator is an iterator implementation of set
type SetIterator[T any] struct {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *SetIterator[K]) IsValid() bool {
	 //TODO: Complete me!
}

// Next moves the pointer of the iterator to the next node and returns itself
func (iter *SetIterator[T]) Next() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Prev moves the pointer of the iterator to the previous node and returns itself
func (iter *SetIterator[T]) Prev() iterator.ConstBidIterator[T] {
	 //TODO: Complete me!
}

// Value returns the element of the iterator point to
func (iter *SetIterator[T]) Value() T {
	 //TODO: Complete me!
}

// Clone clones the iterator into a new SetIterator
func (iter *SetIterator[T]) Clone() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SetIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	 //TODO: Complete me!
}
