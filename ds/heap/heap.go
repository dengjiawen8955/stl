package heap

type Interface[T any] interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)

	// Push add x as element Len()
	Push(x T)

	// Pop remove and return element Len() - 1.
	Pop() T
}

func Init[T any](h Interface[T]) {
	 //TODO: Complete me!
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push[T any](h Interface[T], x T) {
	 //TODO: Complete me!
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop[T any](h Interface[T]) T {
	 //TODO: Complete me!
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove[T any](h Interface[T], i int) T {
	 //TODO: Complete me!
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix[T any](h Interface[T], i int) {
	 //TODO: Complete me!
}

func up[T any](h Interface[T], j int) {
	 //TODO: Complete me!
}

func down[T any](h Interface[T], i0, n int) bool {
	 //TODO: Complete me!
}
