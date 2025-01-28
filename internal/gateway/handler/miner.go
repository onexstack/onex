package handler

import (
	"context"

	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
	"github.com/onexstack/onex/pkg/apis/apps/v1beta1"
)

// CreateMiner handles the creation of a new miner.
func (h *Handler) CreateMiner(ctx context.Context, rq *v1beta1.Miner) (*v1beta1.Miner, error) {
	return h.biz.MinerV1().Create(ctx, rq)
}

// UpdateMiner handles updating an existing miner's details.
func (h *Handler) UpdateMiner(ctx context.Context, rq *v1beta1.Miner) (*v1beta1.Miner, error) {
	return h.biz.MinerV1().Update(ctx, rq)
}

// DeleteMiner handles the deletion of one or more miners.
func (h *Handler) DeleteMiner(ctx context.Context, rq *v1.DeleteMinerRequest) (*v1.DeleteMinerResponse, error) {
	return h.biz.MinerV1().Delete(ctx, rq)
}

// GetMiner retrieves information about a specific miner.
func (h *Handler) GetMiner(ctx context.Context, rq *v1.GetMinerRequest) (*v1beta1.Miner, error) {
	return h.biz.MinerV1().Get(ctx, rq)
}

// ListMiner retrieves a list of miners based on query parameters.
func (h *Handler) ListMiner(ctx context.Context, rq *v1.ListMinerRequest) (*v1.ListMinerResponse, error) {
	return h.biz.MinerV1().List(ctx, rq)
}
