package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/onexstack/onexstack/pkg/core"

	"github.com/onexstack/onex/internal/toyblc/pkg/ws"
	v1 "github.com/onexstack/onex/pkg/api/toyblc/v1"
)

// CreatePeer handles the creation of a new peer.
func (h *Handler) CreatePeer(c *gin.Context) {
	var rq v1.CreatePeerRequest
	if err := core.ShouldBindJSON(c, &rq); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	ws.ConnectToPeers(c, h.bs, h.ss, []string{rq.Peer})

	core.WriteResponse(c, nil, nil)
}

// ListPeer retrieves a list of peers based on query parameters.
func (h *Handler) ListPeer(c *gin.Context) {
	var slice []string

	for _, socket := range h.ss.List() {
		if socket.IsClientConn() {
			slice = append(slice, strings.Replace(socket.LocalAddr().String(), "ws://", "", 1))
		} else {
			slice = append(slice, socket.Request().RemoteAddr)
		}
	}

	core.WriteResponse(c, slice, nil)
}
