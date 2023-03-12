package slice

//SliceWrapper wraps a slice in order to provide functions related to iterators
type SliceWrapper[T any] struct {
	 //TODO: Complete me!
}

// NewSliceWrapper creates a SliceWrapper
func NewSliceWrapper[T any](slice []T) *SliceWrapper[T] {
	 //TODO: Complete me!
}

// Attach update the internal slice to newSlice
func (s *SliceWrapper[T]) Attach(newSlice []T) {
	 //TODO: Complete me!
}

// Len returns the length of s
func (s *SliceWrapper[T]) Len() int {
	 //TODO: Complete me!
}

// At returns the value at position
func (s *SliceWrapper[T]) At(position int) T {
	 //TODO: Complete me!
}

// Set sets value at position
func (s *SliceWrapper[T]) Set(position int, val T) {
	 //TODO: Complete me!
}

// Begin returns the first iterator of s
func (s *SliceWrapper[T]) Begin() *SliceIterator[T] {
	 //TODO: Complete me!
}

// End returns the end iterator of s
func (s *SliceWrapper[T]) End() *SliceIterator[T] {
	 //TODO: Complete me!
}

// First returns the first iterator of s
func (s *SliceWrapper[T]) First() *SliceIterator[T] {
	 //TODO: Complete me!
}

// Last returns the last iterator of s
func (s *SliceWrapper[T]) Last() *SliceIterator[T] {
	 //TODO: Complete me!
}
