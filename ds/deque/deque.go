package deque

import (
	"errors"
	"fmt"
)

// Constants definition
const (
	SegmentCapacity = 128
)

// Define internal errors
var (
	ErrOutOfRange = errors.New("out off range")
)

// Deque is double-ended queue supports efficient data insertion from the head and tail, random access and iterator access.
type Deque[T any] struct {
	 //TODO: Complete me!
}

// New creates a new deque
func New[T any]() *Deque[T] {
	 //TODO: Complete me!
}

// Size returns the amount of values in the deque
func (d *Deque[T]) Size() int {
	 //TODO: Complete me!
}

// Empty returns true if the deque is empty,otherwise returns false.
func (d *Deque[T]) Empty() bool {
	 //TODO: Complete me!
}

func (d *Deque[T]) segUsed() int {
	 //TODO: Complete me!
}

// PushFront pushed a value to the front of the deque
func (d *Deque[T]) PushFront(value T) {
	 //TODO: Complete me!
}

// PushBack pushed a value to the back of deque
func (d *Deque[T]) PushBack(value T) {
	 //TODO: Complete me!
}

// Insert inserts a value to the position pos of the deque
func (d *Deque[T]) Insert(pos int, value T) {
	 //TODO: Complete me!
}

func (d *Deque[T]) moveFrontInsert(seg, pos int, value T) {
	 //TODO: Complete me!
}

func (d *Deque[T]) moveBackInsert(seg, pos int, value T) {
	 //TODO: Complete me!
}

// Front returns the value at the first position of the deque
func (d *Deque[T]) Front() T {
	 //TODO: Complete me!
}

// Back returns the value at the last position of the deque
func (d *Deque[T]) Back() T {
	 //TODO: Complete me!
}

// At returns the value at position pos of the deque
func (d *Deque[T]) At(pos int) T {
	 //TODO: Complete me!
}

// Set sets the value of the deque's position pos with value val
func (d *Deque[T]) Set(pos int, val T) error {
	 //TODO: Complete me!
}

// PopFront returns the value at the first position of the deque and removes it
func (d *Deque[T]) PopFront() T {
	 //TODO: Complete me!
}

// PopBack returns the value at the lase position of the deque and removes it
func (d *Deque[T]) PopBack() T {
	 //TODO: Complete me!
}

// EraseAt erases the element at the position pos
func (d *Deque[T]) EraseAt(pos int) {
	 //TODO: Complete me!
}

// EraseRange erases elements in range [firstPos, lastPos)
func (d *Deque[T]) EraseRange(firstPos, lastPos int) {
	 //TODO: Complete me!
}

// Clear erases all elements in the deque
func (d *Deque[T]) Clear() {
	 //TODO: Complete me!
}

func (d *Deque[T]) putToPool(s *Segment[T]) {
	 //TODO: Complete me!
}

func (d *Deque[T]) firstAvailableSegment() *Segment[T] {
	 //TODO: Complete me!
}

func (d *Deque[T]) lastAvailableSegment() *Segment[T] {
	 //TODO: Complete me!
}

func (d *Deque[T]) firstSegment() *Segment[T] {
	 //TODO: Complete me!
}

func (d *Deque[T]) lastSegment() *Segment[T] {
	 //TODO: Complete me!
}

func (d *Deque[T]) segmentAt(seg int) *Segment[T] {
	 //TODO: Complete me!
}

func (d *Deque[T]) pos(position int) (int, int) {
	 //TODO: Complete me!
}

func (d *Deque[T]) expand() {
	 //TODO: Complete me!
}

// shrinkIfNeeded shrinks the deque if it has too many unused space.
func (d *Deque[T]) shrinkIfNeeded() {
	 //TODO: Complete me!
}

func (d *Deque[T]) nextIndex(index int) int {
	 //TODO: Complete me!
}

func (d *Deque[T]) preIndex(index int) int {
	 //TODO: Complete me!
}

// String returns a string representation of the deque
func (d *Deque[T]) String() string {
	 //TODO: Complete me!
}

// Begin returns an iterator of the deque with the first position
func (d *Deque[T]) Begin() *DequeIterator[T] {
	 //TODO: Complete me!
}

// End returns an iterator of the deque with the position d.Size()
func (d *Deque[T]) End() *DequeIterator[T] {
	 //TODO: Complete me!
}

// First returns an iterator of the deque with the first position
func (d *Deque[T]) First() *DequeIterator[T] {
	 //TODO: Complete me!
}

// Last returns an iterator of the deque with the last position
func (d *Deque[T]) Last() *DequeIterator[T] {
	 //TODO: Complete me!
}

// IterAt returns an iterator of the deque with the position pos
func (d *Deque[T]) IterAt(pos int) *DequeIterator[T] {
	 //TODO: Complete me!
}
