package main

import (
	"fmt"
	"stl/ds/priorityqueue"
	"stl/utils/comparator"
)

func main() {
	q := priorityqueue.New(priorityqueue.WithComparator(comparator.Reverse(comparator.BuiltinTypeComparator)),
		priorityqueue.WithGoroutineSafe())
	q.Push(5)
	q.Push(13)
	q.Push(7)
	q.Push(9)
	q.Push(0)
	q.Push(88)
	for !q.Empty() {
		fmt.Printf("%v\n", q.Pop())
	}
}
