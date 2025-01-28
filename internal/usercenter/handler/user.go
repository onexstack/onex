package handler

import (
	"context"

	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// CreateUser handles the creation of a new user.
func (h *Handler) CreateUser(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return h.biz.UserV1().Create(ctx, rq)
}

// UpdateUser handles updating an existing user's details.
func (h *Handler) UpdateUser(ctx context.Context, rq *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return h.biz.UserV1().Update(ctx, rq)
}

// DeleteUser handles the deletion of one or more users.
func (h *Handler) DeleteUser(ctx context.Context, rq *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	return h.biz.UserV1().Delete(ctx, rq)
}

// GetUser retrieves information about a specific user.
func (h *Handler) GetUser(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	return h.biz.UserV1().Get(ctx, rq)
}

// ListUser retrieves a list of users based on query parameters.
func (h *Handler) ListUser(ctx context.Context, rq *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	return h.biz.UserV1().List(ctx, rq)
}

// UpdatePassword receives an UpdatePasswordRequest and updates the user's password in the datastore.
func (h *Handler) UpdatePassword(ctx context.Context, rq *v1.UpdatePasswordRequest) (*v1.UpdatePasswordResponse, error) {
	return h.biz.UserV1().UpdatePassword(ctx, rq)
}
