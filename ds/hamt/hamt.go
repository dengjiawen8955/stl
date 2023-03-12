package hamt

import (
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
type BitmapNode struct {
	 //TODO: Complete me!
}

// KvPair is a list node with actually value
type KvPair struct {
	 //TODO: Complete me!
}

//KvNode is Hamt's key-value node
type KvNode struct {
	 //TODO: Complete me!
}

// Hamt is an implementation of hash-array-mapped-trie
type Hamt struct {
	 //TODO: Complete me!
}

// Type returns the node type
func (h *BitmapNode) Type() int {
	 //TODO: Complete me!
}

// BitPosNum returns the number from a bit position
func (h *BitmapNode) BitPosNum(int) uint64 {
	 //TODO: Complete me!
}

// Index returns the index of a bitPos int bitmap
func (h *BitmapNode) Index(bitPos uint64) int {
	 //TODO: Complete me!
}

func (h *BitmapNode) insert(depth int, hash uint64, kv *KvPair) {
	 //TODO: Complete me!
}

func (h *BitmapNode) find(depth int, hash uint64, key Key) interface{} {
	 //TODO: Complete me!
}

func (h *BitmapNode) traversal(visitor visitor.KvVisitor) {
	 //TODO: Complete me!
}

func (h *BitmapNode) erase(depth int, hash uint64, key Key) bool {
	 //TODO: Complete me!
}

// Type returns the node type
func (h *KvNode) Type() int {
	 //TODO: Complete me!
}

// BitPosNum returns the bit position
func (h *KvNode) BitPosNum(depth int) uint64 {
	 //TODO: Complete me!
}

// New creates a Hamt(hash array mapped trie) instance
func New(opts ...Option) *Hamt {
	 //TODO: Complete me!
}

// Insert inserts a key-value pair into the hamt
func (h *Hamt) Insert(key Key, value interface{}) {
	 //TODO: Complete me!
}

// Get returns the value by the passed key if the key is in the hamt, otherwise returns nil
func (h *Hamt) Get(key Key) interface{} {
	 //TODO: Complete me!
}

// Erase erases the key-value pair in hamt, and returns true if succeed.
func (h *Hamt) Erase(key Key) bool {
	 //TODO: Complete me!
}

// Keys returns keys in Hamt
func (h *Hamt) Keys() []Key {
	 //TODO: Complete me!
}

// StringKeys returns keys in Hamt
func (h *Hamt) StringKeys() []string {
	 //TODO: Complete me!
}

// Traversal traversals elements in Hamt, it will not stop until to the end or the visitor returns false
func (h *Hamt) Traversal(visitor visitor.KvVisitor) {
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
