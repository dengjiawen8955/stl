package bidlist

import "stl/utils/iterator"

//ListIterator is an implementation of BidIterator
var _ iterator.BidIterator = (*ListIterator)(nil)

// ListIterator is an implementation of list iterator
type ListIterator struct {
	 //TODO: Complete me!
}

// NewIterator creates a ListIterator
func NewIterator(node *Node) *ListIterator {
	 //TODO: Complete me!
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *ListIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Next moves the pointer of iterator to the next node and returns itself
func (iter *ListIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Prev moves the pointer of iterator to the previous node and returns itself
func (iter *ListIterator) Prev() iterator.ConstBidIterator {
	 //TODO: Complete me!
}

// Value returns the node's value of the iterator point to
func (iter *ListIterator) Value() interface{} {
	 //TODO: Complete me!
}

// SetValue sets the node's value of the iterator point to
func (iter *ListIterator) SetValue(value interface{}) {
	 //TODO: Complete me!
}

// Clone clones the iterator to a new iterator
func (iter *ListIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ListIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
