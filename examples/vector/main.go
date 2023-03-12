package main

import (
	"fmt"
	"stl/algorithm/sort"
	"stl/ds/vector"
	"stl/utils/comparator"
)

func main() {
	v := vector.New[int]()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	for i := 0; i < v.Size(); i++ {
		fmt.Printf("%v ", v.At(i))
	}
	fmt.Printf("\n")

	// sort in descending
	sort.Sort[int](v.Begin(), v.End(), comparator.Reverse(comparator.IntComparator))
	for iter := v.Begin(); iter.IsValid(); iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
}
