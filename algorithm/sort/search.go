package sort

import (
	"stl/utils/comparator"
	"stl/utils/iterator"
)

//BinarySearch returns true if exist an element witch value is val in the range [first, last), or false if not exist
func BinarySearch(first, last iterator.RandomAccessIterator, val interface{}, cmp ...comparator.Comparator) bool {
	 //TODO: Complete me!
}

func binarySearch(first, last iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) bool {
	 //TODO: Complete me!
}

//LowerBound returns the iterator pointing to the first element greater than or equal to value passed in the range [first, last), or iterator last if not exist.
func LowerBound(first, last iterator.RandomAccessIterator, val interface{}, cmp ...comparator.Comparator) iterator.RandomAccessIterator {
	 //TODO: Complete me!
}

func lowerBound(first, last iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) iterator.RandomAccessIterator {
	 //TODO: Complete me!
}

//UpperBound returns the iterator pointing to the first element greater than val in the range [first, last), or iterator last if not exist.
func UpperBound(first, last iterator.RandomAccessIterator, val interface{}, cmp ...comparator.Comparator) iterator.RandomAccessIterator {
	 //TODO: Complete me!
}

func upperBound(first, last iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) iterator.RandomAccessIterator {
	 //TODO: Complete me!
}
