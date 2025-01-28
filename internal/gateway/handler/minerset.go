package handler

import (
	"context"

	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
	"github.com/onexstack/onex/pkg/apis/apps/v1beta1"
)

// CreateMinerSet handles the creation of a new minerset.
func (h *Handler) CreateMinerSet(ctx context.Context, rq *v1beta1.MinerSet) (*v1beta1.MinerSet, error) {
	return h.biz.MinerSetV1().Create(ctx, rq)
}

// UpdateMinerSet handles updating an existing minerset's details.
func (h *Handler) UpdateMinerSet(ctx context.Context, rq *v1beta1.MinerSet) (*v1beta1.MinerSet, error) {
	return h.biz.MinerSetV1().Update(ctx, rq)
}

// DeleteMinerSet handles the deletion of one or more minersets.
func (h *Handler) DeleteMinerSet(ctx context.Context, rq *v1.DeleteMinerSetRequest) (*v1.DeleteMinerSetResponse, error) {
	return h.biz.MinerSetV1().Delete(ctx, rq)
}

// GetMinerSet retrieves information about a specific minerset.
func (h *Handler) GetMinerSet(ctx context.Context, rq *v1.GetMinerSetRequest) (*v1beta1.MinerSet, error) {
	return h.biz.MinerSetV1().Get(ctx, rq)
}

// ListMinerSet retrieves a list of minersets based on query parameters.
func (h *Handler) ListMinerSet(ctx context.Context, rq *v1.ListMinerSetRequest) (*v1.ListMinerSetResponse, error) {
	return h.biz.MinerSetV1().List(ctx, rq)
}

// ScaleMinerSet handles scaling operations for a specific minerset.
func (h *Handler) ScaleMinerSet(ctx context.Context, rq *v1.ScaleMinerSetRequest) (*v1.ScaleMinerSetResponse, error) {
	return h.biz.MinerSetV1().Scale(ctx, rq)
}
