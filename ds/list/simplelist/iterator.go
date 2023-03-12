package simplelist

import "stl/utils/iterator"

//ListIterator is an implementation of Iterator
var _ iterator.Iterator = (*ListIterator)(nil)

// ListIterator is an iterator for list
type ListIterator struct {
	 //TODO: Complete me!
}

// NewIterator news a ListIterator
func NewIterator(node *Node) *ListIterator {
	 //TODO: Complete me!
}

// IsValid returns whether iter is valid
func (iter *ListIterator) IsValid() bool {
	 //TODO: Complete me!
}

// Next returns the next iterator
func (iter *ListIterator) Next() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Value returns the internal value of iter
func (iter *ListIterator) Value() interface{} {
	 //TODO: Complete me!
}

// SetValue sets the value of iter
func (iter *ListIterator) SetValue(value interface{}) {
	 //TODO: Complete me!
}

// Clone clones iter to a new ListIterator
func (iter *ListIterator) Clone() iterator.ConstIterator {
	 //TODO: Complete me!
}

// Equal returns whether iter is equal to other
func (iter *ListIterator) Equal(other iterator.ConstIterator) bool {
	 //TODO: Complete me!
}
