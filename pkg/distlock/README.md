# 分布式锁

支持以下类型的分布式锁：
- Noop（假的，勿使用）
- MySQL
- PostgreSQL
- Redis
- Etcd
- Zookeeper
- Consul
- Memcached
- MongoDB

## 测试情况

- 已测试：MySQL、PostgreSQL、Redis
- 未测试（使用前建议你自己充分测试下）：Etcd、Zookeeper、Consul、Memcached、MongoDB

## GPT Prompt

```
package distlock

import (
    "context"
    "fmt"
    "sync"
    "time"

    "github.com/go-sql-driver/mysql"
    "github.com/jackc/pgx/v5/pgconn"
    "gorm.io/gorm"

    "github.com/superproj/onex/pkg/logger"
)

// GORMLocker provides a distributed locking mechanism using GORM.
type GORMLocker struct {
    db          *gorm.DB
    lockName    string
    lockTimeout time.Duration
    renewTicker *time.Ticker
    stopChan    chan struct{}
    mu          sync.Mutex
    ownerID     string
    logger      logger.Logger
}

// Lock represents a database record for a distributed lock.
type Lock struct {
    ID        uint   `gorm:"primarykey"`
    Name      string `gorm:"unique"`
    OwnerID   string
    ExpiredAt time.Time
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Ensure GORMLocker implements the Locker interface.
var _ Locker = (*GORMLocker)(nil)

// NewGORMLocker initializes a new GORMLocker instance.
func NewGORMLocker(db *gorm.DB, opts ...Option) (*GORMLocker, error) {
    o := ApplyOptions(opts...)

    if err := db.AutoMigrate(&Lock{}); err != nil {
        return nil, err
    }

    locker := &GORMLocker{
        db:          db,
        ownerID:     o.ownerID,
        lockName:    o.lockName,
        lockTimeout: o.lockTimeout,
        stopChan:    make(chan struct{}),
        logger:      o.logger,
    }

    locker.logger.Info("GORMLocker initialized", "lockName", locker.lockName, "ownerID", locker.ownerID)

    return locker, nil
}

// Lock acquires the distributed lock.
func (l *GORMLocker) Lock(ctx context.Context) error {
    l.mu.Lock()
    defer l.mu.Unlock()

    now := time.Now()
    expiredAt := now.Add(l.lockTimeout)

    err := l.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(&Lock{Name: l.lockName, OwnerID: l.ownerID, ExpiredAt: expiredAt}).Error; err != nil {
            if !isDuplicateEntry(err) {
                l.logger.Error("failed to create lock", "error", err)
                return err
            }

            var lock Lock
            if err := tx.First(&lock, "name = ?", l.lockName).Error; err != nil {
                l.logger.Error("failed to fetch existing lock", "error", err)
                return err
            }

            if !lock.ExpiredAt.Before(now) {
                l.logger.Warn("lock is already held by another owner", "ownerID", lock.OwnerID)
                return fmt.Errorf("lock is already held by %s", lock.OwnerID)
            }

            lock.OwnerID = l.ownerID
            lock.ExpiredAt = expiredAt
            if err := tx.Save(&lock).Error; err != nil {
                l.logger.Error("failed to update expired lock", "error", err)
                return err
            }
            l.logger.Info("Lock expired, updated owner", "lockName", l.lockName, "newOwnerID", l.ownerID)
        }

        l.renewTicker = time.NewTicker(l.lockTimeout / 2)
        go l.renewLock(ctx)

        l.logger.Info("Lock acquired", "lockName", l.lockName, "ownerID", l.ownerID)
        return nil
    })

    return err
}

// Unlock releases the distributed lock.
func (l *GORMLocker) Unlock(ctx context.Context) error {
    l.mu.Lock()
    defer l.mu.Unlock()

    if l.renewTicker != nil {
        l.renewTicker.Stop()
        l.renewTicker = nil
        l.logger.Info("Stopped renewing lock", "lockName", l.lockName)
    }

    err := l.db.Delete(&Lock{}, "name = ?", l.lockName).Error
    if err != nil {
        l.logger.Error("failed to delete lock", "error", err)
        return err
    }

    l.logger.Info("Lock released", "lockName", l.lockName)
    return nil
}

// Renew refreshes the lease for the distributed lock.
func (l *GORMLocker) Renew(ctx context.Context) error {
    l.mu.Lock()
    defer l.mu.Unlock()

    now := time.Now()
    expiredAt := now.Add(l.lockTimeout)

    err := l.db.Model(&Lock{}).Where("name = ?", l.lockName).Update("expired_at", expiredAt).Error
    if err != nil {
        l.logger.Error("failed to renew lock", "error", err)
        return err
    }

    l.logger.Info("Lock renewed", "lockName", l.lockName, "newExpiration", expiredAt)
    return nil
}

// renewLock periodically renews the lock lease.
func (l *GORMLocker) renewLock(ctx context.Context) {
    for {
        select {
        case <-l.stopChan:
            return
        case <-l.renewTicker.C:
            if err := l.Renew(ctx); err != nil {
                l.logger.Error("failed to renew lock", "error", err)
            }
        }
    }
}

// isDuplicateEntry checks if the error is a duplicate entry error for MySQL and PostgreSQL.
func isDuplicateEntry(err error) bool {
    if err == nil {
        return false
    }

    if mysqlErr, ok := err.(*mysql.MySQLError); ok {
        return mysqlErr.Number == 1062 // MySQL error code for duplicate entry
    }

    if pgErr, ok := err.(*pgconn.PgError); ok {
        return pgErr.Code == "23505" // PostgreSQL error code for unique violation
    }

    return false
}

参考上述基于MySQL的分布式锁实现，使用Consul实现一个分布式锁，该分布式锁同样实现了以下接口：

type Locker interface {
    Lock(ctx context.Context) error
    Unlock(ctx context.Context) error
    Renew(ctx context.Context) error
}

并且需要满足以下要求：
1. 分布式锁启动后，会有一个异步的协程序，根据创建时的参数，定期续锁；
2. 如果代码需要创建数据库表，需要使用gorm的AutoMigrate方法自动创建表
3. 使用logger记录必要的日志
```
