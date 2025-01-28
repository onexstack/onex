package handler

import (
	"context"

	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// Login authenticates the user credentials and returns a token on success.
func (h *Handler) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	return h.biz.AuthV1().Login(ctx, rq)
}

// Logout invalidates the user token.
func (h *Handler) Logout(ctx context.Context, rq *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	return h.biz.AuthV1().Logout(ctx, rq)
}

// RefreshToken generates a new token using the refresh token.
func (h *Handler) RefreshToken(ctx context.Context, rq *v1.RefreshTokenRequest) (*v1.LoginReply, error) {
	return h.biz.AuthV1().RefreshToken(ctx, rq)
}

// Authenticate validates the user token and returns the user ID.
func (h *Handler) Authenticate(ctx context.Context, rq *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	return h.biz.AuthV1().Authenticate(ctx, rq.Token)
}

// Authorize checks whether the user is authorized for the object/action.
func (h *Handler) Authorize(ctx context.Context, rq *v1.AuthorizeRequest) (*v1.AuthorizeResponse, error) {
	return h.biz.AuthV1().Authorize(ctx, rq.Sub, rq.Obj, rq.Act)
}

// Auth authenticates and authorizes the user token for an object/action.
func (h *Handler) Auth(ctx context.Context, rq *v1.AuthRequest) (*v1.AuthResponse, error) {
	authn, err := h.Authenticate(ctx, &v1.AuthenticateRequest{Token: rq.Token})
	if err != nil {
		return nil, err
	}

	authz, err := h.Authorize(ctx, &v1.AuthorizeRequest{Sub: authn.UserID, Obj: rq.Obj, Act: rq.Act})
	if err != nil {
		return nil, err
	}

	return &v1.AuthResponse{UserID: authn.UserID, Allowed: authz.Allowed}, nil
}
