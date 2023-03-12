package sort

import (
	"stl/utils/comparator"
	"stl/utils/iterator"
)

//BinarySearch returns true if exist an element witch value is val in the range [first, last), or false if not exist
func BinarySearch[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) bool {
	 //TODO: Complete me!
}

func binarySearch[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) bool {
	 //TODO: Complete me!
}

//LowerBound returns the iterator pointing to the first element greater than or equal to value passed in the range [first, last), or iterator last if not exist.
func LowerBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	 //TODO: Complete me!
}

func lowerBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	 //TODO: Complete me!
}

//UpperBound returns the iterator pointing to the first element greater than val in the range [first, last), or iterator last if not exist.
func UpperBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	 //TODO: Complete me!
}

func upperBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	 //TODO: Complete me!
}
