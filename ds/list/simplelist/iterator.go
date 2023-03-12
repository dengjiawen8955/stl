package simplelist

import "stl/utils/iterator"

//ListIterator is an implementation of Iterator
var _ iterator.Iterator[any] = (*ListIterator[any])(nil)

// ListIterator is an iterator for list
type ListIterator[T any] struct {
	 //TODO: Complete me!
}

// NewIterator news a ListIterator
func NewIterator[T any](node *Node[T]) *ListIterator[T] {
	 //TODO: Complete me!
}

// IsValid returns whether iter is valid
func (iter *ListIterator[T]) IsValid() bool {
	 //TODO: Complete me!
}

// Next returns the next iterator
func (iter *ListIterator[T]) Next() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Value returns the internal value of iter
func (iter *ListIterator[T]) Value() T {
	 //TODO: Complete me!
}

// SetValue sets the value of iter
func (iter *ListIterator[T]) SetValue(value T) {
	 //TODO: Complete me!
}

// Clone clones iter to a new ListIterator
func (iter *ListIterator[T]) Clone() iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// Equal returns whether iter is equal to other
func (iter *ListIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	 //TODO: Complete me!
}
