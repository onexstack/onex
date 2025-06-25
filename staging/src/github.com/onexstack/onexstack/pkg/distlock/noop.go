package distlock

import (
	"context"
	"sync"
	"time"

	"github.com/onexstack/onexstack/pkg/logger"
)

// NoopLocker provides a no-operation implementation of a distributed lock.
type NoopLocker struct {
	lockTimeout time.Duration
	renewTicker *time.Ticker
	stopChan    chan struct{}
	mu          sync.Mutex
	ownerID     string // Records the owner ID
	logger      logger.Logger
}

// Ensure NoopLocker implements the Locker interface.
var _ Locker = (*NoopLocker)(nil)

// NewNoopLocker creates a new NoopLocker instance.
func NewNoopLocker(opts ...Option) *NoopLocker {
	o := ApplyOptions(opts...)
	return &NoopLocker{
		lockTimeout: o.lockTimeout,
		ownerID:     o.ownerID,
		stopChan:    make(chan struct{}),
		logger:      o.logger, // Initialize logger
	}
}

// Lock simulates acquiring a distributed lock.
func (l *NoopLocker) Lock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Start the renewal goroutine
	l.renewTicker = time.NewTicker(l.lockTimeout / 2)
	go l.renewLock(ctx)

	l.logger.Info("Lock acquired", "ownerID", l.ownerID)
	return nil
}

// Unlock simulates releasing a distributed lock.
func (l *NoopLocker) Unlock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Stop the renewal process
	if l.renewTicker != nil {
		l.renewTicker.Stop()
		l.renewTicker = nil
	}

	l.logger.Info("Lock released", "ownerID", l.ownerID)
	l.ownerID = "" // Clear the owner ID
	return nil
}

// Renew simulates refreshing the lock's expiration time.
func (l *NoopLocker) Renew(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Simulate the renewal operation
	l.logger.Info("Lock renewed", "ownerID", l.ownerID)
	return nil
}

// renewLock periodically renews the lock.
func (l *NoopLocker) renewLock(ctx context.Context) {
	for {
		select {
		case <-l.stopChan:
			return
		case <-l.renewTicker.C:
			if err := l.Renew(ctx); err != nil {
				l.logger.Error("Failed to renew lock", "error", err)
			}
		}
	}
}
