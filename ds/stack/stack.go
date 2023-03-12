package stack

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

// Options holds the Stack's options
type Options[T any] struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option[T any] func(option *Options[T])

// WithGoroutineSafe is used to set a stack goroutine-safe
func WithGoroutineSafe[T any]() Option[T] {
	 //TODO: Complete me!
}

// WithContainer is used to set a stack's underlying container
func WithContainer[T any](c container.Container[T]) Option[T] {
	 //TODO: Complete me!
}

// WithListContainer is used to set List for a stack's underlying container
func WithListContainer[T any]() Option[T] {
	 //TODO: Complete me!
}

//Stack is a last-in-first-out data structure
type Stack[T any] struct {
	 //TODO: Complete me!
}

// New creates a new stack
func New[T any](opts ...Option[T]) *Stack[T] {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the stack
func (s *Stack[T]) Size() int {
	 //TODO: Complete me!
}

// Empty returns true if the stack is empty, otherwise returns false
func (s *Stack[T]) Empty() bool {
	 //TODO: Complete me!
}

// Push pushes a value to the stack
func (s *Stack[T]) Push(value T) {
	 //TODO: Complete me!
}

// Top returns the top value in the stack
func (s *Stack[T]) Top() T {
	 //TODO: Complete me!
}

// Pop removes the top value in the stack and returns it
func (s *Stack[T]) Pop() T {
	 //TODO: Complete me!
}

// Clear clears all elements in the stack
func (s *Stack[T]) Clear() {
	 //TODO: Complete me!
}

// String returns a string representation of the stack
func (s *Stack[T]) String() string {
	 //TODO: Complete me!
}
