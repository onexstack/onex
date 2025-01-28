package handler

import (
	"github.com/onexstack/onex/internal/toyblc/pkg/blc"
	"github.com/onexstack/onex/internal/toyblc/pkg/ws"
)

// Handler implements a gRPC service.
type Handler struct {
	bs *blc.BlockSet
	ss *ws.Sockets
}

// NewHandler creates a new instance of Handler.
func NewHandler(bs *blc.BlockSet, ss *ws.Sockets) *Handler {
	return &Handler{bs: bs, ss: ss}
}
