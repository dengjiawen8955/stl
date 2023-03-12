package array

import (
	"fmt"
)

// Array is a fixed size slice
type Array[T any] struct {
	 //TODO: Complete me!
}

// New creates a new array with passed size
func New[T any](size int) *Array[T] {
	 //TODO: Complete me!
}

// NewFromArray creates a new array from another array, and copy its values
func NewFromArray[T any](other *Array[T]) *Array[T] {
	 //TODO: Complete me!
}

// Fill fills Array a with value val
func (a *Array[T]) Fill(val T) {
	 //TODO: Complete me!
}

// Set sets value val to the position pos of the array
func (a *Array[T]) Set(pos int, val T) {
	 //TODO: Complete me!
}

// At returns the value at position pos in the array
func (a *Array[T]) At(pos int) T {
	 //TODO: Complete me!
}

// Front returns the first value in the array
func (a *Array[T]) Front() T {
	 //TODO: Complete me!
}

// Back returns the last value in the array
func (a *Array[T]) Back() T {
	 //TODO: Complete me!
}

// Size returns number of elements within the array
func (a *Array[T]) Size() int {
	 //TODO: Complete me!
}

// Empty returns whether the array is empty or not
func (a *Array[T]) Empty() bool {
	 //TODO: Complete me!
}

// SwapArray swaps the values of two arrays
func (a *Array[T]) SwapArray(other *Array[T]) {
	 //TODO: Complete me!
}

// Data returns the internal values of the array
func (a *Array[T]) Data() []T {
	 //TODO: Complete me!
}

// Begin returns an iterator of the array with the first position
func (a *Array[T]) Begin() *ArrayIterator[T] {
	 //TODO: Complete me!
}

// End returns an iterator of the array with the position a.Size()
func (a *Array[T]) End() *ArrayIterator[T] {
	 //TODO: Complete me!
}

// First returns an iterator of the array with the first position
func (a *Array[T]) First() *ArrayIterator[T] {
	 //TODO: Complete me!
}

// Last returns an iterator of the array with the last position
func (a *Array[T]) Last() *ArrayIterator[T] {
	 //TODO: Complete me!
}

// IterAt returns an iterator of the array with position pos
func (a *Array[T]) IterAt(pos int) *ArrayIterator[T] {
	 //TODO: Complete me!
}

// String returns a string representation of the array
func (a *Array[T]) String() string {
	 //TODO: Complete me!
}
