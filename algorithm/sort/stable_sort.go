package sort

import (
	"stl/utils/comparator"
	"stl/utils/iterator"
)

//Stable sorts the container by using merge sort
func Stable[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	 //TODO: Complete me!
}

func mergeSort[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T], tempSlice []T) {
	 //TODO: Complete me!
}

func merge[T any](first, mid, end iterator.RandomAccessIterator[T], cmp comparator.Comparator[T], tempSlice []T) {
	 //TODO: Complete me!
}
