package set

import (
	"fmt"
	gosync "sync"

	"stl/ds/rbtree"
	"stl/utils/comparator"
	"stl/utils/sync"
	"stl/utils/visitor"
)

// constants definition
const (
	Empty = true
)

var (
	defaultLocker sync.FakeLocker
)

// Options holds the Set's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function  type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set the set goroutine-safe
// Note that iterators are not goroutine safe, and it is useless to turn on the setting option here.
// so don't use iterators in multi goroutines
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// Set uses RbTress for internal data structure, and every key can must bee unique.
type Set[T any] struct {
	 //TODO: Complete me!
}

// New creates a new set
func New[T any](cmp comparator.Comparator[T], opts ...Option) *Set[T] {
	 //TODO: Complete me!
}

// Insert inserts an element to the set
func (s *Set[T]) Insert(element T) {
	 //TODO: Complete me!
}

// Erase erases an element from the set
func (s *Set[T]) Erase(element T) {
	 //TODO: Complete me!
}

// Find finds the element's node in the set, and return its iterator
func (s *Set[T]) Find(element T) *SetIterator[T] {
	 //TODO: Complete me!
}

// LowerBound finds the first element that equal or greater than the passed element in the set, and returns its iterator
func (s *Set[T]) LowerBound(element T) *SetIterator[T] {
	 //TODO: Complete me!
}

// UpperBound finds the first element that greater than the passed element in the set, and returns its iterator
func (s *Set[T]) UpperBound(element T) *SetIterator[T] {
	 //TODO: Complete me!
}

// Begin returns the iterator with the minimum element in the set
func (s *Set[T]) Begin() *SetIterator[T] {
	 //TODO: Complete me!
}

// First returns the iterator with the minimum element in the set
func (s *Set[T]) First() *SetIterator[T] {
	 //TODO: Complete me!
}

// Last returns the iterator with the maximum element in the set
func (s *Set[T]) Last() *SetIterator[T] {
	 //TODO: Complete me!
}

// Clear clears the set
func (s *Set[T]) Clear() {
	 //TODO: Complete me!
}

// Contains returns true if the passed element is in the Set. otherwise returns false.
func (s *Set[T]) Contains(element T) bool {
	 //TODO: Complete me!
}

// Size returns the amount of element in the set
func (s *Set[T]) Size() int {
	 //TODO: Complete me!
}

// Traversal traversals elements in the set, it will not stop until to the end of the set or the visitor returns false
func (s *Set[T]) Traversal(visitor visitor.Visitor[T]) {
	 //TODO: Complete me!
}

// String returns a string representation of the set
func (s *Set[T]) String() string {
	 //TODO: Complete me!
}

// Intersect returns a new set with the common elements in the set s and the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	 //TODO: Complete me!
}

// Union returns a new set with the all elements in the set s and the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	 //TODO: Complete me!
}

// Diff returns a new set with the elements in the set s but not in the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set[T]) Diff(other *Set[T]) *Set[T] {
	 //TODO: Complete me!
}
