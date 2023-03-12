package set

import (
	"fmt"
	"stl/ds/rbtree"
	"stl/utils/comparator"
	"stl/utils/sync"
	"stl/utils/visitor"
)

// MultiSet uses RbTress for internal data structure, and keys can bee repeated.
type MultiSet struct {
	 //TODO: Complete me!
}

// NewMultiSet creates a new MultiSet
func NewMultiSet(opts ...Option) *MultiSet {
	 //TODO: Complete me!
}

// Insert inserts an element to the MultiSet
func (ms *MultiSet) Insert(element interface{}) {
	 //TODO: Complete me!
}

// Erase erases all node with passed element in the MultiSet
func (ms *MultiSet) Erase(element interface{}) {
	 //TODO: Complete me!
}

// Find finds the first element that is equal to the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet) Find(element interface{}) *SetIterator {
	 //TODO: Complete me!
}

//LowerBound finds the first element that is equal to or greater than the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet) LowerBound(element interface{}) *SetIterator {
	 //TODO: Complete me!
}

// Begin returns the iterator with the minimum element in the MultiSet
func (ms *MultiSet) Begin() *SetIterator {
	 //TODO: Complete me!
}

// First returns the iterator with the minimum element in the MultiSet
func (ms *MultiSet) First() *SetIterator {
	 //TODO: Complete me!
}

//Last returns the iterator with the maximum element in the MultiSet
func (ms *MultiSet) Last() *SetIterator {
	 //TODO: Complete me!
}

// Clear clears all elements in the MultiSet
func (ms *MultiSet) Clear() {
	 //TODO: Complete me!
}

// Contains returns true if the passed element is in the MultiSet. otherwise returns false.
func (ms *MultiSet) Contains(element interface{}) bool {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the MultiSet
func (ms *MultiSet) Size() int {
	 //TODO: Complete me!
}

// Traversal traversals elements in the MultiSet, it will not stop until to the end of the MultiSet or the visitor returns false
func (ms *MultiSet) Traversal(visitor visitor.Visitor) {
	 //TODO: Complete me!
}

// String returns s string representation of the MultiSet
func (ms *MultiSet) String() string {
	 //TODO: Complete me!
}
