package store

//go:generate mockgen -destination mock_store.go -package store onex/internal/nightwatch/store IStore,CronJobStore,JobStore

import (
	"context"
	"sync"

	"github.com/google/wire"
	"github.com/onexstack/onexstack/pkg/store/where"
	"gorm.io/gorm"
)

// ProviderSet is a Wire provider set that declares dependency injection rules.
// It includes the NewStore constructor function to generate datastore instances.
// wire.Bind is used to bind the IStore interface to the concrete implementation *datastore,
// allowing automatic injection of *datastore instances wherever IStore is required.
var ProviderSet = wire.NewSet(NewStore, wire.Bind(new(IStore), new(*datastore)))

var (
	once sync.Once
	// S is a global variable for convenient access to the initialized datastore
	// instance from other packages.
	S *datastore
)

// IStore defines the methods that the Store layer needs to implement.
type IStore interface {
	// DB returns the *gorm.DB instance of the Store layer, which might be used in rare cases.
	DB(ctx context.Context, wheres ...where.Where) *gorm.DB
	// TX is used to implement transactions in the Biz layer.
	TX(ctx context.Context, fn func(ctx context.Context) error) error
	// CronJob returns an implementation of the CronJobStore.
	CronJob() CronJobStore
	// Job returns an implementation of the JobStore.
	Job() JobStore
}

// transactionKey is the key used to store transaction context in context.Context.
type transactionKey struct{}

// datastore is the concrete implementation of the IStore.
type datastore struct {
	core *gorm.DB

	// Additional database instances can be added as needed.
	// Example: fake *gorm.DB
}

// Ensure datastore implements the IStore.
var _ IStore = (*datastore)(nil)

// NewStore initializes a singleton instance of type IStore.
// It ensures that the datastore is only created once using sync.Once.
func NewStore(db *gorm.DB) *datastore {
	// Initialize the singleton datastore instance only once.
	once.Do(func() {
		S = &datastore{db}
	})

	return S
}

// DB filters the database instance based on the input conditions (wheres).
// If no conditions are provided, the function returns the database instance
// from the context (transaction instance or core database instance).
func (store *datastore) DB(ctx context.Context, wheres ...where.Where) *gorm.DB {
	db := store.core
	// Attempt to retrieve the transaction instance from the context.
	if tx, ok := ctx.Value(transactionKey{}).(*gorm.DB); ok {
		db = tx
	}

	// Apply each provided 'where' condition to the query.
	for _, whr := range wheres {
		db = whr.Where(db)
	}
	return db
}

// FakeDB is used to demonstrate multiple database instances.
// It returns a nil gorm.DB, indicating a fake database.
func (ds *datastore) FakeDB(ctx context.Context) *gorm.DB { return nil }

// TX starts a new transaction instance.
// nolint: fatcontext
func (store *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return store.core.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}

// CronJob returns an instance that implements the CronJobStore.
func (store *datastore) CronJob() CronJobStore {
	return newCronJobStore(store)
}

// Job returns an instance that implements the JobStore.
func (store *datastore) Job() JobStore {
	return newJobStore(store)
}
