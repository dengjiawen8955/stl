package algorithm

import (
	"stl/utils/comparator"
	"stl/utils/iterator"
)

// Count returns the number of elements that their value is equal to value in range [first, last)
func Count(first, last iterator.ConstIterator, value interface{}, cmps ...comparator.Comparator) int {
	 //TODO: Complete me!
}

// CountIf returns the number of elements are satisfied the function f in range [first, last)
func CountIf(first, last iterator.ConstIterator, f func(iterator.ConstIterator) bool) int {
	 //TODO: Complete me!
}

// Find finds the first element that its value is equal to value in range [first, last), and returns its iterator, or last if not found
func Find(first, last iterator.ConstIterator, value interface{}, cmps ...comparator.Comparator) iterator.ConstIterator {
	 //TODO: Complete me!
}

// FindIf finds the first element that is satisfied the function f, and returns its iterator, or last if not found
func FindIf(first, last iterator.ConstIterator, f func(iterator.ConstIterator) bool) iterator.ConstIterator {
	 //TODO: Complete me!
}
