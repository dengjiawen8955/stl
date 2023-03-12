package treemap

import (
	"stl/ds/rbtree"
	"stl/utils/comparator"
	"stl/utils/iterator"
	"stl/utils/sync"
	"stl/utils/visitor"
	gosync "sync"
)

var (
	defaultKeyComparator = comparator.BuiltinTypeComparator
	defaultLocker        sync.FakeLocker
)

// Options holds Map's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithKeyComparator is used to set the key comparator of map
func WithKeyComparator(cmp comparator.Comparator) Option {
	 //TODO: Complete me!
}

// WithGoroutineSafe is used to set a map goroutine-safe
// Note that iterators are not goroutine safe, and it is useless to turn on the setting option here.
// so don't use iterator in multi goroutines
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// Map uses RbTress for internal data structure, and every key can must bee unique.
type Map struct {
	 //TODO: Complete me!
}

// New creates a new map
func New(opts ...Option) *Map {
	 //TODO: Complete me!
}

//Insert inserts a key-value to the map
func (m *Map) Insert(key, value interface{}) {
	 //TODO: Complete me!
}

//Get returns the value of the passed key if the key is in the map, otherwise returns nil
func (m *Map) Get(key interface{}) interface{} {
	 //TODO: Complete me!
}

//Erase erases the node by the passed key from the map if the key in the Map
func (m *Map) Erase(key interface{}) {
	 //TODO: Complete me!
}

//EraseIter erases the node that iterator iter point to from the map
func (m *Map) EraseIter(iter iterator.ConstKvIterator) {
	 //TODO: Complete me!
}

//Find finds a node by the passed key and returns its iterator
func (m *Map) Find(key interface{}) *MapIterator {
	 //TODO: Complete me!
}

//LowerBound finds a node that its key is equal or greater than the passed key and returns its iterator
func (m *Map) LowerBound(key interface{}) *MapIterator {
	 //TODO: Complete me!
}

//Begin returns the first node's iterator
func (m *Map) Begin() *MapIterator {
	 //TODO: Complete me!
}

//First returns the first node's iterator
func (m *Map) First() *MapIterator {
	 //TODO: Complete me!
}

//Last returns the last node's iterator
func (m *Map) Last() *MapIterator {
	 //TODO: Complete me!
}

//Clear clears the map
func (m *Map) Clear() {
	 //TODO: Complete me!
}

// Contains returns true if the key is in the map. otherwise returns false.
func (m *Map) Contains(key interface{}) bool {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the map
func (m *Map) Size() int {
	 //TODO: Complete me!
}

// Traversal traversals elements in the map, it will not stop until to the end or the visitor returns false
func (m *Map) Traversal(visitor visitor.KvVisitor) {
	 //TODO: Complete me!
}
