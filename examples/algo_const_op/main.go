package main

import (
	"fmt"
	"stl/algorithm"
	"stl/ds/deque"
	"stl/utils/comparator"
	"stl/utils/iterator"
)

func isEven(iter iterator.ConstIterator[int]) bool {
	return iter.Value()%2 == 0
}

func greaterThan5(iter iterator.ConstIterator[int]) bool {
	return iter.Value() > 5
}

func main() {
	a := deque.New[int]()
	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	for i := 0; i < 5; i++ {
		a.PushBack(i)
	}
	fmt.Printf("%v\n", a)

	fmt.Printf("Count 2: %v\n", algorithm.Count[int](a.Begin(), a.End(), 2, comparator.IntComparator))

	fmt.Printf("Count 2: %v\n", algorithm.CountIf[int](a.Begin(), a.End(), isEven))

	iter := algorithm.Find[int](a.Begin(), a.End(), 2, comparator.IntComparator)
	if !iter.Equal(a.End()) {
		fmt.Printf("Fund %v\n", iter.Value())
	}
	iter = algorithm.FindIf[int](a.Begin(), a.End(), greaterThan5)
	if !iter.Equal(a.End()) {
		fmt.Printf("FindIf greaterThan5 : %v\n", iter.Value())
	}
	iter = algorithm.MaxElement[int](a.Begin(), a.End(), comparator.IntComparator)
	if !iter.Equal(a.End()) {
		fmt.Printf("largest value : %v\n", iter.Value())
	}
	iter = algorithm.MinElement[int](a.Begin(), a.End(), comparator.IntComparator)
	if !iter.Equal(a.End()) {
		fmt.Printf("largest value : %v\n", iter.Value())
	}
}
