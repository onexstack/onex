package handler

import (
	"context"

	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// CreateSecret handles the creation of a new secret.
func (h *Handler) CreateSecret(ctx context.Context, rq *v1.CreateSecretRequest) (*v1.CreateSecretResponse, error) {
	return h.biz.SecretV1().Create(ctx, rq)
}

// UpdateSecret handles updating an existing secret's details.
func (h *Handler) UpdateSecret(ctx context.Context, rq *v1.UpdateSecretRequest) (*v1.UpdateSecretResponse, error) {
	return h.biz.SecretV1().Update(ctx, rq)
}

// DeleteSecret handles the deletion of one or more secrets.
func (h *Handler) DeleteSecret(ctx context.Context, rq *v1.DeleteSecretRequest) (*v1.DeleteSecretResponse, error) {
	return h.biz.SecretV1().Delete(ctx, rq)
}

// GetSecret retrieves information about a specific secret.
func (h *Handler) GetSecret(ctx context.Context, rq *v1.GetSecretRequest) (*v1.GetSecretResponse, error) {
	return h.biz.SecretV1().Get(ctx, rq)
}

// ListSecret retrieves a list of secrets based on query parameters.
func (h *Handler) ListSecret(ctx context.Context, rq *v1.ListSecretRequest) (*v1.ListSecretResponse, error) {
	return h.biz.SecretV1().List(ctx, rq)
}
