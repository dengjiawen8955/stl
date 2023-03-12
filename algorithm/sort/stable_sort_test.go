package sort

import (
	"math/rand"
	"stl/ds/vector"
	"stl/utils/comparator"
	"testing"
	"time"
)

func TestStableSort(t *testing.T) {
	// test size = 0
	rand.Seed(time.Now().UnixNano())
	v := vector.New[int]()
	Sort[int](v.Begin(), v.End(), comparator.IntComparator)
	t.Logf("v: %v", v.String())

	// test size = 1
	v = vector.New[int]()
	for i := 0; i < 1; i++ {
		v.PushBack(rand.Int() % 10)
	}
	Stable[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}

	// test size = 2
	v = vector.New[int]()
	for i := 0; i < 2; i++ {
		v.PushBack(rand.Int() % 10)
	}
	Stable[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}
}

func TestStableSort10(t *testing.T) {
	// test size = 10
	v := vector.New[int]()
	for i := 0; i < 10; i++ {
		v.PushBack(rand.Int() % 10)
	}
	Stable[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}
}

func TestStableSort31(t *testing.T) {
	// test size = 31
	v := vector.New[int]()
	for i := 0; i < 31; i++ {
		v.PushBack(rand.Int() % 10)
	}
	Stable[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}
}

func TestStableSort50(t *testing.T) {
	// test size = 50
	v := vector.New[int]()
	for i := 0; i < 50; i++ {
		v.PushBack(rand.Int() % 100)
	}
	Stable[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}
}
