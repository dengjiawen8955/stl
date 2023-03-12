package set

import (
	"fmt"
	"stl/ds/rbtree"
	"stl/utils/comparator"
	"stl/utils/sync"
	"stl/utils/visitor"
)

// MultiSet uses RbTress for internal data structure, and keys can bee repeated.
type MultiSet[T any] struct {
	 //TODO: Complete me!
}

// NewMultiSet creates a new MultiSet
func NewMultiSet[T any](cmp comparator.Comparator[T], opts ...Option) *MultiSet[T] {
	 //TODO: Complete me!
}

// Insert inserts an element to the MultiSet
func (ms *MultiSet[T]) Insert(element T) {
	 //TODO: Complete me!
}

// Erase erases all node with passed element in the MultiSet
func (ms *MultiSet[T]) Erase(element T) {
	 //TODO: Complete me!
}

// Find finds the first element that is equal to the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) Find(element T) *SetIterator[T] {
	 //TODO: Complete me!
}

//LowerBound finds the first element that is equal to or greater than the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) LowerBound(element T) *SetIterator[T] {
	 //TODO: Complete me!
}

//UpperBound finds the first element that is greater than the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) UpperBound(element T) *SetIterator[T] {
	 //TODO: Complete me!
}

// Begin returns the iterator with the minimum element in the MultiSet
func (ms *MultiSet[T]) Begin() *SetIterator[T] {
	 //TODO: Complete me!
}

// First returns the iterator with the minimum element in the MultiSet
func (ms *MultiSet[T]) First() *SetIterator[T] {
	 //TODO: Complete me!
}

//Last returns the iterator with the maximum element in the MultiSet
func (ms *MultiSet[T]) Last() *SetIterator[T] {
	 //TODO: Complete me!
}

// Clear clears all elements in the MultiSet
func (ms *MultiSet[T]) Clear() {
	 //TODO: Complete me!
}

// Contains returns true if the passed element is in the MultiSet. otherwise returns false.
func (ms *MultiSet[T]) Contains(element T) bool {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the MultiSet
func (ms *MultiSet[T]) Size() int {
	 //TODO: Complete me!
}

// Traversal traversals elements in the MultiSet, it will not stop until to the end of the MultiSet or the visitor returns false
func (ms *MultiSet[T]) Traversal(visitor visitor.Visitor[T]) {
	 //TODO: Complete me!
}

// String returns s string representation of the MultiSet
func (ms *MultiSet[T]) String() string {
	 //TODO: Complete me!
}
