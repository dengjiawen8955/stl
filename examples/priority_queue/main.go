package main

import (
	"fmt"
	"stl/ds/priorityqueue"
	"stl/utils/comparator"
)

func main() {
	q := priorityqueue.New[int](comparator.Reverse(comparator.IntComparator),
		priorityqueue.WithGoroutineSafe())
	q.Push(4)
	q.Push(13)
	q.Push(7)
	q.Push(9)
	q.Push(0)
	q.Push(88)

	for !q.Empty() {
		fmt.Printf("%v\n", q.Pop())
	}
}
