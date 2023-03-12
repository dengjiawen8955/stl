package slice

import "reflect"

//SliceWrapper wraps a slice in order to provide functions related to iterators
type SliceWrapper struct {
	 //TODO: Complete me!
}

// NewSliceWrapper creates a SliceWrapper
func NewSliceWrapper(slice interface{}) *SliceWrapper {
	 //TODO: Complete me!
}

// Attach update the internal slice to newSlice
func (s *SliceWrapper) Attach(newSlice interface{}) {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s *SliceWrapper) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s *SliceWrapper) At(position int) interface{} {
	 //TODO: Complete me!
}

// Set sets value at position
func (s *SliceWrapper) Set(position int, val interface{}) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s *SliceWrapper) Begin() *SliceIterator {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s *SliceWrapper) End() *SliceIterator {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s *SliceWrapper) First() *SliceIterator {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s *SliceWrapper) Last() *SliceIterator {
	 //TODO: Complete me!
}
