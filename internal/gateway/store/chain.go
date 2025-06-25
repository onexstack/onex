// nolint: dupl
package store

import (
	"context"

	storelogger "github.com/onexstack/onexstack/pkg/log/logger/store"
	genericstore "github.com/onexstack/onexstack/pkg/store"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/onex/internal/gateway/model"
)

// ChainStore defines the interface for managing chain-related data operations.
type ChainStore interface {
	// Create inserts a new Chain record into the store.
	Create(ctx context.Context, obj *model.ChainM) error

	// Update modifies an existing Chain record in the store based on the given model.
	Update(ctx context.Context, obj *model.ChainM) error

	// Delete removes Chain records that satisfy the given query options.
	Delete(ctx context.Context, opts *where.Options) error

	// Get retrieves a single Chain record that satisfies the given query options.
	Get(ctx context.Context, opts *where.Options) (*model.ChainM, error)

	// List retrieves a list of Chain records and their total count based on the given query options.
	List(ctx context.Context, opts *where.Options) (int64, []*model.ChainM, error)

	// ChainExpansion is a placeholder for extension methods for chains,
	// to be implemented by additional interfaces if needed.
	ChainExpansion
}

// ChainExpansion is an empty interface provided for extending
// the ChainStore interface.
// Developers can define chain-specific additional methods
// in this interface for future expansion.
type ChainExpansion interface{}

// chainStore implements the ChainStore interface and provides
// default implementations of the methods.
type chainStore struct {
	*genericstore.Store[model.ChainM]
}

// Ensure that chainStore satisfies the ChainStore interface at compile time.
var _ ChainStore = (*chainStore)(nil)

// newChainStore creates a new chainStore instance with the provided
// datastore and logger.
func newChainStore(store *datastore) *chainStore {
	return &chainStore{
		Store: genericstore.NewStore[model.ChainM](store, storelogger.NewLogger()),
	}
}
