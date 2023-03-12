package queue

import (
	"stl/ds/container"
	"stl/ds/deque"
	"stl/ds/list/bidlist"
	"stl/utils/sync"
	gosync "sync"
)

var (
	defaultLocker    sync.FakeLocker
	defaultContainer = deque.New()
)

// Options holds Queue's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set a Queue goroutine-safe
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// WithContainer is used to set a Queue's underlying container
func WithContainer(c container.Container) Option {
	 //TODO: Complete me!
}

// WithListContainer is used to set List as a Queue's underlying container
func WithListContainer() Option {
	 //TODO: Complete me!
}

// Queue is a first-in-first-out data structure
type Queue struct {
	 //TODO: Complete me!
}

//New creates a new queue
func New(opts ...Option) *Queue {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the queue
func (q *Queue) Size() int {
	 //TODO: Complete me!
}

// Empty returns true if the queue is empty, otherwise returns false
func (q *Queue) Empty() bool {
	 //TODO: Complete me!
}

// Push pushes a value to the end of the queue
func (q *Queue) Push(value interface{}) {
	 //TODO: Complete me!
}

// Front returns the front value in the queue
func (q *Queue) Front() interface{} {
	 //TODO: Complete me!
}

// Back returns the back value in the queue
func (q *Queue) Back() interface{} {
	 //TODO: Complete me!
}

// Pop removes the the front element in the queue, and returns its value
func (q *Queue) Pop() interface{} {
	 //TODO: Complete me!
}

// Clear clears all elements in the queue
func (q *Queue) Clear() {
	 //TODO: Complete me!
}

// String returns a string representation of the queue
func (q *Queue) String() string {
	 //TODO: Complete me!
}
