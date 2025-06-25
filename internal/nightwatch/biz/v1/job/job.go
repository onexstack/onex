package job

//go:generate mockgen -destination mock_job.go -package job onex/internal/nightwatch/biz/v1/job JobBiz

import (
	"context"
	"sync"

	"github.com/onexstack/onexstack/pkg/core"
	"github.com/onexstack/onexstack/pkg/log"
	"github.com/onexstack/onexstack/pkg/ptr"
	"github.com/onexstack/onexstack/pkg/store/where"
	"golang.org/x/sync/errgroup"

	"github.com/onexstack/onex/internal/nightwatch/model"
	"github.com/onexstack/onex/internal/nightwatch/pkg/conversion"
	"github.com/onexstack/onex/internal/nightwatch/store"
	"github.com/onexstack/onex/internal/pkg/known"
	v1 "github.com/onexstack/onex/pkg/api/nightwatch/v1"
)

// JobBiz defines the interface that contains methods for handling job requests.
type JobBiz interface {
	// Create creates a new job based on the provided request parameters.
	Create(ctx context.Context, rq *v1.CreateJobRequest) (*v1.CreateJobResponse, error)

	// Update updates an existing job based on the provided request parameters.
	Update(ctx context.Context, rq *v1.UpdateJobRequest) (*v1.UpdateJobResponse, error)

	// Delete removes one or more jobs based on the provided request parameters.
	Delete(ctx context.Context, rq *v1.DeleteJobRequest) (*v1.DeleteJobResponse, error)

	// Get retrieves the details of a specific job based on the provided request parameters.
	Get(ctx context.Context, rq *v1.GetJobRequest) (*v1.GetJobResponse, error)

	// List retrieves a list of jobs and their total count based on the provided request parameters.
	List(ctx context.Context, rq *v1.ListJobRequest) (*v1.ListJobResponse, error)

	// JobExpansion defines additional methods for extended job operations, if needed.
	JobExpansion
}

// JobExpansion defines additional methods for job operations.
type JobExpansion interface{}

// jobBiz is the implementation of the JobBiz.
type jobBiz struct {
	store store.IStore
}

// Ensure that *jobBiz implements the JobBiz.
var _ JobBiz = (*jobBiz)(nil)

// New creates and returns a new instance of *jobBiz.
func New(store store.IStore) *jobBiz {
	return &jobBiz{store: store}
}

// Create implements the Create method of the JobBiz.
func (b *jobBiz) Create(ctx context.Context, rq *v1.CreateJobRequest) (*v1.CreateJobResponse, error) {
	var jobM model.JobM
	_ = core.Copy(&jobM, rq)
	// TODO: Retrieve the UserID from the custom context and assign it as needed.
	// jobM.UserID = contextx.UserID(ctx)

	if err := b.store.Job().Create(ctx, &jobM); err != nil {
		return nil, err
	}

	return &v1.CreateJobResponse{JobID: jobM.JobID}, nil
}

// Update implements the Update method of the JobBiz.
func (b *jobBiz) Update(ctx context.Context, rq *v1.UpdateJobRequest) (*v1.UpdateJobResponse, error) {
	whr := where.T(ctx).F("jobID", rq.GetJobID())
	jobM, err := b.store.Job().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	if rq.Name != nil {
		jobM.Name = *rq.Name
	}
	if rq.Description != nil {
		jobM.Description = *rq.Description
	}
	if rq.Params != nil {
		jobM.Params = ptr.To(model.JobParams(*rq.Params))
	}
	if rq.Results != nil {
		jobM.Results = ptr.To(model.JobResults(*rq.Results))
	}
	if rq.Status != nil {
		jobM.Status = *rq.Status
	}

	if err := b.store.Job().Update(ctx, jobM); err != nil {
		return nil, err
	}

	return &v1.UpdateJobResponse{}, nil
}

// Delete implements the Delete method of the JobBiz.
func (b *jobBiz) Delete(ctx context.Context, rq *v1.DeleteJobRequest) (*v1.DeleteJobResponse, error) {
	whr := where.T(ctx).F("jobID", rq.GetJobIDs())
	if err := b.store.Job().Delete(ctx, whr); err != nil {
		return nil, err
	}

	return &v1.DeleteJobResponse{}, nil
}

// Get implements the Get method of the JobBiz.
func (b *jobBiz) Get(ctx context.Context, rq *v1.GetJobRequest) (*v1.GetJobResponse, error) {
	whr := where.T(ctx).F("jobID", rq.GetJobID())
	jobM, err := b.store.Job().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	return &v1.GetJobResponse{Job: conversion.JobMToJobV1(jobM)}, nil
}

// List implements the List method of the JobBiz.
func (b *jobBiz) List(ctx context.Context, rq *v1.ListJobRequest) (*v1.ListJobResponse, error) {
	whr := where.T(ctx).P(int(rq.GetOffset()), int(rq.GetLimit()))
	count, jobList, err := b.store.Job().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)

	// Set the maximum concurrency limit using the constant MaxConcurrency
	eg.SetLimit(known.MaxErrGroupConcurrency)

	// Use goroutines to improve API performance
	for _, job := range jobList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				converted := conversion.JobMToJobV1(job)
				// TODO: Add additional processing logic and assign values to fields
				// that need updating, for example:
				// xxx := doSomething()
				// converted.XXX = xxx
				m.Store(job.ID, converted)

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.W(ctx).Errorw(err, "Failed to wait all function calls returned")
		return nil, err
	}

	jobs := make([]*v1.Job, 0, len(jobList))
	for _, item := range jobList {
		job, _ := m.Load(item.ID)
		jobs = append(jobs, job.(*v1.Job))
	}

	return &v1.ListJobResponse{Total: count, Jobs: jobs}, nil
}
