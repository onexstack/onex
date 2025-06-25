package validation

import (
	"context"

	genericvalidation "github.com/onexstack/onexstack/pkg/validation"

	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
)

// ValidateMinerSetRules returns a set of validation rules for minerset-related requests.
func (v *Validator) ValidateMinerSetRules() genericvalidation.Rules {
	return genericvalidation.Rules{}
}

// ValidateCreateMinerSetRequest validates the fields of a CreateMinerSetRequest.
func (v *Validator) ValidateCreateMinerSetRequest(ctx context.Context, rq *v1.CreateMinerSetRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateMinerSetRules())
}

// ValidateUpdateMinerSetRequest validates the fields of an UpdateMinerSetRequest.
func (v *Validator) ValidateUpdateMinerSetRequest(ctx context.Context, rq *v1.UpdateMinerSetRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateMinerSetRules())
}

// ValidateDeleteMinerSetRequest validates the fields of a DeleteMinerSetRequest.
func (v *Validator) ValidateDeleteMinerSetRequest(ctx context.Context, rq *v1.DeleteMinerSetRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateMinerSetRules())
}

// ValidateGetMinerSetRequest validates the fields of a GetMinerSetRequest.
func (v *Validator) ValidateGetMinerSetRequest(ctx context.Context, rq *v1.GetMinerSetRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateMinerSetRules())
}

// ValidateListMinerSetRequest validates the fields of a ListMinerSetRequest, focusing on selected fields ("Offset" and "Limit").
func (v *Validator) ValidateListMinerSetRequest(ctx context.Context, rq *v1.ListMinerSetRequest) error {
	return genericvalidation.ValidateSelectedFields(rq, v.ValidateMinerSetRules(), "Offset", "Limit")
}
