package skiplist

import (
	"math/rand"
	"stl/utils/comparator"
	"stl/utils/sync"
	"stl/utils/visitor"
	gosync "sync"
	"time"
)

var (
	defaultKeyComparator = comparator.BuiltinTypeComparator
	defaultMaxLevel      = 10
	defaultLocker        sync.FakeLocker
)

// Options holds Skiplist's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function used to set Options
type Option func(option *Options)

// WithKeyComparator sets Key comparator option
func WithKeyComparator(cmp comparator.Comparator) Option {
	 //TODO: Complete me!
}

// WithGoroutineSafe sets Skiplist goroutine-safety,
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// WithMaxLevel sets max level of Skiplist
func WithMaxLevel(maxLevel int) Option {
	 //TODO: Complete me!
}

// Node is a list node
type Node struct {
	 //TODO: Complete me!
}

// Element is a kind of node with key-value data
type Element struct {
	 //TODO: Complete me!
}

// Skiplist is a kind of data structure which can search quickly by exchanging space for time
type Skiplist struct {
	 //TODO: Complete me!
}

// New news a Skiplist
func New(opts ...Option) *Skiplist {
	 //TODO: Complete me!
}

// Insert inserts a key-value pair into the skiplist
func (sl *Skiplist) Insert(key, value interface{}) {
	 //TODO: Complete me!
}

// Get returns the value associated with the passed key if the key is in the skiplist, otherwise returns nil
func (sl *Skiplist) Get(key interface{}) interface{} {
	 //TODO: Complete me!
}

// Remove removes the key-value pair associated with the passed key and returns true if the key is in the skiplist, otherwise returns false
func (sl *Skiplist) Remove(key interface{}) bool {
	 //TODO: Complete me!
}

// Len returns the amount of key-value pair in the skiplist
func (sl *Skiplist) Len() int {
	 //TODO: Complete me!
}

func (sl *Skiplist) randomLevel() int {
	 //TODO: Complete me!
}

func (sl *Skiplist) findPrevNodes(key interface{}) []*Node {
	 //TODO: Complete me!
}

// Traversal traversals elements in the skiplist, it will stop until to the end or the visitor returns false
func (sl *Skiplist) Traversal(visitor visitor.KvVisitor) {
	 //TODO: Complete me!
}

// Keys returns all keys in the skiplist
func (sl *Skiplist) Keys() []interface{} {
	 //TODO: Complete me!
}
