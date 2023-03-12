package vector

import (
	"fmt"
	"stl/utils/iterator"
)

// Options holds the Vector's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithCapacity is used to set the capacity of a Vector
func WithCapacity(capacity int) Option {
	 //TODO: Complete me!
}

// Vector is a linear data structure, the internal is a slice
type Vector[T any] struct {
	 //TODO: Complete me!
}

// New creates a new Vector
func New[T any](opts ...Option) *Vector[T] {
	 //TODO: Complete me!
}

// NewFromVector news a Vector from other Vector
func NewFromVector[T any](other *Vector[T]) *Vector[T] {
	 //TODO: Complete me!
}

// Size returns the size of the vector
func (v *Vector[T]) Size() int {
	 //TODO: Complete me!
}

// Capacity returns the capacity of the vector
func (v *Vector[T]) Capacity() int {
	 //TODO: Complete me!
}

// Empty returns true if the vector is empty, otherwise returns false
func (v *Vector[T]) Empty() bool {
	 //TODO: Complete me!
}

// PushBack pushes val to the back of the vector
func (v *Vector[T]) PushBack(val T) {
	 //TODO: Complete me!
}

// SetAt sets the value val to the vector at position pos
func (v *Vector[T]) SetAt(pos int, val T) {
	 //TODO: Complete me!
}

// InsertAt inserts the value val to the vector at position pos
func (v *Vector[T]) InsertAt(pos int, val T) {
	 //TODO: Complete me!
}

// EraseAt erases the value at position pos
func (v *Vector[T]) EraseAt(pos int) {
	 //TODO: Complete me!
}

// EraseIndexRange erases values at range[first, last)
func (v *Vector[T]) EraseIndexRange(first, last int) {
	 //TODO: Complete me!
}

// At returns the value at position pos, returns nil if pos is out off range .
func (v *Vector[T]) At(pos int) T {
	 //TODO: Complete me!
}

//Front returns the first value in the vector, returns nil if the vector is empty.
func (v *Vector[T]) Front() T {
	 //TODO: Complete me!
}

//Back returns the last value in the vector, returns nil if the vector is empty.
func (v *Vector[T]) Back() T {
	 //TODO: Complete me!
}

//PopBack returns the last value of the vector and erase it, returns nil if the vector is empty.
func (v *Vector[T]) PopBack() T {
	 //TODO: Complete me!
}

//Reserve makes a new space for the vector with passed capacity
func (v *Vector[T]) Reserve(capacity int) {
	 //TODO: Complete me!
}

// ShrinkToFit shrinks the capacity of the vector to the fit size
func (v *Vector[T]) ShrinkToFit() {
	 //TODO: Complete me!
}

// Clear clears all data in the vector
func (v *Vector[T]) Clear() {
	 //TODO: Complete me!
}

// Data returns internal data of the vector
func (v *Vector[T]) Data() []T {
	 //TODO: Complete me!
}

// Begin returns the first iterator of the vector
func (v *Vector[T]) Begin() *VectorIterator[T] {
	 //TODO: Complete me!
}

// End returns the end iterator of the vector
func (v *Vector[T]) End() *VectorIterator[T] {
	 //TODO: Complete me!
}

// First returns the first iterator of the vector
func (v *Vector[T]) First() *VectorIterator[T] {
	 //TODO: Complete me!
}

// Last returns the last iterator of the vector
func (v *Vector[T]) Last() *VectorIterator[T] {
	 //TODO: Complete me!
}

// IterAt  returns the iterator at position of the vector
func (v *Vector[T]) IterAt(pos int) *VectorIterator[T] {
	 //TODO: Complete me!
}

// Insert inserts a value val to the vector at the position of the iterator iter point to
func (v *Vector[T]) Insert(iter iterator.ConstIterator[T], val T) *VectorIterator[T] {
	 //TODO: Complete me!
}

// Erase erases the element of the iterator iter point to
func (v *Vector[T]) Erase(iter iterator.ConstIterator[T]) *VectorIterator[T] {
	 //TODO: Complete me!
}

// EraseRange erases all elements in the range[first, last)
func (v *Vector[T]) EraseRange(first, last iterator.ConstIterator[T]) *VectorIterator[T] {
	 //TODO: Complete me!
}

// Resize resizes the size of the vector to the passed size
func (v *Vector[T]) Resize(size int) {
	 //TODO: Complete me!
}

// String returns a string representation of the vector
func (v *Vector[T]) String() string {
	 //TODO: Complete me!
}
