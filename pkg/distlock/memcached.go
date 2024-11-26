package distlock

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/bradfitz/gomemcache/memcache"

	"github.com/superproj/onex/pkg/logger"
)

// MemcachedLocker provides a distributed locking mechanism using Memcached.
type MemcachedLocker struct {
	client      *memcache.Client
	lockKey     string
	lockTimeout time.Duration
	renewTicker *time.Ticker
	stopChan    chan struct{}
	mu          sync.Mutex
	ownerID     string
	logger      logger.Logger
}

// Ensure MemcachedLocker implements the Locker interface.
var _ Locker = (*MemcachedLocker)(nil)

// NewMemcachedLocker creates a new MemcachedLocker instance.
func NewMemcachedLocker(memcachedAddr string, opts ...Option) *MemcachedLocker {
	o := ApplyOptions(opts...)
	client := memcache.New(memcachedAddr)
	locker := &MemcachedLocker{
		client:      client,
		lockKey:     o.lockName,
		lockTimeout: o.lockTimeout,
		stopChan:    make(chan struct{}),
		ownerID:     o.ownerID,
		logger:      o.logger,
	}

	locker.logger.Info("MemcachedLocker initialized", "lockKey", locker.lockKey, "ownerID", locker.ownerID)
	return locker
}

// Lock attempts to acquire the distributed lock.
func (l *MemcachedLocker) Lock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Attempt to acquire the lock
	item := &memcache.Item{
		Key:        l.lockKey,
		Value:      []byte(l.ownerID),
		Expiration: int32(l.lockTimeout.Seconds()),
	}

	// Use Add method to acquire the lock, which succeeds only if the key does not exist
	err := l.client.Add(item)
	if err == memcache.ErrNotStored {
		l.logger.Warn("Lock is already held by another owner", "lockKey", l.lockKey)
		return fmt.Errorf("lock is already held by another owner")
	} else if err != nil {
		l.logger.Error("Failed to acquire lock", "error", err)
		return fmt.Errorf("failed to acquire lock: %v", err)
	}

	// Start the renewal goroutine
	l.renewTicker = time.NewTicker(l.lockTimeout / 2)
	go l.renewLock(ctx)

	l.logger.Info("Lock acquired", "ownerID", l.ownerID, "lockKey", l.lockKey)
	return nil
}

// Unlock releases the distributed lock.
func (l *MemcachedLocker) Unlock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Stop renewing the lock
	if l.renewTicker != nil {
		l.renewTicker.Stop()
		l.renewTicker = nil
		l.logger.Info("Stopped renewing lock", "lockKey", l.lockKey)
	}

	// Remove the lock
	err := l.client.Delete(l.lockKey)
	if err != nil {
		l.logger.Error("Failed to release lock", "error", err)
		return fmt.Errorf("failed to release lock: %v", err)
	}

	l.logger.Info("Lock released", "ownerID", l.ownerID)
	return nil
}

// Renew refreshes the expiration time of the lock.
func (l *MemcachedLocker) Renew(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Attempt to renew the lock
	item := &memcache.Item{
		Key:        l.lockKey,
		Value:      []byte(l.ownerID),
		Expiration: int32(l.lockTimeout.Seconds()),
	}

	// Use Replace method to update the expiration time of the lock
	err := l.client.Replace(item)
	if err == memcache.ErrNotStored {
		l.logger.Warn("Lock is not held by this owner anymore", "lockKey", l.lockKey)
		return fmt.Errorf("lock is not held by this owner anymore")
	} else if err != nil {
		l.logger.Error("Failed to renew lock", "error", err)
		return fmt.Errorf("failed to renew lock: %v", err)
	}

	l.logger.Info("Lock renewed", "ownerID", l.ownerID, "lockKey", l.lockKey)
	return nil
}

// renewLock periodically renews the lock.
func (l *MemcachedLocker) renewLock(ctx context.Context) {
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
