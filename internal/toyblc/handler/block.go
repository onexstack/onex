package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/onexstack/onexstack/pkg/core"

	"github.com/onexstack/onex/internal/toyblc/pkg/miner"
	v1 "github.com/onexstack/onex/pkg/api/toyblc/v1"
)

// CreateBlock handles the creation of a new block.
func (h *Handler) CreateBlock(c *gin.Context) {
	var rq v1.CreateBlockRequest
	if err := core.ShouldBindJSON(c, &rq); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	_ = miner.MinerBlock(h.bs, h.ss, rq.Data)
	core.WriteResponse(c, nil, nil)
}

// ListBlock retrieves a list of blocks based on query parameters.
func (h *Handler) ListBlock(c *gin.Context) {
	core.WriteResponse(c, h.bs.List(), nil)
}
