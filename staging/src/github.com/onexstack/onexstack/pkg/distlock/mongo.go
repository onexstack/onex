package distlock

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/onexstack/onexstack/pkg/logger"
)

// MongoLocker provides a distributed locking mechanism using MongoDB.
type MongoLocker struct {
	client         *mongo.Client
	lockCollection *mongo.Collection
	lockName       string
	lockTimeout    time.Duration
	renewTicker    *time.Ticker
	stopChan       chan struct{}
	mu             sync.Mutex
	ownerID        string
	logger         logger.Logger
}

// Ensure MongoLocker implements the Locker interface.
var _ Locker = (*MongoLocker)(nil)

// NewMongoLocker creates a new MongoLocker instance.
func NewMongoLocker(mongoURI string, dbName string, opts ...Option) (*MongoLocker, error) {
	o := ApplyOptions(opts...)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	locker := &MongoLocker{
		client:         client,
		lockCollection: client.Database(dbName).Collection("locks"),
		lockName:       o.lockName,
		lockTimeout:    o.lockTimeout,
		stopChan:       make(chan struct{}),
		ownerID:        o.ownerID,
		logger:         o.logger,
	}

	locker.logger.Info("MongoLocker initialized", "lockName", locker.lockName, "ownerID", locker.ownerID)
	return locker, nil
}

// Lock attempts to acquire the distributed lock.
func (l *MongoLocker) Lock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	expiredAt := now.Add(l.lockTimeout)

	filter := bson.M{"name": l.lockName}
	update := bson.M{
		"$setOnInsert": bson.M{
			"ownerID":   l.ownerID,
			"expiredAt": expiredAt,
		},
		"$set": bson.M{
			"ownerID":   l.ownerID,
			"expiredAt": expiredAt,
		},
	}

	opts := options.Update().SetUpsert(true)
	result, err := l.lockCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		l.logger.Error("Failed to acquire lock", "error", err)
		return fmt.Errorf("failed to acquire lock: %v", err)
	}

	if result.MatchedCount == 0 {
		l.logger.Warn("Lock is already held by another owner", "lockName", l.lockName)
		return fmt.Errorf("lock is already held by another owner")
	}

	l.renewTicker = time.NewTicker(l.lockTimeout / 2)
	go l.renewLock(ctx)

	l.logger.Info("Lock acquired", "ownerID", l.ownerID)
	return nil
}

// Unlock releases the distributed lock.
func (l *MongoLocker) Unlock(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.renewTicker != nil {
		l.renewTicker.Stop()
		l.renewTicker = nil
		l.logger.Info("Stopped renewing lock", "lockName", l.lockName)
	}

	_, err := l.lockCollection.DeleteOne(ctx, bson.M{"name": l.lockName})
	if err != nil {
		l.logger.Error("Failed to release lock", "error", err)
		return fmt.Errorf("failed to release lock: %v", err)
	}

	l.logger.Info("Lock released", "ownerID", l.ownerID)
	l.ownerID = ""
	return nil
}

// Renew refreshes the lock's expiration time.
func (l *MongoLocker) Renew(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	expiredAt := now.Add(l.lockTimeout)

	_, err := l.lockCollection.UpdateOne(ctx, bson.M{"name": l.lockName}, bson.M{"$set": bson.M{"expiredAt": expiredAt}})
	if err != nil {
		l.logger.Error("Failed to renew lock", "error", err)
		return fmt.Errorf("failed to renew lock: %v", err)
	}

	l.logger.Info("Lock renewed", "ownerID", l.ownerID)
	return nil
}

// renewLock periodically renews the lock.
func (l *MongoLocker) renewLock(ctx context.Context) {
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
