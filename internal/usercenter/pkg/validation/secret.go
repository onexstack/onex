package validation

import (
	"context"

	"github.com/onexstack/onexstack/pkg/store/where"
	genericvalidation "github.com/onexstack/onexstack/pkg/validation"

	ucknown "github.com/onexstack/onex/internal/pkg/known/usercenter"
	"github.com/onexstack/onex/pkg/api/errno"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// ValidateSecretRules returns a set of validation rules for secret-related requests.
func (v *Validator) ValidateSecretRules() genericvalidation.Rules {
	return genericvalidation.Rules{}
}

// ValidateCreateSecretRequest validates the fields of a CreateSecretRequest.
func (v *Validator) ValidateCreateSecretRequest(ctx context.Context, rq *v1.CreateSecretRequest) error {
	_, secrets, err := v.store.Secret().List(ctx, where.T(ctx))
	if err != nil {
		return err
	}

	if len(secrets) >= ucknown.MaxSecretCount {
		return errno.ErrorInvalidParameter("secret reach the max count %d", ucknown.MaxSecretCount)
	}

	return nil
}

// ValidateUpdateSecretRequest validates the fields of an UpdateSecretRequest.
func (v *Validator) ValidateUpdateSecretRequest(ctx context.Context, rq *v1.UpdateSecretRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateSecretRules())
}

// ValidateDeleteSecretRequest validates the fields of a DeleteSecretRequest.
func (v *Validator) ValidateDeleteSecretRequest(ctx context.Context, rq *v1.DeleteSecretRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateSecretRules())
}

// ValidateGetSecretRequest validates the fields of a GetSecretRequest.
func (v *Validator) ValidateGetSecretRequest(ctx context.Context, rq *v1.GetSecretRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateSecretRules())
}

// ValidateListSecretRequest validates the fields of a ListSecretRequest, focusing on selected fields ("Offset" and "Limit").
func (v *Validator) ValidateListSecretRequest(ctx context.Context, rq *v1.ListSecretRequest) error {
	return genericvalidation.ValidateSelectedFields(rq, v.ValidateSecretRules(), "Offset", "Limit")
}
