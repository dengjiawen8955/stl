package deque

// Pool is a memory pool for holding Segments
type Pool struct {
	 //TODO: Complete me!
}

func newPool() *Pool {
	 //TODO: Complete me!
}

func (p *Pool) get() *Segment {
	 //TODO: Complete me!
}

func (p *Pool) put(s *Segment) {
	 //TODO: Complete me!
}

func (p *Pool) shrinkToSize(size int) {
	 //TODO: Complete me!
}

func (p *Pool) size() int {
	 //TODO: Complete me!
}
