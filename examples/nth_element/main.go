package main

import (
	"fmt"
	"stl/algorithm/sort"
	"stl/ds/deque"
	"stl/utils/comparator"
)

func main() {
	a := deque.New[int]()
	a.PushBack(9)
	a.PushBack(8)
	a.PushBack(7)
	a.PushBack(6)
	a.PushBack(5)
	a.PushBack(4)
	a.PushBack(3)
	a.PushBack(2)
	a.PushBack(1)
	fmt.Printf("%v\n", a)
	sort.NthElement[int](a.Begin(), a.End(), 3, comparator.IntComparator)
	fmt.Printf("%v\n", a.At(3))
	fmt.Printf("%v\n", a)
}
