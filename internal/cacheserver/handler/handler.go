package handler

import (
	"github.com/google/wire"

	"github.com/onexstack/onex/internal/cacheserver/biz"
	v1 "github.com/onexstack/onex/pkg/api/cacheserver/v1"
)

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewHandler, wire.Bind(new(v1.CacheServerServer), new(*Handler)))

// Handler provides gRPC methods to handle cache operations.
type Handler struct {
	v1.UnimplementedCacheServerServer

	biz biz.IBiz
}

// Ensure that Handler implements the v1.CacheServerServer interface.
var _ v1.CacheServerServer = (*Handler)(nil)

// NewCacheServerHandler creates and returns a new instance of Handler.
func NewHandler(biz biz.IBiz) *Handler {
	return &Handler{biz: biz}
}
