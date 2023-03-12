package main

import (
	"fmt"
	"stl/algorithm"
	"stl/ds/deque"
)

func main() {
	a := deque.New()
	for i := 0; i < 9; i++ {
		a.PushBack(i)
	}
	fmt.Printf("%v\n", a)

	algorithm.Swap(a.First(), a.Last())
	fmt.Printf("%v\n", a)

	algorithm.Reverse(a.Begin(), a.End())
	fmt.Printf("%v\n", a)
}
