package deque

//Segment is a fixed capacity ring
type Segment[T any] struct {
	 //TODO: Complete me!
}

func newSegment[T any](capacity int) *Segment[T] {
	 //TODO: Complete me!
}

func (s *Segment[T]) pushBack(value T) {
	 //TODO: Complete me!
}

func (s *Segment[T]) pushFront(val T) {
	 //TODO: Complete me!
}

func (s *Segment[T]) insert(position int, value T) {
	 //TODO: Complete me!
}

func (s *Segment[T]) popBack() T {
	 //TODO: Complete me!
}

func (s *Segment[T]) popFront() T {
	 //TODO: Complete me!
}

func (s *Segment[T]) eraseAt(position int) {
	 //TODO: Complete me!
}

func (s *Segment[T]) size() int {
	 //TODO: Complete me!
}

func (s *Segment[T]) capacity() int {
	 //TODO: Complete me!
}

func (s *Segment[T]) full() bool {
	 //TODO: Complete me!
}

func (s *Segment[T]) empty() bool {
	 //TODO: Complete me!
}

func (s *Segment[T]) nextIndex(index int) int {
	 //TODO: Complete me!
}

func (s *Segment[T]) preIndex(index int) int {
	 //TODO: Complete me!
}

func (s *Segment[T]) at(position int) T {
	 //TODO: Complete me!
}

func (s *Segment[T]) set(position int, val T) {
	 //TODO: Complete me!
}

func (s *Segment[T]) back() T {
	 //TODO: Complete me!
}

func (s *Segment[T]) front() T {
	 //TODO: Complete me!
}

func (s *Segment[T]) clear() {
	 //TODO: Complete me!
}
