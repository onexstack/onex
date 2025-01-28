// nolint: dupl
package store

import (
	"context"

	storelogger "github.com/onexstack/onexstack/pkg/log/logger/store"
	genericstore "github.com/onexstack/onexstack/pkg/store"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/onex/internal/gateway/model"
)

// MinerStore defines the interface for managing miner-related data operations.
type MinerStore interface {
	// Create inserts a new Miner record into the store.
	Create(ctx context.Context, obj *model.MinerM) error

	// Update modifies an existing Miner record in the store based on the given model.
	Update(ctx context.Context, obj *model.MinerM) error

	// Delete removes Miner records that satisfy the given query options.
	Delete(ctx context.Context, opts *where.Options) error

	// Get retrieves a single Miner record that satisfies the given query options.
	Get(ctx context.Context, opts *where.Options) (*model.MinerM, error)

	// List retrieves a list of Miner records and their total count based on the given query options.
	List(ctx context.Context, opts *where.Options) (int64, []*model.MinerM, error)

	// MinerExpansion is a placeholder for extension methods for miners,
	// to be implemented by additional interfaces if needed.
	MinerExpansion
}

// MinerExpansion is an empty interface provided for extending
// the MinerStore interface.
// Developers can define miner-specific additional methods
// in this interface for future expansion.
type MinerExpansion interface{}

// minerStore implements the MinerStore interface and provides
// default implementations of the methods.
type minerStore struct {
	*genericstore.Store[model.MinerM]
}

// Ensure that minerStore satisfies the MinerStore interface at compile time.
var _ MinerStore = (*minerStore)(nil)

// newMinerStore creates a new minerStore instance with the provided
// datastore and logger.
func newMinerStore(store *datastore) *minerStore {
	return &minerStore{
		Store: genericstore.NewStore[model.MinerM](store, storelogger.NewLogger()),
	}
}
