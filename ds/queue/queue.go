package queue

import (
	"stl/ds/container"
	"stl/ds/deque"
	"stl/ds/list/bidlist"
	"stl/utils/sync"
	gosync "sync"
)

var (
	defaultLocker sync.FakeLocker
)

// Options holds Queue's options
type Options[T any] struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option[T any] func(option *Options[T])

// WithGoroutineSafe is used to set a Queue goroutine-safe
func WithGoroutineSafe[T any]() Option[T] {
	 //TODO: Complete me!
}

// WithContainer is used to set a Queue's underlying container
func WithContainer[T any](c container.Container[T]) Option[T] {
	 //TODO: Complete me!
}

// WithListContainer is used to set List as a Queue's underlying container
func WithListContainer[T any]() Option[T] {
	 //TODO: Complete me!
}

// Queue is a first-in-first-out data structure
type Queue[T any] struct {
	 //TODO: Complete me!
}

//New creates a new queue
func New[T any](opts ...Option[T]) *Queue[T] {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the queue
func (q *Queue[T]) Size() int {
	 //TODO: Complete me!
}

// Empty returns true if the queue is empty, otherwise returns false
func (q *Queue[T]) Empty() bool {
	 //TODO: Complete me!
}

// Push pushes a value to the end of the queue
func (q *Queue[T]) Push(value T) {
	 //TODO: Complete me!
}

// Front returns the front value in the queue
func (q *Queue[T]) Front() any {
	 //TODO: Complete me!
}

// Back returns the back value in the queue
func (q *Queue[T]) Back() any {
	 //TODO: Complete me!
}

// Pop removes the the front element in the queue, and returns its value
func (q *Queue[T]) Pop() any {
	 //TODO: Complete me!
}

// Clear clears all elements in the queue
func (q *Queue[T]) Clear() {
	 //TODO: Complete me!
}

// String returns a string representation of the queue
func (q *Queue[T]) String() string {
	 //TODO: Complete me!
}
