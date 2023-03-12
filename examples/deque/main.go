package main

import (
	"fmt"
	"math/rand"
	"stl/algorithm/sort"
	"stl/ds/deque"
	"stl/utils/comparator"
)

func main() {
	q := deque.New[int]()
	for i := 0; i < 100; i++ {
		r := rand.Int() % 100
		q.PushBack(r)
		q.PushFront(r)
	}
	fmt.Printf("%v\n", q)

	sort.Sort[int](q.Begin(), q.End(), comparator.IntComparator)
	fmt.Printf("%v\n", q)

	for !q.Empty() {
		r := rand.Int() % q.Size()
		q.EraseAt(r)
	}
	fmt.Printf("%v\n", q)
}
