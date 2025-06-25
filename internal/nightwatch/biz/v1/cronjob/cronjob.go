package cronjob

//go:generate mockgen -destination mock_cronjob.go -package cronjob onex/internal/nightwatch/biz/v1/cronjob CronJobBiz

import (
	"context"
	"sync"

	"github.com/onexstack/onexstack/pkg/core"
	"github.com/onexstack/onexstack/pkg/log"
	"github.com/onexstack/onexstack/pkg/store/where"
	"golang.org/x/sync/errgroup"

	"github.com/onexstack/onex/internal/nightwatch/model"
	"github.com/onexstack/onex/internal/nightwatch/pkg/conversion"
	"github.com/onexstack/onex/internal/nightwatch/store"
	"github.com/onexstack/onex/internal/pkg/known"
	v1 "github.com/onexstack/onex/pkg/api/nightwatch/v1"
)

// CronJobBiz defines the interface that contains methods for handling cronjob requests.
type CronJobBiz interface {
	// Create creates a new cronjob based on the provided request parameters.
	Create(ctx context.Context, rq *v1.CreateCronJobRequest) (*v1.CreateCronJobResponse, error)

	// Update updates an existing cronjob based on the provided request parameters.
	Update(ctx context.Context, rq *v1.UpdateCronJobRequest) (*v1.UpdateCronJobResponse, error)

	// Delete removes one or more cronjobs based on the provided request parameters.
	Delete(ctx context.Context, rq *v1.DeleteCronJobRequest) (*v1.DeleteCronJobResponse, error)

	// Get retrieves the details of a specific cronjob based on the provided request parameters.
	Get(ctx context.Context, rq *v1.GetCronJobRequest) (*v1.GetCronJobResponse, error)

	// List retrieves a list of cronjobs and their total count based on the provided request parameters.
	List(ctx context.Context, rq *v1.ListCronJobRequest) (*v1.ListCronJobResponse, error)

	// CronJobExpansion defines additional methods for extended cronjob operations, if needed.
	CronJobExpansion
}

// CronJobExpansion defines additional methods for cronjob operations.
type CronJobExpansion interface{}

// cronJobBiz is the implementation of the CronJobBiz.
type cronJobBiz struct {
	store store.IStore
}

// Ensure that *cronJobBiz implements the CronJobBiz.
var _ CronJobBiz = (*cronJobBiz)(nil)

// New creates and returns a new instance of *cronJobBiz.
func New(store store.IStore) *cronJobBiz {
	return &cronJobBiz{store: store}
}

// Create implements the Create method of the CronJobBiz.
func (b *cronJobBiz) Create(ctx context.Context, rq *v1.CreateCronJobRequest) (*v1.CreateCronJobResponse, error) {
	var cronJobM model.CronJobM
	_ = core.Copy(&cronJobM, rq)
	// TODO: Retrieve the UserID from the custom context and assign it as needed.
	// cronJobM.UserID = contextx.UserID(ctx)

	if err := b.store.CronJob().Create(ctx, &cronJobM); err != nil {
		return nil, err
	}

	return &v1.CreateCronJobResponse{CronJobID: cronJobM.CronJobID}, nil
}

// Update implements the Update method of the CronJobBiz.
func (b *cronJobBiz) Update(ctx context.Context, rq *v1.UpdateCronJobRequest) (*v1.UpdateCronJobResponse, error) {
	whr := where.T(ctx).F("cronJobID", rq.GetCronJobID())
	cronJobM, err := b.store.CronJob().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	if rq.Name != nil {
		cronJobM.Name = *rq.Name
	}
	if rq.Description != nil {
		cronJobM.Description = *rq.Description
	}
	if rq.Schedule != nil {
		cronJobM.Schedule = *rq.Schedule
	}
	if rq.ConcurrencyPolicy != nil {
		cronJobM.ConcurrencyPolicy = int32(*rq.ConcurrencyPolicy)
	}
	if rq.Suspend != nil {
		cronJobM.Suspend = *rq.Suspend
	}
	if rq.SuccessHistoryLimit != nil {
		cronJobM.SuccessHistoryLimit = *rq.SuccessHistoryLimit
	}
	if rq.FailedHistoryLimit != nil {
		cronJobM.FailedHistoryLimit = *rq.FailedHistoryLimit
	}
	if err := b.store.CronJob().Update(ctx, cronJobM); err != nil {
		return nil, err
	}

	return &v1.UpdateCronJobResponse{}, nil
}

// Delete implements the Delete method of the CronJobBiz.
func (b *cronJobBiz) Delete(ctx context.Context, rq *v1.DeleteCronJobRequest) (*v1.DeleteCronJobResponse, error) {
	whr := where.T(ctx).F("cronJobID", rq.GetCronJobIDs())
	if err := b.store.CronJob().Delete(ctx, whr); err != nil {
		return nil, err
	}

	return &v1.DeleteCronJobResponse{}, nil
}

// Get implements the Get method of the CronJobBiz.
func (b *cronJobBiz) Get(ctx context.Context, rq *v1.GetCronJobRequest) (*v1.GetCronJobResponse, error) {
	whr := where.T(ctx).F("cronJobID", rq.GetCronJobID())
	cronJobM, err := b.store.CronJob().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	return &v1.GetCronJobResponse{CronJob: conversion.CronJobMToCronJobV1(cronJobM)}, nil
}

// List implements the List method of the CronJobBiz.
func (b *cronJobBiz) List(ctx context.Context, rq *v1.ListCronJobRequest) (*v1.ListCronJobResponse, error) {
	whr := where.T(ctx).P(int(rq.GetOffset()), int(rq.GetLimit()))
	count, cronJobList, err := b.store.CronJob().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)

	// Set the maximum concurrency limit using the constant MaxConcurrency
	eg.SetLimit(known.MaxErrGroupConcurrency)

	// Use goroutines to improve API performance
	for _, cronJob := range cronJobList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				converted := conversion.CronJobMToCronJobV1(cronJob)
				// TODO: Add additional processing logic and assign values to fields
				// that need updating, for example:
				// xxx := doSomething()
				// converted.XXX = xxx
				m.Store(cronJob.ID, converted)

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.W(ctx).Errorw(err, "Failed to wait all function calls returned")
		return nil, err
	}

	cronJobs := make([]*v1.CronJob, 0, len(cronJobList))
	for _, item := range cronJobList {
		cronJob, _ := m.Load(item.ID)
		cronJobs = append(cronJobs, cronJob.(*v1.CronJob))
	}

	return &v1.ListCronJobResponse{Total: count, CronJobs: cronJobs}, nil
}
