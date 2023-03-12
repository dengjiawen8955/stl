package ketama

import (
	"stl/algorithm/hash"
	"stl/ds/treemap"
	"stl/utils/comparator"
	"stl/utils/sync"
	gosync "sync"
)

var (
	defaultReplicas = 10
	defaultLocker   sync.FakeLocker
)

const salt = "ni9fkh72hgh1g"

// Options hold Ketama's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to config a Ketama with goroutine-safe
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// WithReplicas is used to config the hash replicas of a Ketama
func WithReplicas(replicas int) Option {
	 //TODO: Complete me!
}

// Ketama is an implementation of consistent-hash
type Ketama struct {
	 //TODO: Complete me!
}

// New creates a new ketama
func New(opts ...Option) *Ketama {
	 //TODO: Complete me!
}

// Empty returns true if the ketama is empty, otherwise returns false
func (k *Ketama) Empty() bool {
	 //TODO: Complete me!
}

// Add adds nodes to the ketama ring
func (k *Ketama) Add(nodes ...string) {
	 //TODO: Complete me!
}

// Remove removes nodes from the ketama ring
func (k *Ketama) Remove(nodes ...string) {
	 //TODO: Complete me!
}

// Get returns the node closest to key in the clockwise direction
func (k *Ketama) Get(key string) (string, bool) {
	 //TODO: Complete me!
}
