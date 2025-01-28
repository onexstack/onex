package validation

import (
	"context"

	genericvalidation "github.com/onexstack/onexstack/pkg/validation"

	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
)

// ValidateMinerRules returns a set of validation rules for miner-related requests.
func (v *Validator) ValidateMinerRules() genericvalidation.Rules {
	return genericvalidation.Rules{}
}

// ValidateCreateMinerRequest validates the fields of a CreateMinerRequest.
func (v *Validator) ValidateCreateMinerRequest(ctx context.Context, rq *v1.CreateMinerRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateMinerRules())
}

// ValidateUpdateMinerRequest validates the fields of an UpdateMinerRequest.
func (v *Validator) ValidateUpdateMinerRequest(ctx context.Context, rq *v1.UpdateMinerRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateMinerRules())
}

// ValidateDeleteMinerRequest validates the fields of a DeleteMinerRequest.
func (v *Validator) ValidateDeleteMinerRequest(ctx context.Context, rq *v1.DeleteMinerRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateMinerRules())
}

// ValidateGetMinerRequest validates the fields of a GetMinerRequest.
func (v *Validator) ValidateGetMinerRequest(ctx context.Context, rq *v1.GetMinerRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateMinerRules())
}

// ValidateListMinerRequest validates the fields of a ListMinerRequest, focusing on selected fields ("Offset" and "Limit").
func (v *Validator) ValidateListMinerRequest(ctx context.Context, rq *v1.ListMinerRequest) error {
	return genericvalidation.ValidateSelectedFields(rq, v.ValidateMinerRules(), "Offset", "Limit")
}
