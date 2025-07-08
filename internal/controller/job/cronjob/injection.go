/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cronjob

import (
	"context"
	"sync"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/onexstack/onex/pkg/apis/batch/v1beta1"
)

// cronJobControlInterface is an interface that knows how to update CronJob status
// created as an interface to allow testing.
type cronJobControlInterface interface {
	// GetCronJob retrieves a CronJob.
	GetCronJob(ctx context.Context, namespace, name string) (*v1beta1.CronJob, error)
	// UpdateStatus(ctx context.Context, cronJob *v1beta1.CronJob) (*v1beta1.CronJob, error)
}

// realCJControl is the default implementation of cronJobControlInterface.
type realCJControl struct {
	client client.Client
}

var _ cronJobControlInterface = &realCJControl{}

func (c *realCJControl) GetCronJob(ctx context.Context, namespace, name string) (*v1beta1.CronJob, error) {
	cronJob := new(v1beta1.CronJob)
	key := client.ObjectKey{Namespace: namespace, Name: name}

	if err := c.client.Get(ctx, key, cronJob); err != nil {
		return nil, err
	}
	return cronJob, nil
}

func (c *realCJControl) UpdateStatus(ctx context.Context, cronJob *v1beta1.CronJob) (*v1beta1.CronJob, error) {
	return cronJob, c.client.Status().Update(ctx, cronJob)
}

// fakeCJControl is the default implementation of cronJobControlInterface.
type fakeCJControl struct {
	CronJob *v1beta1.CronJob
	Updates []v1beta1.CronJob
}

var _ cronJobControlInterface = &fakeCJControl{}

func (c *fakeCJControl) GetCronJob(ctx context.Context, namespace, name string) (*v1beta1.CronJob, error) {
	if name == c.CronJob.Name && namespace == c.CronJob.Namespace {
		return c.CronJob, nil
	}
	return nil, errors.NewNotFound(schema.GroupResource{
		Group:    "v1beta1",
		Resource: "cronjobs",
	}, name)
}

func (c *fakeCJControl) UpdateStatus(ctx context.Context, cronJob *v1beta1.CronJob) (*v1beta1.CronJob, error) {
	c.Updates = append(c.Updates, *cronJob)
	return cronJob, nil
}

// ------------------------------------------------------------------ //

// jobControlInterface is an interface that knows how to add or delete jobs
// created as an interface to allow testing.
type jobControlInterface interface {
	// GetJob retrieves a Job.
	GetJob(ctx context.Context, namespace, name string) (*v1beta1.Job, error)
	// CreateJob creates new Jobs according to the spec.
	CreateJob(ctx context.Context, job *v1beta1.Job) (*v1beta1.Job, error)
	// DeleteJob deletes the Job identified by name.
	// TODO: delete by UID?
	DeleteJob(ctx context.Context, job *v1beta1.Job) error
}

// realJobControl is the default implementation of jobControlInterface.
type realJobControl struct {
	client   client.Client
	Recorder record.EventRecorder
}

var _ jobControlInterface = &realJobControl{}

func (r realJobControl) GetJob(ctx context.Context, namespace, name string) (*v1beta1.Job, error) {
	job := new(v1beta1.Job)
	key := client.ObjectKey{Namespace: namespace, Name: name}

	if err := r.client.Get(ctx, key, job); err != nil {
		return nil, err
	}

	return job, nil
}

func (r realJobControl) CreateJob(ctx context.Context, job *v1beta1.Job) (*v1beta1.Job, error) {
	if err := r.client.Create(ctx, job); err != nil {
		return job, err
	}
	return job, nil
}

func (r realJobControl) DeleteJob(ctx context.Context, job *v1beta1.Job) error {
	return r.client.Delete(ctx, job, client.PropagationPolicy(metav1.DeletePropagationBackground))
}

type fakeJobControl struct {
	sync.Mutex
	Job           *v1beta1.Job
	Jobs          []v1beta1.Job
	DeleteJobName []string
	Err           error
	CreateErr     error
	UpdateJobName []string
	PatchJobName  []string
	Patches       [][]byte
}

var _ jobControlInterface = &fakeJobControl{}

func (f *fakeJobControl) CreateJob(ctx context.Context, job *v1beta1.Job) (*v1beta1.Job, error) {
	f.Lock()
	defer f.Unlock()
	if f.CreateErr != nil {
		return nil, f.CreateErr
	}
	f.Jobs = append(f.Jobs, *job)
	job.UID = "test-uid"
	return job, nil
}

func (f *fakeJobControl) GetJob(ctx context.Context, namespace, name string) (*v1beta1.Job, error) {
	f.Lock()
	defer f.Unlock()
	if f.Err != nil {
		return nil, f.Err
	}
	return f.Job, nil
}

func (f *fakeJobControl) DeleteJob(ctx context.Context, job *v1beta1.Job) error {
	f.Lock()
	defer f.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.DeleteJobName = append(f.DeleteJobName, job.Name)
	return nil
}
