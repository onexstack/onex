// nolint: dupl
package store

import (
	"context"

	storelogger "github.com/onexstack/onexstack/pkg/log/logger/store"
	genericstore "github.com/onexstack/onexstack/pkg/store"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/onex/internal/usercenter/model"
)

// UserStore defines the interface for managing user-related data operations.
type UserStore interface {
	// Create inserts a new User record into the store.
	Create(ctx context.Context, obj *model.UserM) error

	// Update modifies an existing User record in the store based on the given model.
	Update(ctx context.Context, obj *model.UserM) error

	// Delete removes User records that satisfy the given query options.
	Delete(ctx context.Context, opts *where.Options) error

	// Get retrieves a single User record that satisfies the given query options.
	Get(ctx context.Context, opts *where.Options) (*model.UserM, error)

	// List retrieves a list of User records and their total count based on the given query options.
	List(ctx context.Context, opts *where.Options) (int64, []*model.UserM, error)

	// UserExpansion is a placeholder for extension methods for users,
	// to be implemented by additional interfaces if needed.
	UserExpansion
}

// UserExpansion is an empty interface provided for extending
// the UserStore interface.
// Developers can define user-specific additional methods
// in this interface for future expansion.
type UserExpansion interface{}

// userStore implements the UserStore interface and provides
// default implementations of the methods.
type userStore struct {
	*genericstore.Store[model.UserM]
}

// Ensure that userStore satisfies the UserStore interface at compile time.
var _ UserStore = (*userStore)(nil)

// newUserStore creates a new userStore instance with the provided
// datastore and logger.
func newUserStore(store *datastore) *userStore {
	return &userStore{
		Store: genericstore.NewStore[model.UserM](store, storelogger.NewLogger()),
	}
}
