package handler

import (
	"golang.org/x/net/websocket"

	"github.com/onexstack/onex/internal/toyblc/pkg/ws"
)

func (h *Handler) WSHandler(w *websocket.Conn) {
	ws.WSHandler(h.bs, h.ss, w)
}
