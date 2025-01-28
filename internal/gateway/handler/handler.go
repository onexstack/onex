package handler

import (
	"github.com/google/wire"

	"github.com/onexstack/onex/internal/gateway/biz"
	"github.com/onexstack/onex/internal/pkg/idempotent"
	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
)

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewHandler, wire.Bind(new(v1.GatewayServer), new(*Handler)))

// Handler implements a gRPC service.
type Handler struct {
	v1.UnimplementedGatewayServer

	biz biz.IBiz
	idt *idempotent.Idempotent
}

// Ensure that Handler implements the v1.GatewayServer interface.
var _ v1.GatewayServer = (*Handler)(nil)

// NewHandler creates a new instance of *Handler.
func NewHandler(biz biz.IBiz, idt *idempotent.Idempotent) *Handler {
	return &Handler{biz: biz, idt: idt}
}
