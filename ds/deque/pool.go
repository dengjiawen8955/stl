package deque

// Pool is a memory pool for holding Segments
type Pool[T any] struct {
	 //TODO: Complete me!
}

func newPool[T any]() *Pool[T] {
	 //TODO: Complete me!
}

func (p *Pool[T]) get() *Segment[T] {
	 //TODO: Complete me!
}

func (p *Pool[T]) put(s *Segment[T]) {
	 //TODO: Complete me!
}

func (p *Pool[T]) shrinkToSize(size int) {
	 //TODO: Complete me!
}

func (p *Pool[T]) size() int {
	 //TODO: Complete me!
}
