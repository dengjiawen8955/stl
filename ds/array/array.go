package array

import (
	"fmt"
)

// Array is a fixed size slice
type Array struct {
	 //TODO: Complete me!
}

// New creates a new array with passed size
func New(size int) *Array {
	 //TODO: Complete me!
}

// NewFromArray creates a new array from another array, and copy its values
func NewFromArray(other *Array) *Array {
	 //TODO: Complete me!
}

// Fill fills Array a with value val
func (a *Array) Fill(val interface{}) {
	 //TODO: Complete me!
}

// Set sets value val to the position pos of the array
func (a *Array) Set(pos int, val interface{}) {
	 //TODO: Complete me!
}

// At returns the value at position pos in the array
func (a *Array) At(pos int) interface{} {
	 //TODO: Complete me!
}

// Front returns the first value in the array
func (a *Array) Front() interface{} {
	 //TODO: Complete me!
}

// Back returns the last value in the array
func (a *Array) Back() interface{} {
	 //TODO: Complete me!
}

// Size returns number of elements within the array
func (a *Array) Size() int {
	 //TODO: Complete me!
}

// Empty returns whether the array is empty or not
func (a *Array) Empty() bool {
	 //TODO: Complete me!
}

// SwapArray swaps the values of two arrays
func (a *Array) SwapArray(other *Array) {
	 //TODO: Complete me!
}

// Data returns the internal values of the array
func (a *Array) Data() []interface{} {
	 //TODO: Complete me!
}

// Begin returns an iterator of the array with the first position
func (a *Array) Begin() *ArrayIterator {
	 //TODO: Complete me!
}

// End returns an iterator of the array with the position a.Size()
func (a *Array) End() *ArrayIterator {
	 //TODO: Complete me!
}

// First returns an iterator of the array with the first position
func (a *Array) First() *ArrayIterator {
	 //TODO: Complete me!
}

// Last returns an iterator of the array with the last position
func (a *Array) Last() *ArrayIterator {
	 //TODO: Complete me!
}

// IterAt returns an iterator of the array with position pos
func (a *Array) IterAt(pos int) *ArrayIterator {
	 //TODO: Complete me!
}

// String returns a string representation of the array
func (a *Array) String() string {
	 //TODO: Complete me!
}
