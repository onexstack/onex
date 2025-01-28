package validation

import (
	"context"

	"github.com/onexstack/onex/pkg/api/errno"
	genericvalidation "github.com/onexstack/onexstack/pkg/validation"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/onexstack/onex/internal/pkg/contextx"
	known "github.com/onexstack/onex/internal/pkg/known/nightwatch"
	v1 "github.com/onexstack/onex/pkg/api/nightwatch/v1"
)

var availableScope = sets.New(
	known.LLMJobScope,
)

// ValidateJobRules returns a set of validation rules for job-related requests.
func (v *Validator) ValidateJobRules() genericvalidation.Rules {
	return genericvalidation.Rules{}
}

// ValidateCreateJobRequest validates the fields of a CreateJobRequest.
func (v *Validator) ValidateCreateJobRequest(ctx context.Context, rq *v1.CreateJobRequest) error {
	if err := validateJob(ctx, rq.Job); err != nil {
		return err
	}

	rq.Job.UserID = contextx.UserID(ctx)

	return genericvalidation.ValidateAllFields(rq, v.ValidateJobRules())
}

// ValidateUpdateJobRequest validates the fields of an UpdateJobRequest.
func (v *Validator) ValidateUpdateJobRequest(ctx context.Context, rq *v1.UpdateJobRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateJobRules())
}

// ValidateDeleteJobRequest validates the fields of a DeleteJobRequest.
func (v *Validator) ValidateDeleteJobRequest(ctx context.Context, rq *v1.DeleteJobRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateJobRules())
}

// ValidateGetJobRequest validates the fields of a GetJobRequest.
func (v *Validator) ValidateGetJobRequest(ctx context.Context, rq *v1.GetJobRequest) error {
	return genericvalidation.ValidateAllFields(rq, v.ValidateJobRules())
}

// ValidateListJobRequest validates the fields of a ListJobRequest, focusing on selected fields ("Offset" and "Limit").
func (v *Validator) ValidateListJobRequest(ctx context.Context, rq *v1.ListJobRequest) error {
	return genericvalidation.ValidateSelectedFields(rq, v.ValidateJobRules(), "Offset", "Limit")
}

func validateJob(ctx context.Context, job *v1.Job) error {
	if job.Name == "" {
		return errno.ErrorInvalidParameter("job.name cannot be empty")
	}
	if job.Scope == "" {
		return errno.ErrorInvalidParameter("job.scope cannot be empty")
	}

	if !availableScope.Has(job.Scope) {
		return errno.ErrorInvalidParameter("invalid job.scope: %s", job.Scope)
	}

	if job.Watcher == "" {
		return errno.ErrorInvalidParameter("job.watcher cannot be empty")
	}

	job.UserID = contextx.UserID(ctx)

	return nil
}
