package hamt

import (
	"errors"
	"hash/fnv"
	"math/bits"
	"stl/utils/sync"
	"stl/utils/visitor"
	gosync "sync"
)

// Some constants
const (
	BITMAP_NODE = 0
	KV_NODE     = 1
	Fanout      = 6 //each bitmap node has 6 bits, so the max depth of tree is 64/6 = 10.666 = 11
	Mask        = (1 << Fanout) - 1
)

var ErrorNotFound = errors.New("not found")

// Key is a redefinition of []byte
type Key []byte

var (
	defaultLocker sync.FakeLocker
)

// Options holds Hamt's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to config a Hamt with goroutine-safe
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// Entry is a tree node interface
type Entry interface {
	// Type returns the node type
	Type() int

	// BitPosNum returns number from a bit position
	BitPosNum(depth int) uint64
}

// BitmapNode defines Hamt's bitmap node
type BitmapNode[T any] struct {
	 //TODO: Complete me!
}

// KvPair is a list node with actually value
type KvPair[T any] struct {
	 //TODO: Complete me!
}

//KvNode is Hamt's key-value node
type KvNode[T any] struct {
	 //TODO: Complete me!
}

// Hamt is an implementation of hash-array-mapped-trie
type Hamt[T any] struct {
	 //TODO: Complete me!
}

// Type returns the node type
func (h *BitmapNode[T]) Type() int {
	 //TODO: Complete me!
}

// BitPosNum returns the number from a bit position
func (h *BitmapNode[T]) BitPosNum(int) uint64 {
	 //TODO: Complete me!
}

// Index returns the index of a bitPos int bitmap
func (h *BitmapNode[T]) Index(bitPos uint64) int {
	 //TODO: Complete me!
}

func (h *BitmapNode[T]) insert(depth int, hash uint64, kv *KvPair[T]) {
	 //TODO: Complete me!
}

func (h *BitmapNode[T]) find(depth int, hash uint64, key Key) (T, error) {
	 //TODO: Complete me!
}

func (h *BitmapNode[T]) traversal(visitor visitor.KvVisitor[Key, T]) {
	 //TODO: Complete me!
}

func (h *BitmapNode[T]) erase(depth int, hash uint64, key Key) bool {
	 //TODO: Complete me!
}

// Type returns the node type
func (h *KvNode[T]) Type() int {
	 //TODO: Complete me!
}

// BitPosNum returns the bit position
func (h *KvNode[T]) BitPosNum(depth int) uint64 {
	 //TODO: Complete me!
}

// New creates a Hamt(hash array mapped trie) instance
func New[T any](opts ...Option) *Hamt[T] {
	 //TODO: Complete me!
}

// Insert inserts a key-value pair into the hamt
func (h *Hamt[T]) Insert(key Key, value T) {
	 //TODO: Complete me!
}

// Get returns the value by the passed key if the key is in the hamt, otherwise returns nil
func (h *Hamt[T]) Get(key Key) (T, error) {
	 //TODO: Complete me!
}

// Erase erases the key-value pair in hamt, and returns true if succeed.
func (h *Hamt[T]) Erase(key Key) bool {
	 //TODO: Complete me!
}

// Keys returns keys in Hamt
func (h *Hamt[T]) Keys() []Key {
	 //TODO: Complete me!
}

// StringKeys returns keys in Hamt
func (h *Hamt[T]) StringKeys() []string {
	 //TODO: Complete me!
}

// Traversal traversals elements in Hamt, it will not stop until to the end or the visitor returns false
func (h *Hamt[T]) Traversal(visitor visitor.KvVisitor[Key, T]) {
	 //TODO: Complete me!
}

func hash(a []byte) uint64 {
	 //TODO: Complete me!
}

func pos(hash uint64, depth int) uint8 {
	 //TODO: Complete me!
}

func bitPos(pos uint8) uint64 {
	 //TODO: Complete me!
}
