package main

import (
	"fmt"
	"stl/algorithm/sort"
	"stl/ds/slice"
	"stl/utils/comparator"
)

func main() {
	a := make([]int, 0)
	for i := 1; i <= 3; i++ {
		a = append(a, i)
	}
	wa := slice.NewSliceWrapper(a)
	fmt.Println("NextPermutation")
	for {
		fmt.Printf("%v\n", a)
		if !sort.NextPermutation[int](wa.Begin(), wa.End(), comparator.IntComparator) {
			break
		}
	}
	fmt.Println("PrePermutation")
	for {
		fmt.Printf("%v\n", a)
		if !sort.NextPermutation[int](wa.Begin(), wa.End(), comparator.Reverse(comparator.IntComparator)) {
			break
		}
	}
}
