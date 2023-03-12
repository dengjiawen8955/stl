package treemap

import (
	"errors"
	"stl/ds/rbtree"
	"stl/utils/comparator"
	"stl/utils/iterator"
	"stl/utils/sync"
	"stl/utils/visitor"
	gosync "sync"
)

var (
	defaultLocker sync.FakeLocker
)

var ErrorNotFound = errors.New("not found")

// Options holds Map's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set a map goroutine-safe
// Note that iterators are not goroutine safe, and it is useless to turn on the setting option here.
// so don't use iterator in multi goroutines
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// Map uses RbTress for internal data structure, and every key can must bee unique.
type Map[K, V any] struct {
	 //TODO: Complete me!
}

// New creates a new map
func New[K, V any](cmp comparator.Comparator[K], opts ...Option) *Map[K, V] {
	 //TODO: Complete me!
}

//Insert inserts a key-value to the map
func (m *Map[K, V]) Insert(key K, value V) {
	 //TODO: Complete me!
}

//Get returns the value of the passed key if the key is in the map, otherwise returns nil
func (m *Map[K, V]) Get(key K) (V, error) {
	 //TODO: Complete me!
}

//Erase erases the node by the passed key from the map if the key in the Map
func (m *Map[K, V]) Erase(key K) {
	 //TODO: Complete me!
}

//EraseIter erases the node that iterator iter point to from the map
func (m *Map[K, V]) EraseIter(iter iterator.ConstKvIterator[K, V]) {
	 //TODO: Complete me!
}

//Find finds a node by the passed key and returns its iterator
func (m *Map[K, V]) Find(key K) *MapIterator[K, V] {
	 //TODO: Complete me!
}

//LowerBound finds a node that its key is equal or greater than the passed key and returns its iterator
func (m *Map[K, V]) LowerBound(key K) *MapIterator[K, V] {
	 //TODO: Complete me!
}

//UpperBound finds a node that its key is greater than the passed key and returns its iterator
func (m *Map[K, V]) UpperBound(key K) *MapIterator[K, V] {
	 //TODO: Complete me!
}

//Begin returns the first node's iterator
func (m *Map[K, V]) Begin() *MapIterator[K, V] {
	 //TODO: Complete me!
}

//First returns the first node's iterator
func (m *Map[K, V]) First() *MapIterator[K, V] {
	 //TODO: Complete me!
}

//Last returns the last node's iterator
func (m *Map[K, V]) Last() *MapIterator[K, V] {
	 //TODO: Complete me!
}

//Clear clears the map
func (m *Map[K, V]) Clear() {
	 //TODO: Complete me!
}

// Contains returns true if the key is in the map. otherwise returns false.
func (m *Map[K, V]) Contains(key K) bool {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the map
func (m *Map[K, V]) Size() int {
	 //TODO: Complete me!
}

// Traversal traversals elements in the map, it will not stop until to the end or the visitor returns false
func (m *Map[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	 //TODO: Complete me!
}
