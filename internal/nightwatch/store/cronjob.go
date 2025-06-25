// nolint: dupl
package store

import (
	"context"

	storelogger "github.com/onexstack/onexstack/pkg/log/logger/store"
	genericstore "github.com/onexstack/onexstack/pkg/store"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/onex/internal/nightwatch/model"
)

// CronJobStore defines the interface for managing cronjob-related data operations.
type CronJobStore interface {
	// Create inserts a new CronJob record into the store.
	Create(ctx context.Context, obj *model.CronJobM) error

	// Update modifies an existing CronJob record in the store based on the given model.
	Update(ctx context.Context, obj *model.CronJobM) error

	// Delete removes CronJob records that satisfy the given query options.
	Delete(ctx context.Context, opts *where.Options) error

	// Get retrieves a single CronJob record that satisfies the given query options.
	Get(ctx context.Context, opts *where.Options) (*model.CronJobM, error)

	// List retrieves a list of CronJob records and their total count based on the given query options.
	List(ctx context.Context, opts *where.Options) (int64, []*model.CronJobM, error)

	// CronJobExpansion is a placeholder for extension methods for cronjobs,
	// to be implemented by additional interfaces if needed.
	CronJobExpansion
}

// CronJobExpansion is an empty interface provided for extending
// the CronJobStore interface.
// Developers can define cronjob-specific additional methods
// in this interface for future expansion.
type CronJobExpansion interface{}

// cronJobStore implements the CronJobStore interface and provides
// default implementations of the methods.
type cronJobStore struct {
	*genericstore.Store[model.CronJobM]
}

// Ensure that cronJobStore satisfies the CronJobStore interface at compile time.
var _ CronJobStore = (*cronJobStore)(nil)

// newCronJobStore creates a new cronJobStore instance with the provided
// datastore and logger.
func newCronJobStore(store *datastore) *cronJobStore {
	return &cronJobStore{
		Store: genericstore.NewStore[model.CronJobM](store, storelogger.NewLogger()),
	}
}
