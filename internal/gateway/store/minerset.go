// nolint: dupl
package store

import (
	"context"

	storelogger "github.com/onexstack/onexstack/pkg/log/logger/store"
	genericstore "github.com/onexstack/onexstack/pkg/store"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/onex/internal/gateway/model"
)

// MinerSetStore defines the interface for managing minerset-related data operations.
type MinerSetStore interface {
	// Create inserts a new MinerSet record into the store.
	Create(ctx context.Context, obj *model.MinerSetM) error

	// Update modifies an existing MinerSet record in the store based on the given model.
	Update(ctx context.Context, obj *model.MinerSetM) error

	// Delete removes MinerSet records that satisfy the given query options.
	Delete(ctx context.Context, opts *where.Options) error

	// Get retrieves a single MinerSet record that satisfies the given query options.
	Get(ctx context.Context, opts *where.Options) (*model.MinerSetM, error)

	// List retrieves a list of MinerSet records and their total count based on the given query options.
	List(ctx context.Context, opts *where.Options) (int64, []*model.MinerSetM, error)

	// MinerSetExpansion is a placeholder for extension methods for minersets,
	// to be implemented by additional interfaces if needed.
	MinerSetExpansion
}

// MinerSetExpansion is an empty interface provided for extending
// the MinerSetStore interface.
// Developers can define minerset-specific additional methods
// in this interface for future expansion.
type MinerSetExpansion interface{}

// minerSetStore implements the MinerSetStore interface and provides
// default implementations of the methods.
type minerSetStore struct {
	*genericstore.Store[model.MinerSetM]
}

// Ensure that minerSetStore satisfies the MinerSetStore interface at compile time.
var _ MinerSetStore = (*minerSetStore)(nil)

// newMinerSetStore creates a new minerSetStore instance with the provided
// datastore and logger.
func newMinerSetStore(store *datastore) *minerSetStore {
	return &minerSetStore{
		Store: genericstore.NewStore[model.MinerSetM](store, storelogger.NewLogger()),
	}
}
