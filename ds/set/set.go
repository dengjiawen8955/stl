package set

import (
	"fmt"
	"stl/ds/rbtree"
	"stl/utils/comparator"
	"stl/utils/sync"
	"stl/utils/visitor"
	gosync "sync"
)

// constants definition
const (
	Empty = 0
)

var (
	defaultKeyComparator = comparator.BuiltinTypeComparator
	defaultLocker        sync.FakeLocker
)

// Options holds the Set's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function  type used to set Options
type Option func(option *Options)

// WithKeyComparator is used to set the key comparator for a set
func WithKeyComparator(cmp comparator.Comparator) Option {
	 //TODO: Complete me!
}

// WithGoroutineSafe is used to set the set goroutine-safe
// Note that iterators are not goroutine safe, and it is useless to turn on the setting option here.
// so don't use iterators in multi goroutines
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// Set uses RbTress for internal data structure, and every key can must bee unique.
type Set struct {
	 //TODO: Complete me!
}

// New creates a new set
func New(opts ...Option) *Set {
	 //TODO: Complete me!
}

// Insert inserts an element to the set
func (s *Set) Insert(element interface{}) {
	 //TODO: Complete me!
}

// Erase erases an element from the set
func (s *Set) Erase(element interface{}) {
	 //TODO: Complete me!
}

// Find finds the element's node in the set, and return its iterator
func (s *Set) Find(element interface{}) *SetIterator {
	 //TODO: Complete me!
}

// LowerBound finds the first element that equal or greater than the passed element in the set, and returns its iterator
func (s *Set) LowerBound(element interface{}) *SetIterator {
	 //TODO: Complete me!
}

// Begin returns the iterator with the minimum element in the set
func (s *Set) Begin() *SetIterator {
	 //TODO: Complete me!
}

// First returns the iterator with the minimum element in the set
func (s *Set) First() *SetIterator {
	 //TODO: Complete me!
}

// Last returns the iterator with the maximum element in the set
func (s *Set) Last() *SetIterator {
	 //TODO: Complete me!
}

// Clear clears the set
func (s *Set) Clear() {
	 //TODO: Complete me!
}

// Contains returns true if the passed element is in the Set. otherwise returns false.
func (s *Set) Contains(element interface{}) bool {
	 //TODO: Complete me!
}

// Size returns the amount of element in the set
func (s *Set) Size() int {
	 //TODO: Complete me!
}

// Traversal traversals elements in the set, it will not stop until to the end of the set or the visitor returns false
func (s *Set) Traversal(visitor visitor.Visitor) {
	 //TODO: Complete me!
}

// String returns a string representation of the set
func (s *Set) String() string {
	 //TODO: Complete me!
}

// Intersect returns a new set with the common elements in the set s and the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set) Intersect(other *Set) *Set {
	 //TODO: Complete me!
}

// Union returns a new set with the all elements in the set s and the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set) Union(other *Set) *Set {
	 //TODO: Complete me!
}

// Diff returns a new set with the elements in the set s but not in the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set) Diff(other *Set) *Set {
	 //TODO: Complete me!
}
