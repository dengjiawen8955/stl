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
type Vector struct {
	 //TODO: Complete me!
}

// New creates a new Vector
func New(opts ...Option) *Vector {
	 //TODO: Complete me!
}

// NewFromVector news a Vector from other Vector
func NewFromVector(other *Vector) *Vector {
	 //TODO: Complete me!
}

// Size returns the size of the vector
func (v *Vector) Size() int {
	 //TODO: Complete me!
}

// Capacity returns the capacity of the vector
func (v *Vector) Capacity() int {
	 //TODO: Complete me!
}

// Empty returns true if the vector is empty, otherwise returns false
func (v *Vector) Empty() bool {
	 //TODO: Complete me!
}

// PushBack pushes val to the back of the vector
func (v *Vector) PushBack(val interface{}) {
	 //TODO: Complete me!
}

// SetAt sets the value val to the vector at position pos
func (v *Vector) SetAt(pos int, val interface{}) {
	 //TODO: Complete me!
}

// InsertAt inserts the value val to the vector at position pos
func (v *Vector) InsertAt(pos int, val interface{}) {
	 //TODO: Complete me!
}

// EraseAt erases the value at position pos
func (v *Vector) EraseAt(pos int) {
	 //TODO: Complete me!
}

// EraseIndexRange erases values at range[first, last)
func (v *Vector) EraseIndexRange(first, last int) {
	 //TODO: Complete me!
}

// At returns the value at position pos, returns nil if pos is out off range .
func (v *Vector) At(pos int) interface{} {
	 //TODO: Complete me!
}

//Front returns the first value in the vector, returns nil if the vector is empty.
func (v *Vector) Front() interface{} {
	 //TODO: Complete me!
}

//Back returns the last value in the vector, returns nil if the vector is empty.
func (v *Vector) Back() interface{} {
	 //TODO: Complete me!
}

//PopBack returns the last value of the vector and erase it, returns nil if the vector is empty.
func (v *Vector) PopBack() interface{} {
	 //TODO: Complete me!
}

//Reserve makes a new space for the vector with passed capacity
func (v *Vector) Reserve(capacity int) {
	 //TODO: Complete me!
}

// ShrinkToFit shrinks the capacity of the vector to the fit size
func (v *Vector) ShrinkToFit() {
	 //TODO: Complete me!
}

// Clear clears all data in the vector
func (v *Vector) Clear() {
	 //TODO: Complete me!
}

// Data returns internal data of the vector
func (v *Vector) Data() []interface{} {
	 //TODO: Complete me!
}

// Begin returns the first iterator of the vector
func (v *Vector) Begin() *VectorIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of the vector
func (v *Vector) End() *VectorIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of the vector
func (v *Vector) First() *VectorIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of the vector
func (v *Vector) Last() *VectorIterator {
	 //TODO: Complete me!
}

// IterAt  returns the iterator at position of the vector
func (v *Vector) IterAt(pos int) *VectorIterator {
	 //TODO: Complete me!
}

// Insert inserts a value val to the vector at the position of the iterator iter point to
func (v *Vector) Insert(iter iterator.ConstIterator, val interface{}) *VectorIterator {
	 //TODO: Complete me!
}

// Erase erases the element of the iterator iter point to
func (v *Vector) Erase(iter iterator.ConstIterator) *VectorIterator {
	 //TODO: Complete me!
}

// EraseRange erases all elements in the range[first, last)
func (v *Vector) EraseRange(first, last iterator.ConstIterator) *VectorIterator {
	 //TODO: Complete me!
}

// Resize resizes the size of the vector to the passed size
func (v *Vector) Resize(size int) {
	 //TODO: Complete me!
}

// String returns a string representation of the vector
func (v *Vector) String() string {
	 //TODO: Complete me!
}
