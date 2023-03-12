package stack

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

// Options holds the Stack's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set a stack goroutine-safe
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// WithContainer is used to set a stack's underlying container
func WithContainer(c container.Container) Option {
	 //TODO: Complete me!
}

// WithListContainer is used to set List for a stack's underlying container
func WithListContainer() Option {
	 //TODO: Complete me!
}

//Stack is a last-in-first-out data structure
type Stack struct {
	 //TODO: Complete me!
}

// New creates a new stack
func New(opts ...Option) *Stack {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the stack
func (s *Stack) Size() int {
	 //TODO: Complete me!
}

// Empty returns true if the stack is empty, otherwise returns false
func (s *Stack) Empty() bool {
	 //TODO: Complete me!
}

// Push pushes a value to the stack
func (s *Stack) Push(value interface{}) {
	 //TODO: Complete me!
}

// Top returns the top value in the stack
func (s *Stack) Top() interface{} {
	 //TODO: Complete me!
}

// Pop removes the the top value in the stack and returns it
func (s *Stack) Pop() interface{} {
	 //TODO: Complete me!
}

// Clear clears all elements in the stack
func (s *Stack) Clear() {
	 //TODO: Complete me!
}

// String returns a string representation of the stack
func (s *Stack) String() string {
	 //TODO: Complete me!
}
