package distlock

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/superproj/onex/pkg/logger"
)

// ConsulLocker is a structure that implements distributed locking using Consul.
type ConsulLocker struct {
	client      *api.Client   // Consul client for interacting with the Consul API
	lockKey     string        // Key for the distributed lock
	lockTimeout time.Duration // Duration for which the lock is valid
	renewTicker *time.Ticker  // Ticker for renewing the lock periodically
	stopChan    chan struct{} // Channel to signal stopping the renewal process
	mu          sync.Mutex    // Mutex for synchronizing access to the locker
	ownerID     string        // Identifier for the owner of the lock
	logger      logger.Logger // Logger for logging events and errors
}

// Ensure ConsulLocker implements the Locker interface
var _ Locker = (*ConsulLocker)(nil)

// NewConsulLocker creates a new ConsulLocker instance.
func NewConsulLocker(consulAddr string, opts ...Option) (*ConsulLocker, error) {
	o := ApplyOptions(opts...)
	// Create a Consul client
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	// Initialize a new ConsulLocker with the provided options
	locker := &ConsulLocker{
		client:      client,
		lockKey:     o.lockName,
		lockTimeout: o.lockTimeout,
		stopChan:    make(chan struct{}),
		ownerID:     o.ownerID,
		logger:      o.logger,
	}

	locker.logger.Info("ConsulLocker initialized", "lockKey", locker.lockKey, "ownerID", locker.ownerID)
	return locker, nil
}

// Lock attempts to acquire the distributed lock.
func (l *ConsulLocker) Lock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Create a new session for the lock with a TTL
	session := &api.SessionEntry{
		TTL:      fmt.Sprintf("%s", l.lockTimeout),
		Behavior: api.SessionBehaviorRelease,
	}

	// Create a session and handle any errors
	sessionID, _, err := l.client.Session().Create(session, nil)
	if err != nil {
		l.logger.Error("Failed to create session", "error", err)
		return fmt.Errorf("failed to create session: %v", err)
	}

	// Create a KV pair for the lock
	kv := &api.KVPair{
		Key:     l.lockKey,
		Value:   []byte(l.ownerID),
		Session: sessionID,
	}

	// Attempt to put the lock in the KV store and handle any errors
	_, err = l.client.KV().Put(kv, nil)
	if err != nil {
		l.logger.Error("Failed to acquire lock", "error", err)
		return fmt.Errorf("failed to acquire lock: %v", err)
	}

	// Start a ticker to renew the lock periodically
	l.renewTicker = time.NewTicker(l.lockTimeout / 2)
	go l.renewLock(ctx, sessionID)

	l.logger.Info("Lock acquired", "ownerID", l.ownerID, "sessionID", sessionID)
	return nil
}

// Unlock releases the distributed lock.
func (l *ConsulLocker) Unlock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Stop the renewal ticker if it is running
	if l.renewTicker != nil {
		l.renewTicker.Stop()
		l.renewTicker = nil
		l.logger.Info("Stopped renewing lock", "lockKey", l.lockKey)
	}

	// Delete the lock from the KV store and handle any errors
	_, err := l.client.KV().Delete(l.lockKey, nil)
	if err != nil {
		l.logger.Error("Failed to release lock", "error", err)
		return fmt.Errorf("failed to release lock: %v", err)
	}

	l.logger.Info("Lock released", "ownerID", l.ownerID)
	return nil
}

// Renew refreshes the lock's expiration time.
func (l *ConsulLocker) Renew(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Renew the session associated with the lock and handle any errors
	_, _, err := l.client.Session().Renew(l.ownerID, nil)
	if err != nil {
		l.logger.Error("Failed to renew lock", "error", err)
		return fmt.Errorf("failed to renew lock: %v", err)
	}

	l.logger.Info("Lock renewed", "ownerID", l.ownerID)
	return nil
}

// renewLock periodically renews the lock.
func (l *ConsulLocker) renewLock(ctx context.Context, sessionID string) {
	for {
		select {
		case <-l.stopChan:
			return
		case <-l.renewTicker.C:
			// Attempt to renew the lock and log any errors
			if err := l.Renew(ctx); err != nil {
				l.logger.Error("Failed to renew lock", "error", err)
			}
		}
	}
}
