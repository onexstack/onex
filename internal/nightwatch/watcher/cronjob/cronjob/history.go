// Package cronjob is a watcher implement.
package cronjob

import (
	"context"

	"github.com/onexstack/onexstack/pkg/log"
	"github.com/onexstack/onexstack/pkg/store/where"
	"github.com/onexstack/onexstack/pkg/watch/registry"

	"github.com/onexstack/onex/internal/nightwatch/model"
	"github.com/onexstack/onex/internal/nightwatch/store"
	known "github.com/onexstack/onex/internal/pkg/known/nightwatch"
)

var _ registry.Watcher = (*Watcher)(nil)

// watcher implement.
type History struct {
	store store.IStore
}

// Run runs the watcher.
func (h *History) Run() {
	ctx := context.Background()
	_, cronjobs, err := h.store.CronJob().List(ctx, where.F("suspend", known.JobNonSuspended))
	if err != nil {
		return
	}

	for _, cronjob := range cronjobs {
		h.retainRecords(ctx, known.JobSucceeded, cronjob.SuccessHistoryLimit)
		h.retainRecords(ctx, known.JobFailed, cronjob.FailedHistoryLimit)
	}
}

// Spec is parsed using the time zone of task Cron instance as the default.
func (h *History) Spec() string {
	return "@every 1s"
}

// SetStore sets the persistence store for the Watcher.
func (h *History) SetStore(store store.IStore) {
	h.store = store
}

func (h *History) retainRecords(ctx context.Context, status string, maxRecords int32) {
	_, jobs, err := h.store.Job().List(ctx, where.F("status", status))
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to list jobs")
		return
	}
	removedIDs := retainMaxElements(jobs, maxRecords)
	if err := h.store.Job().Delete(ctx, where.F("job_id", removedIDs)); err != nil {
		log.W(ctx).Errorw(err, "Failed to delete jobs")
	}
}

func retainMaxElements(jobs []*model.JobM, maxRecords int32) []string {
	all := make([]string, len(jobs))
	for i, job := range jobs {
		all[i] = job.JobID
	}

	if len(all) <= int(maxRecords) {
		return []string{}
	}

	return all[maxRecords:]
}

func init() {
	registry.Register("history", &Watcher{})
}
