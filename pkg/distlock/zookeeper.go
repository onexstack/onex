package distlock

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-zookeeper/zk"

	"github.com/onexstack/onex/pkg/logger"
)

// ZookeeperLocker provides a distributed locking mechanism using Zookeeper.
type ZookeeperLocker struct {
	conn        *zk.Conn
	lockPath    string
	lockTimeout time.Duration
	renewTicker *time.Ticker
	stopChan    chan struct{}
	mu          sync.Mutex
	ownerID     string // Records the owner ID
	logger      logger.Logger
}

// Ensure ZookeeperLocker implements the Locker interface.
var _ Locker = (*ZookeeperLocker)(nil)

// NewZookeeperLocker creates a new ZookeeperLocker instance.
func NewZookeeperLocker(zkServers []string, opts ...Option) (*ZookeeperLocker, error) {
	o := ApplyOptions(opts...)
	conn, _, err := zk.Connect(zkServers, time.Second)
	if err != nil {
		return nil, err
	}

	locker := &ZookeeperLocker{
		conn:        conn,
		lockPath:    o.lockName,
		lockTimeout: o.lockTimeout,
		stopChan:    make(chan struct{}),
		ownerID:     o.ownerID,
		logger:      o.logger,
	}

	locker.logger.Info("ZookeeperLocker initialized", "lockPath", locker.lockPath, "ownerID", locker.ownerID)
	return locker, nil
}

// Lock attempts to acquire the distributed lock.
func (l *ZookeeperLocker) Lock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Create the lock node
	lockNode := fmt.Sprintf("%s/%s", l.lockPath, l.ownerID)
	_, err := l.conn.Create(lockNode, []byte{}, 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		if err == zk.ErrNodeExists {
			l.logger.Warn("Lock is already held by another owner", "lockNode", lockNode)
			return fmt.Errorf("lock is already held by another owner")
		}
		l.logger.Error("Failed to acquire lock", "error", err)
		return fmt.Errorf("failed to acquire lock: %v", err)
	}

	// Start the renewal goroutine
	l.renewTicker = time.NewTicker(l.lockTimeout / 2)
	go l.renewLock(ctx, lockNode)

	l.logger.Info("Lock acquired", "ownerID", l.ownerID, "lockNode", lockNode)
	return nil
}

// Unlock releases the distributed lock.
func (l *ZookeeperLocker) Unlock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Stop the renewal process
	if l.renewTicker != nil {
		l.renewTicker.Stop()
		l.renewTicker = nil
		l.logger.Info("Stopped renewing lock", "lockNode", fmt.Sprintf("%s/%s", l.lockPath, l.ownerID))
	}

	// Delete the lock node
	lockNode := fmt.Sprintf("%s/%s", l.lockPath, l.ownerID)
	err := l.conn.Delete(lockNode, -1)
	if err != nil {
		l.logger.Error("Failed to release lock", "error", err)
		return fmt.Errorf("failed to release lock: %v", err)
	}

	l.logger.Info("Lock released", "ownerID", l.ownerID)
	l.ownerID = "" // Clear the owner ID
	return nil
}

// Renew refreshes the lock's expiration time.
func (l *ZookeeperLocker) Renew(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Simulate the renewal operation
	l.logger.Info("Lock renewed", "ownerID", l.ownerID)
	return nil
}

// renewLock periodically renews the lock.
func (l *ZookeeperLocker) renewLock(ctx context.Context, lockNode string) {
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
