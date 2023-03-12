package main

import (
	"fmt"
	"stl/algorithm"
	"stl/ds/deque"
)

func main() {
	a := deque.New[int]()
	for i := 0; i < 9; i++ {
		a.PushBack(i)
	}
	fmt.Printf("%v\n", a)

	algorithm.Swap[int](a.First(), a.Last())
	fmt.Printf("%v\n", a)

	algorithm.Reverse[int](a.Begin(), a.End())
	fmt.Printf("%v\n", a)
}
