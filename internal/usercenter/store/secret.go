// nolint: dupl
package store

import (
	"context"

	storelogger "github.com/onexstack/onexstack/pkg/log/logger/store"
	genericstore "github.com/onexstack/onexstack/pkg/store"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/onex/internal/usercenter/model"
)

// SecretStore defines the interface for managing secret-related data operations.
type SecretStore interface {
	// Create inserts a new Secret record into the store.
	Create(ctx context.Context, obj *model.SecretM) error

	// Update modifies an existing Secret record in the store based on the given model.
	Update(ctx context.Context, obj *model.SecretM) error

	// Delete removes Secret records that satisfy the given query options.
	Delete(ctx context.Context, opts *where.Options) error

	// Get retrieves a single Secret record that satisfies the given query options.
	Get(ctx context.Context, opts *where.Options) (*model.SecretM, error)

	// List retrieves a list of Secret records and their total count based on the given query options.
	List(ctx context.Context, opts *where.Options) (int64, []*model.SecretM, error)

	// SecretExpansion is a placeholder for extension methods for secrets,
	// to be implemented by additional interfaces if needed.
	SecretExpansion
}

// SecretExpansion is an empty interface provided for extending
// the SecretStore interface.
// Developers can define secret-specific additional methods
// in this interface for future expansion.
type SecretExpansion interface{}

// secretStore implements the SecretStore interface and provides
// default implementations of the methods.
type secretStore struct {
	*genericstore.Store[model.SecretM]
}

// Ensure that secretStore satisfies the SecretStore interface at compile time.
var _ SecretStore = (*secretStore)(nil)

// newSecretStore creates a new secretStore instance with the provided
// datastore and logger.
func newSecretStore(store *datastore) *secretStore {
	return &secretStore{
		Store: genericstore.NewStore[model.SecretM](store, storelogger.NewLogger()),
	}
}
