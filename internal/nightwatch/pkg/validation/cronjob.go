package validation

import (
	"context"

	genericvalidation "github.com/onexstack/onexstack/pkg/validation"

	v1 "github.com/onexstack/onex/pkg/api/nightwatch/v1"
)

// ValidateCronJobRules returns a set of validation rules for cronjob-related requests.
func (v *Validator) ValidateCronJobRules() genericvalidation.Rules {
	return genericvalidation.Rules{}
}

// ValidateCreateCronJobRequest validates the fields of a CreateCronJobRequest.
func (v *Validator) ValidateCreateCronJobRequest(ctx context.Context, rq *v1.CreateCronJobRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateCronJobRules())
}

// ValidateUpdateCronJobRequest validates the fields of an UpdateCronJobRequest.
func (v *Validator) ValidateUpdateCronJobRequest(ctx context.Context, rq *v1.UpdateCronJobRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateCronJobRules())
}

// ValidateDeleteCronJobRequest validates the fields of a DeleteCronJobRequest.
func (v *Validator) ValidateDeleteCronJobRequest(ctx context.Context, rq *v1.DeleteCronJobRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateCronJobRules())
}

// ValidateGetCronJobRequest validates the fields of a GetCronJobRequest.
func (v *Validator) ValidateGetCronJobRequest(ctx context.Context, rq *v1.GetCronJobRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateCronJobRules())
}

// ValidateListCronJobRequest validates the fields of a ListCronJobRequest, focusing on selected fields ("Offset" and "Limit").
func (v *Validator) ValidateListCronJobRequest(ctx context.Context, rq *v1.ListCronJobRequest) error {
	return genericvalidation.ValidateSelectedFields(rq, v.ValidateCronJobRules(), "Offset", "Limit")
}
