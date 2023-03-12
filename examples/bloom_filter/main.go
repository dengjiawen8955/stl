package main

import (
	"fmt"
	"stl/ds/bloomfilter"
)

func main() {
	filter := bloomfilter.New(100, 4, bloomfilter.WithGoroutineSafe())
	filter.Add("hhhh")
	filter.Add("gggg")

	fmt.Printf("%v\n", filter.Contains("aaaa"))
	fmt.Printf("%v\n", filter.Contains("gggg"))
}
