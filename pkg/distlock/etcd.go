package distlock

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.etcd.io/etcd/client/v3"

	"github.com/onexstack/onex/pkg/logger"
)

// EtcdLocker provides a distributed locking mechanism using etcd.
type EtcdLocker struct {
	cli         *clientv3.Client
	lease       clientv3.Lease
	leaseID     clientv3.LeaseID
	lockKey     string
	lockTimeout time.Duration
	renewTicker *time.Ticker
	stopChan    chan struct{}
	mu          sync.Mutex
	ownerID     string
	logger      logger.Logger
}

// Ensure EtcdLocker implements the Locker interface.
var _ Locker = (*EtcdLocker)(nil)

// NewEtcdLocker initializes a new EtcdLocker instance.
func NewEtcdLocker(endpoints []string, opts ...Option) (*EtcdLocker, error) {
	o := ApplyOptions(opts...)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	lease := clientv3.NewLease(cli)

	locker := &EtcdLocker{
		cli:         cli,
		lease:       lease,
		lockKey:     o.lockName,
		lockTimeout: o.lockTimeout,
		stopChan:    make(chan struct{}),
		ownerID:     o.ownerID,
		logger:      o.logger,
	}

	return locker, nil
}

// Lock acquires the distributed lock.
func (l *EtcdLocker) Lock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	leaseResp, err := l.lease.Grant(ctx, int64(l.lockTimeout.Seconds()))
	if err != nil {
		return err
	}

	l.leaseID = leaseResp.ID

	_, err = l.cli.Put(ctx, l.lockKey, l.ownerID, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %v", err)
	}

	l.renewTicker = time.NewTicker(l.lockTimeout / 2)
	go l.renewLock(ctx, leaseResp.ID)

	l.logger.Info("Lock acquired", "lockKey", l.lockKey)
	return nil
}

// Unlock releases the distributed lock.
func (l *EtcdLocker) Unlock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.renewTicker != nil {
		l.renewTicker.Stop()
		l.renewTicker = nil
	}

	_, err := l.cli.Delete(ctx, l.lockKey)
	if err != nil {
		return err
	}

	if _, err := l.lease.Revoke(context.Background(), l.leaseID); err != nil {
		return fmt.Errorf("failed to revoke lease: %w", err)
	}

	l.logger.Info("Lock released", "lockKey", l.lockKey)
	return nil
}

// Renew refreshes the lease for the distributed lock.
func (l *EtcdLocker) Renew(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, err := l.lease.KeepAliveOnce(ctx, l.leaseID)
	return err
}

// renewLock periodically renews the lock lease.
func (l *EtcdLocker) renewLock(ctx context.Context, leaseID clientv3.LeaseID) {
	for {
		select {
		case <-l.stopChan:
			return
		case <-l.renewTicker.C:
			if err := l.Renew(ctx); err != nil {
				l.logger.Error("failed to renew lock", "err", err)
			} else {
				l.logger.Info("Lock renewed", "lockKey", l.lockKey)
			}
		}
	}
}
