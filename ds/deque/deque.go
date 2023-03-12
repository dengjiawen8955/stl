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
	ErrOutOffRange = errors.New("out off range")
)

// Deque is double-ended queue supports efficient data insertion from the head and tail, random access and iterator access.
type Deque struct {
	 //TODO: Complete me!
}

// New creates a new deque
func New() *Deque {
	 //TODO: Complete me!
}

// Size returns the amount of values in the deque
func (d *Deque) Size() int {
	 //TODO: Complete me!
}

// Empty returns true if the deque is empty,otherwise returns false.
func (d *Deque) Empty() bool {
	 //TODO: Complete me!
}

func (d *Deque) segUsed() int {
	 //TODO: Complete me!
}

// PushFront pushed a value to the front of the deque
func (d *Deque) PushFront(value interface{}) {
	 //TODO: Complete me!
}

// PushBack pushed a value to the back of deque
func (d *Deque) PushBack(value interface{}) {
	 //TODO: Complete me!
}

// Insert inserts a value to the position pos of the deque
func (d *Deque) Insert(pos int, value interface{}) {
	 //TODO: Complete me!
}

func (d *Deque) moveFrontInsert(seg, pos int, value interface{}) {
	 //TODO: Complete me!
}

func (d *Deque) moveBackInsert(seg, pos int, value interface{}) {
	 //TODO: Complete me!
}

// Front returns the value at the first position of the deque
func (d *Deque) Front() interface{} {
	 //TODO: Complete me!
}

// Back returns the value at the last position of the deque
func (d *Deque) Back() interface{} {
	 //TODO: Complete me!
}

// At returns the value at position pos of the deque
func (d *Deque) At(pos int) interface{} {
	 //TODO: Complete me!
}

// Set sets the value of the deque's position pos with value val
func (d *Deque) Set(pos int, val interface{}) error {
	 //TODO: Complete me!
}

// PopFront returns the value at the first position of the deque and removes it
func (d *Deque) PopFront() interface{} {
	 //TODO: Complete me!
}

// PopBack returns the value at the lase position of the deque and removes it
func (d *Deque) PopBack() interface{} {
	 //TODO: Complete me!
}

// EraseAt erases the element at the position pos
func (d *Deque) EraseAt(pos int) {
	 //TODO: Complete me!
}

// EraseRange erases elements in range [firstPos, lastPos)
func (d *Deque) EraseRange(firstPos, lastPos int) {
	 //TODO: Complete me!
}

// Clear erases all elements in the deque
func (d *Deque) Clear() {
	 //TODO: Complete me!
}

func (d *Deque) putToPool(s *Segment) {
	 //TODO: Complete me!
}

func (d *Deque) firstAvailableSegment() *Segment {
	 //TODO: Complete me!
}

func (d *Deque) lastAvailableSegment() *Segment {
	 //TODO: Complete me!
}

func (d *Deque) firstSegment() *Segment {
	 //TODO: Complete me!
}

func (d *Deque) lastSegment() *Segment {
	 //TODO: Complete me!
}

func (d *Deque) segmentAt(seg int) *Segment {
	 //TODO: Complete me!
}

func (d *Deque) pos(position int) (int, int) {
	 //TODO: Complete me!
}

func (d *Deque) expand() {
	 //TODO: Complete me!
}

//shrinkIfNeeded shrinks the deque if it has too many unused space.
func (d *Deque) shrinkIfNeeded() {
	 //TODO: Complete me!
}

func (d *Deque) nextIndex(index int) int {
	 //TODO: Complete me!
}

func (d *Deque) preIndex(index int) int {
	 //TODO: Complete me!
}

// String returns a string representation of the deque
func (d *Deque) String() string {
	 //TODO: Complete me!
}

// Begin returns an iterator of the deque with the first position
func (d *Deque) Begin() *DequeIterator {
	 //TODO: Complete me!
}

// End returns an iterator of the deque with the position d.Size()
func (d *Deque) End() *DequeIterator {
	 //TODO: Complete me!
}

// First returns an iterator of the deque with the first position
func (d *Deque) First() *DequeIterator {
	 //TODO: Complete me!
}

// Last returns an iterator of the deque with the last position
func (d *Deque) Last() *DequeIterator {
	 //TODO: Complete me!
}

// IterAt returns an iterator of the deque with the position pos
func (d *Deque) IterAt(pos int) *DequeIterator {
	 //TODO: Complete me!
}
