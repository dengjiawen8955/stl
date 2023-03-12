package sort

import (
	"stl/utils/comparator"
	"stl/utils/iterator"
)

//Sort sorts the container by using quick sort
func Sort[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	 //TODO: Complete me!
}

func quickSort[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	 //TODO: Complete me!
}

func doPivot[T any](first, mid, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	 //TODO: Complete me!
}

func swapValue[T any](a, b iterator.RandomAccessIterator[T]) {
	 //TODO: Complete me!
}
