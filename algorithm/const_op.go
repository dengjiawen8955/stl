package algorithm

import (
	"stl/utils/comparator"
	"stl/utils/iterator"
)

// Count returns the number of elements that their value is equal to value in range [first, last)
func Count[T any](first, last iterator.ConstIterator[T], value T, cmp comparator.Comparator[T]) int {
	 //TODO: Complete me!
}

// CountIf returns the number of elements are satisfied the function f in range [first, last)
func CountIf[T any](first, last iterator.ConstIterator[T], f func(iterator.ConstIterator[T]) bool) int {
	 //TODO: Complete me!
}

// Find finds the first element that its value is equal to value in range [first, last), and returns its iterator, or last if not found
func Find[T any](first, last iterator.ConstIterator[T], value T, cmp comparator.Comparator[T]) iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// FindIf finds the first element that is satisfied the function f, and returns its iterator, or last if not found
func FindIf[T any](first, last iterator.ConstIterator[T], f func(iterator.ConstIterator[T]) bool) iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// MaxElement returns an Iterator to the largest element in the range [first, last). If several elements in the range are equivalent to the largest element, returns the iterator to the first such element. Returns last if the range is empty.
func MaxElement[T any](first, last iterator.ConstIterator[T], cmp comparator.Comparator[T]) iterator.ConstIterator[T] {
	 //TODO: Complete me!
}

// MinElement returns an Iterator to the smallest element value in the range [first, last). If several elements in the range are equivalent to the smallest element value, returns the iterator to the first such element. Returns last if the range is empty.
func MinElement[T any](first, last iterator.ConstIterator[T], cmp comparator.Comparator[T]) iterator.ConstIterator[T] {
	 //TODO: Complete me!
}
