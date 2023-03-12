package priorityqueue

import (
	gosync "sync"

	"stl/ds/heap"
	"stl/utils/comparator"
	"stl/utils/sync"
)

var (
	defaultLocker sync.FakeLocker
)

// ElementHolder holds elements of the PriorityQueue
type ElementHolder[T any] struct {
	 //TODO: Complete me!
}

// Push pushes an element to the ElementHolder
func (h *ElementHolder[T]) Push(element T) {
	 //TODO: Complete me!
}

// Pop pops an element from the ElementHolder
func (h *ElementHolder[T]) Pop() T {
	 //TODO: Complete me!
}

func (h *ElementHolder[T]) top() T {
	 //TODO: Complete me!
}

// Len returns the amount of elements in ElementHolder
func (h *ElementHolder[T]) Len() int {
	 //TODO: Complete me!
}

// Len compare two elements at position i and j , and returns true if elements[i] < elements[j]
func (h *ElementHolder[T]) Less(i, j int) bool {
	 //TODO: Complete me!
}

// Swap swaps two elements at position i and j
func (h *ElementHolder[T]) Swap(i, j int) {
	 //TODO: Complete me!
}

// Options holds PriorityQueue's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set the PriorityQueue goroutine-safe
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// PriorityQueue is an implementation of priority queue
type PriorityQueue[T any] struct {
	 //TODO: Complete me!
}

// New creates a PriorityQueue
func New[T any](cmp comparator.Comparator[T], opts ...Option) *PriorityQueue[T] {
	 //TODO: Complete me!
}

// Push pushes an element to the PriorityQueue
func (q *PriorityQueue[T]) Push(e T) {
	 //TODO: Complete me!
}

// Pop pops an element from the PriorityQueue
func (q *PriorityQueue[T]) Pop() T {
	 //TODO: Complete me!
}

// Top returns the top element in the PriorityQueue
func (q *PriorityQueue[T]) Top() T {
	 //TODO: Complete me!
}

// Empty returns true if the PriorityQueue is empty, otherwise returns false
func (q *PriorityQueue[T]) Empty() bool {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the queue
func (q *PriorityQueue[T]) Size() int {
	 //TODO: Complete me!
}
