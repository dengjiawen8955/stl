package sort

import (
	"stl/utils/comparator"
	"stl/utils/iterator"
)

//Stable sorts the container by using merge sort
func Stable(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	 //TODO: Complete me!
}

func mergeSort(first, last iterator.RandomAccessIterator, cmp comparator.Comparator, tempSlice []interface{}) {
	 //TODO: Complete me!
}

func merge(first, mid, end iterator.RandomAccessIterator, cmp comparator.Comparator, tempSlice []interface{}) {
	 //TODO: Complete me!
}
