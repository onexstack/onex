package validation

import (
	"context"

	"github.com/onexstack/onexstack/pkg/i18n"
	"github.com/onexstack/onexstack/pkg/store/where"
	genericvalidation "github.com/onexstack/onexstack/pkg/validation"

	"github.com/onexstack/onex/internal/pkg/contextx"
	"github.com/onexstack/onex/internal/pkg/known"
	"github.com/onexstack/onex/internal/usercenter/pkg/locales"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// ValidateUserRules returns a set of validation rules for user-related requests.
func (v *Validator) ValidateUserRules() genericvalidation.Rules {
	return genericvalidation.Rules{}
}

// ValidateAuthRequest validates the authentication rquest.
// In this sample, no actual validation is needed, so it returns nil directly.
func (v *Validator) ValidateAuthRequest(ctx context.Context, rq *v1.AuthRequest) error {
	return nil
}

// ValidateAuthorizeRequest validates the authorization rquest.
// In this sample, no actual validation is needed, so it returns nil directly.
func (v *Validator) ValidateAuthorizeRequest(ctx context.Context, rq *v1.AuthorizeRequest) error {
	return nil
}

// ValidateCreateUserRequest validates the fields of a CreateUserRequest.
func (v *Validator) ValidateCreateUserRequest(ctx context.Context, rq *v1.CreateUserRequest) error {
	if _, err := v.store.User().Get(ctx, where.F("username", rq.Username)); err == nil {
		return i18n.FromContext(ctx).E(locales.UserAlreadyExists)
	}
	return nil
}

// ValidateUpdateUserRequest validates the fields of an UpdateUserRequest.
func (v *Validator) ValidateUpdateUserRequest(ctx context.Context, rq *v1.UpdateUserRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateUserRules())
}

// ValidateDeleteUserRequest validates the fields of a DeleteUserRequest.
func (v *Validator) ValidateDeleteUserRequest(ctx context.Context, rq *v1.DeleteUserRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateUserRules())
}

// ValidateGetUserRequest validates the fields of a GetUserRequest.
func (v *Validator) ValidateGetUserRequest(ctx context.Context, rq *v1.GetUserRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateUserRules())
}

// ValidateListUserRequest validates the fields of a ListUserRequest, focusing on selected fields ("Offset" and "Limit").
func (v *Validator) ValidateListUserRequest(ctx context.Context, rq *v1.ListUserRequest) error {
	if userID := contextx.UserID(ctx); userID != known.AdminUserID {
		return i18n.FromContext(ctx).E(locales.UserListUnauthorized)
	}
	return nil
}
