package distlock

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/onexstack/onexstack/pkg/logger"
)

// RedisLocker provides a distributed locking mechanism using Redis.
type RedisLocker struct {
	client      *redis.Client
	lockName    string
	lockTimeout time.Duration
	renewTicker *time.Ticker
	stopChan    chan struct{}
	mu          sync.Mutex
	ownerID     string
	logger      logger.Logger
}

// Ensure RedisLocker implements the Locker interface.
var _ Locker = (*RedisLocker)(nil)

// NewRedisLocker creates a new RedisLocker instance.
func NewRedisLocker(client *redis.Client, opts ...Option) *RedisLocker {
	o := ApplyOptions(opts...)
	locker := &RedisLocker{
		client:      client,
		lockName:    o.lockName,
		lockTimeout: o.lockTimeout,
		stopChan:    make(chan struct{}),
		ownerID:     o.ownerID,
		logger:      o.logger,
	}

	locker.logger.Info("RedisLocker initialized", "lockName", locker.lockName, "ownerID", locker.ownerID)
	return locker
}

// Lock attempts to acquire the distributed lock.
func (l *RedisLocker) Lock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	success, err := l.client.SetNX(ctx, l.lockName, l.ownerID, l.lockTimeout).Result()
	if err != nil {
		l.logger.Error("Failed to set lock", "error", err)
		return err
	}
	if !success {
		currentOwnerID, err := l.client.Get(ctx, l.lockName).Result()
		if err != nil {
			l.logger.Error("Failed to get current owner ID", "error", err)
			return err
		}
		if currentOwnerID != l.ownerID {
			l.logger.Warn("Lock is already held by another owner", "currentOwnerID", currentOwnerID)
			return fmt.Errorf("lock is already held by %s", currentOwnerID)
		}
		l.logger.Info("Lock is already held by the current owner, extending the lock if needed")
		return nil
	}

	l.renewTicker = time.NewTicker(l.lockTimeout / 2)
	go l.renewLock(ctx)

	l.logger.Info("Lock acquired", "ownerID", l.ownerID)
	return nil
}

// Unlock releases the distributed lock.
func (l *RedisLocker) Unlock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.renewTicker != nil {
		l.renewTicker.Stop()
		l.renewTicker = nil
		l.logger.Info("Stopped renewing lock", "lockName", l.lockName)
	}

	err := l.client.Del(ctx, l.lockName).Err()
	if err != nil {
		l.logger.Error("Failed to delete lock", "error", err)
		return err
	}

	l.logger.Info("Lock released", "ownerID", l.ownerID)
	return nil
}

// Renew refreshes the lock's expiration time.
func (l *RedisLocker) Renew(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	err := l.client.Expire(ctx, l.lockName, l.lockTimeout).Err()
	if err != nil {
		l.logger.Error("Failed to renew lock", "error", err)
		return err
	}

	l.logger.Info("Lock renewed", "ownerID", l.ownerID)
	return nil
}

// renewLock periodically renews the lock.
func (l *RedisLocker) renewLock(ctx context.Context) {
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
