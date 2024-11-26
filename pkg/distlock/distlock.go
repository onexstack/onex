// Package distlock provides an interface for distributed locking mechanisms.
package distlock

import (
	"context"
	"os"
	"time"

	"github.com/superproj/onex/pkg/logger"
	"github.com/superproj/onex/pkg/logger/empty"
)

// Locker is an interface that defines the methods for a distributed lock.
// It provides methods to acquire, release, and renew a lock in a distributed system.
type Locker interface {
	// Lock attempts to acquire the lock.
	Lock(ctx context.Context) error

	// Unlock releases the previously acquired lock.
	Unlock(ctx context.Context) error

	// Renew updates the expiration time of the lock.
	// It should be called periodically to keep the lock active.
	Renew(ctx context.Context) error
}

// Options holds the configuration for the distributed lock.
type Options struct {
	lockName    string        // Name of the lock
	lockTimeout time.Duration // Duration before the lock expires
	ownerID     string        // Identifier for the lock owner
	logger      logger.Logger // Logger for logging events
}

// Option is a function that modifies Options.
type Option func(o *Options)

// NewOptions initializes Options with default values.
func NewOptions() *Options {
	ownerID, _ := os.Hostname() // Get the hostname as the default owner ID
	return &Options{
		lockName:    "onex-distributed-lock", // Default lock name
		lockTimeout: 10 * time.Second,        // Default lock timeout
		ownerID:     ownerID,                 // Set the owner ID
		logger:      empty.NewLogger(),       // Default logger
	}
}

// ApplyOptions applies a series of Option functions to configure Options.
func ApplyOptions(opts ...Option) *Options {
	o := NewOptions() // Create a new Options instance with default values
	for _, opt := range opts {
		opt(o) // Apply each option to the Options instance
	}

	return o // Return the configured Options
}

// WithLockName sets the lock name in Options.
func WithLockName(name string) Option {
	return func(o *Options) {
		o.lockName = name // Set the lock name
	}
}

// WithLockTimeout sets the lock timeout in Options.
func WithLockTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.lockTimeout = timeout // Set the lock timeout
	}
}

// WithOwnerID sets the owner ID in Options.
func WithOwnerID(ownerID string) Option {
	return func(o *Options) {
		o.ownerID = ownerID // Set the owner ID
	}
}

// WithLogger sets the logger in Options.
func WithLogger(logger logger.Logger) Option {
	return func(o *Options) {
		o.logger = logger // Set the logger
	}
}
