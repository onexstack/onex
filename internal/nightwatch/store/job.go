// nolint: dupl
package store

import (
	"context"

	storelogger "github.com/onexstack/onexstack/pkg/log/logger/store"
	genericstore "github.com/onexstack/onexstack/pkg/store"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/onex/internal/nightwatch/model"
)

// JobStore defines the interface for managing job-related data operations.
type JobStore interface {
	// Create inserts a new Job record into the store.
	Create(ctx context.Context, obj *model.JobM) error

	// Update modifies an existing Job record in the store based on the given model.
	Update(ctx context.Context, obj *model.JobM) error

	// Delete removes Job records that satisfy the given query options.
	Delete(ctx context.Context, opts *where.Options) error

	// Get retrieves a single Job record that satisfies the given query options.
	Get(ctx context.Context, opts *where.Options) (*model.JobM, error)

	// List retrieves a list of Job records and their total count based on the given query options.
	List(ctx context.Context, opts *where.Options) (int64, []*model.JobM, error)

	// JobExpansion is a placeholder for extension methods for jobs,
	// to be implemented by additional interfaces if needed.
	JobExpansion
}

// JobExpansion is an empty interface provided for extending
// the JobStore interface.
// Developers can define job-specific additional methods
// in this interface for future expansion.
type JobExpansion interface{}

// jobStore implements the JobStore interface and provides
// default implementations of the methods.
type jobStore struct {
	*genericstore.Store[model.JobM]
}

// Ensure that jobStore satisfies the JobStore interface at compile time.
var _ JobStore = (*jobStore)(nil)

// newJobStore creates a new jobStore instance with the provided
// datastore and logger.
func newJobStore(store *datastore) *jobStore {
	return &jobStore{
		Store: genericstore.NewStore[model.JobM](store, storelogger.NewLogger()),
	}
}
