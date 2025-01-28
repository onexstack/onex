package where

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	// defaultLimit defines the default limit for pagination.
	defaultLimit = -1
)

// Tenant represents a tenant with a key and a function to retrieve its value.
type Tenant struct {
	Key       string                           // The key associated with the tenant
	ValueFunc func(ctx context.Context) string // Function to retrieve the tenant's value based on the context
}

type Where interface {
	Where(db *gorm.DB) *gorm.DB
}

// Option defines a function type that modifies Options.
type Option func(*Options)

// Options holds the options for GORM's Where query conditions.
type Options struct {
	// Offset defines the starting point for pagination.
	// +optional
	Offset int `json:"offset"`
	// Limit defines the maximum number of results to return.
	// +optional
	Limit int `json:"limit"`
	// Filters contains key-value pairs for filtering records.
	Filters map[any]any
	// Clauses contains custom clauses to be appended to the query.
	Clauses []clause.Expression
}

// tenant holds the registered tenant instance.
var registeredTenant Tenant

// WithOffset initializes the Offset field in Options with the given offset value.
func WithOffset(offset int64) Option {
	return func(whr *Options) {
		if offset < 0 {
			offset = 0
		}
		whr.Offset = int(offset)
	}
}

// WithLimit initializes the Limit field in Options with the given limit value.
func WithLimit(limit int64) Option {
	return func(whr *Options) {
		if limit <= 0 {
			limit = defaultLimit
		}
		whr.Limit = int(limit)
	}
}

// WithPage is a sugar function to convert page and pageSize into limit and offset in Options.
// This function is commonly used in business logic to facilitate pagination.
func WithPage(page int, pageSize int) Option {
	return func(whr *Options) {
		if page == 0 {
			page = 1
		}
		if pageSize == 0 {
			pageSize = defaultLimit
		}

		whr.Offset = (page - 1) * pageSize
		whr.Limit = pageSize
	}
}

// WithFilter initializes the Filters field in Options with the given filter criteria.
func WithFilter(filter map[any]any) Option {
	return func(whr *Options) {
		whr.Filters = filter
	}
}

// WithClauses appends clauses to the Clauses field in Options.
func WithClauses(conds ...clause.Expression) Option {
	return func(whr *Options) {
		whr.Clauses = append(whr.Clauses, conds...)
	}
}

// NewWhere constructs a new Options object, applying the given where options.
func NewWhere(opts ...Option) *Options {
	whr := &Options{
		Offset:  0,
		Limit:   defaultLimit,
		Filters: map[any]any{},
		Clauses: make([]clause.Expression, 0),
	}

	for _, opt := range opts {
		opt(whr) // Apply each Option to the opts.
	}

	return whr
}

// O sets the offset for the query.
func (whr *Options) O(offset int) *Options {
	if offset < 0 {
		offset = 0
	}
	whr.Offset = offset
	return whr
}

// L sets the limit for the query.
func (whr *Options) L(limit int) *Options {
	if limit <= 0 {
		limit = defaultLimit // Ensure defaultLimit is defined elsewhere
	}
	whr.Limit = limit
	return whr
}

// P sets the pagination based on the page number and page size.
func (whr *Options) P(page int, pageSize int) *Options {
	if page < 1 {
		page = 1 // Ensure page is at least 1
	}
	if pageSize <= 0 {
		pageSize = defaultLimit // Ensure defaultLimit is defined elsewhere
	}
	whr.Offset = (page - 1) * pageSize
	whr.Limit = pageSize
	return whr
}

// C adds conditions to the query.
func (whr *Options) C(conds ...clause.Expression) *Options {
	whr.Clauses = append(whr.Clauses, conds...)
	return whr
}

// T retrieves the value associated with the registered tenant using the provided context.
func (whr *Options) T(ctx context.Context) *Options {
	if registeredTenant.Key != "" && registeredTenant.ValueFunc != nil {
		whr.F(registeredTenant.Key, registeredTenant.ValueFunc(ctx))
	}
	return whr
}

// F adds filters to the query.
func (whr *Options) F(kvs ...any) *Options {
	if len(kvs)%2 != 0 {
		// Handle error: uneven number of key-value pairs
		return whr
	}

	for i := 0; i < len(kvs); i += 2 {
		key := kvs[i]
		value := kvs[i+1]
		whr.Filters[key] = value
	}

	return whr
}

// Where applies the filters and clauses to the given gorm.DB instance.
func (whr *Options) Where(db *gorm.DB) *gorm.DB {
	return db.Where(whr.Filters).Clauses(whr.Clauses...).Offset(whr.Offset).Limit(whr.Limit)
}

// O is a convenience function to create a new Options with offset.
func O(offset int) *Options {
	return NewWhere().O(offset)
}

// L is a convenience function to create a new Options with limit.
func L(limit int) *Options {
	return NewWhere().L(limit)
}

// P is a convenience function to create a new Options with page number and page size.
func P(page int, pageSize int) *Options {
	return NewWhere().P(page, pageSize)
}

// C is a convenience function to create a new Options with conditions.
func C(conds ...clause.Expression) *Options {
	return NewWhere().C(conds...)
}

// T is a convenience function to create a new Options with tenant.
func T(ctx context.Context) *Options {
	return NewWhere().F(registeredTenant.Key, registeredTenant.ValueFunc(ctx))
}

// F is a convenience function to create a new Options with filters.
func F(kvs ...any) *Options {
	return NewWhere().F(kvs...)
}

// RegisterTenant registers a new tenant with the specified key and value function.
func RegisterTenant(key string, valueFunc func(context.Context) string) {
	registeredTenant = Tenant{
		Key:       key,
		ValueFunc: valueFunc,
	}
}
