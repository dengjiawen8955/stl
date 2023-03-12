package treemap

import (
	"stl/ds/rbtree"
	"stl/utils/comparator"
	"stl/utils/sync"
	"stl/utils/visitor"
)

// MultiMap uses RbTress for internal data structure, and keys can bee repeated.
type MultiMap[K, V any] struct {
	 //TODO: Complete me!
}

//NewMultiMap creates a new MultiMap
func NewMultiMap[K, V any](cmp comparator.Comparator[K], opts ...Option) *MultiMap[K, V] {
	 //TODO: Complete me!
}

//Insert inserts a key-value to the MultiMap
func (mm *MultiMap[K, V]) Insert(key K, value V) {
	 //TODO: Complete me!
}

//Get returns the first node's value by the passed key if the key is in the MultiMap, otherwise returns nil
func (mm *MultiMap[K, V]) Get(key K) (V, error) {
	 //TODO: Complete me!
}

//Erase erases the key in the MultiMap
func (mm *MultiMap[K, V]) Erase(key K) {
	 //TODO: Complete me!
}

//Find finds the node by the passed key in the MultiMap and returns its iterator
func (mm *MultiMap[K, V]) Find(key K) *MapIterator[K, V] {
	 //TODO: Complete me!
}

//LowerBound find the first node that its key is equal or greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap[K, V]) LowerBound(key K) *MapIterator[K, V] {
	 //TODO: Complete me!
}

//UpperBound find the first node that its key is greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap[K, V]) UpperBound(key K) *MapIterator[K, V] {
	 //TODO: Complete me!
}

//Begin returns the first node's iterator
func (mm *MultiMap[K, V]) Begin() *MapIterator[K, V] {
	 //TODO: Complete me!
}

//First returns the first node's iterator
func (mm *MultiMap[K, V]) First() *MapIterator[K, V] {
	 //TODO: Complete me!
}

//Last returns the last node's iterator
func (mm *MultiMap[K, V]) Last() *MapIterator[K, V] {
	 //TODO: Complete me!
}

//Clear clears the MultiMap
func (mm *MultiMap[K, V]) Clear() {
	 //TODO: Complete me!
}

// Contains returns true if the passed value is in the MultiMap. otherwise returns false.
func (mm *MultiMap[K, V]) Contains(key K) bool {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the MultiMap
func (mm *MultiMap[K, V]) Size() int {
	 //TODO: Complete me!
}

// Traversal traversals elements in the MultiMap, it will not stop until to the end of the MultiMap or the visitor returns false
func (mm *MultiMap[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	 //TODO: Complete me!
}
