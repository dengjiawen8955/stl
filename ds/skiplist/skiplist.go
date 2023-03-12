package skiplist

import (
	"errors"
	"math/rand"
	"stl/utils/comparator"
	"stl/utils/sync"
	"stl/utils/visitor"
	gosync "sync"
	"time"
)

var (
	defaultMaxLevel = 10
	defaultLocker   sync.FakeLocker
)
var ErrorNotFound = errors.New("not found")

// Options holds Skiplist's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function used to set Options
type Option func(option *Options)

// WithGoroutineSafe sets Skiplist goroutine-safety,
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// WithMaxLevel sets max level of Skiplist
func WithMaxLevel(maxLevel int) Option {
	 //TODO: Complete me!
}

// Node is a list node
type Node[K, V any] struct {
	 //TODO: Complete me!
}

// Element is a kind of node with key-value data
type Element[K, V any] struct {
	 //TODO: Complete me!
}

// Skiplist is a kind of data structure which can search quickly by exchanging space for time
type Skiplist[K, V any] struct {
	 //TODO: Complete me!
}

// New news a Skiplist
func New[K, V any](cmp comparator.Comparator[K], opts ...Option) *Skiplist[K, V] {
	 //TODO: Complete me!
}

// Insert inserts a key-value pair into the skiplist
func (sl *Skiplist[K, V]) Insert(key K, value V) {
	 //TODO: Complete me!
}

// Get returns the value associated with the passed key if the key is in the skiplist, otherwise returns error
func (sl *Skiplist[K, V]) Get(key K) (V, error) {
	 //TODO: Complete me!
}

// Remove removes the key-value pair associated with the passed key and returns true if the key is in the skiplist, otherwise returns false
func (sl *Skiplist[K, V]) Remove(key K) bool {
	 //TODO: Complete me!
}

// Len returns the amount of key-value pair in the skiplist
func (sl *Skiplist[K, V]) Len() int {
	 //TODO: Complete me!
}

func (sl *Skiplist[K, V]) randomLevel() int {
	 //TODO: Complete me!
}

func (sl *Skiplist[K, V]) findPrevNodes(key K) []*Node[K, V] {
	 //TODO: Complete me!
}

// Traversal traversals elements in the skiplist, it will stop until to the end or the visitor returns false
func (sl *Skiplist[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	 //TODO: Complete me!
}

// Keys returns all keys in the skiplist
func (sl *Skiplist[K, V]) Keys() []K {
	 //TODO: Complete me!
}
