package priorityqueue

import (
	"container/heap"
	"stl/utils/comparator"
	"stl/utils/sync"
	gosync "sync"
)

var (
	defaultComparator = comparator.BuiltinTypeComparator
	defaultLocker     sync.FakeLocker
)

// ElementHolder holds elements of the PriorityQueue
type ElementHolder struct {
	 //TODO: Complete me!
}

// Push pushes an element to the ElementHolder
func (h *ElementHolder) Push(element interface{}) {
	 //TODO: Complete me!
}

// Pop pops an element from the ElementHolder
func (h *ElementHolder) Pop() interface{} {
	 //TODO: Complete me!
}

func (h *ElementHolder) top() interface{} {
	 //TODO: Complete me!
}

// Len returns the amount of elements in ElementHolder
func (h *ElementHolder) Len() int {
	 //TODO: Complete me!
}

// Len compare two elements at position i and j , and returns true if elements[i] < elements[j]
func (h *ElementHolder) Less(i, j int) bool {
	 //TODO: Complete me!
}

// Swap swaps two elements at position i and j
func (h *ElementHolder) Swap(i, j int) {
	 //TODO: Complete me!
}

// Options holds PriorityQueue's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithComparator is used to set the PriorityQueue's comparator
func WithComparator(cmp comparator.Comparator) Option {
	 //TODO: Complete me!
}

// WithGoroutineSafe is used to set the PriorityQueue goroutine-safe
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// PriorityQueue is an implementation of priority queue
type PriorityQueue struct {
	 //TODO: Complete me!
}

// New creates a PriorityQueue
func New(opts ...Option) *PriorityQueue {
	 //TODO: Complete me!
}

// Push pushes an element to the PriorityQueue
func (q *PriorityQueue) Push(e interface{}) {
	 //TODO: Complete me!
}

// Pop pops an element from the PriorityQueue
func (q *PriorityQueue) Pop() interface{} {
	 //TODO: Complete me!
}

// Top returns the top element in the PriorityQueue
func (q *PriorityQueue) Top() interface{} {
	 //TODO: Complete me!
}

// Empty returns true if the PriorityQueue is empty, otherwise returns false
func (q *PriorityQueue) Empty() bool {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the queue
func (q *PriorityQueue) Size() int {
	 //TODO: Complete me!
}
