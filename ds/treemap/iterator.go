package treemap

import (
	"stl/ds/rbtree"
	"stl/utils/iterator"
)

// MapIterator is a map iterator
type MapIterator[K, V any] struct {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *MapIterator[K, V]) IsValid() bool {
	 //TODO: Complete me!
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *MapIterator[K, V]) Next() iterator.ConstIterator[V] {
	 //TODO: Complete me!
}

// Prev moves the pointer of the iterator to the previous node, and returns itseft
func (iter *MapIterator[K, V]) Prev() iterator.ConstBidIterator[V] {
	 //TODO: Complete me!
}

// Key returns the node's key of the iterator point to
func (iter *MapIterator[K, V]) Key() K {
	 //TODO: Complete me!
}

// Value returns the node's value of the iterator point to
func (iter *MapIterator[K, V]) Value() V {
	 //TODO: Complete me!
}

// SetValue sets the node's value of the iterator point to
func (iter *MapIterator[K, V]) SetValue(val V) {
	 //TODO: Complete me!
}

// Clone clones the iterator to a new MapIterator
func (iter *MapIterator[K, V]) Clone() iterator.ConstIterator[V] {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *MapIterator[K, V]) Equal(other iterator.ConstIterator[V]) bool {
	 //TODO: Complete me!
}
