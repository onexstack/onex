package conversion

import (
	"github.com/onexstack/onexstack/pkg/core"

	"github.com/onexstack/onex/internal/nightwatch/model"
	v1 "github.com/onexstack/onex/pkg/api/nightwatch/v1"
)

// CronJobMToCronJobV1 converts a CronJobM object from the internal model
// to a CronJob object in the v1 API format.
func CronJobMToCronJobV1(cronJobModel *model.CronJobM) *v1.CronJob {
	var cronJob v1.CronJob
	_ = core.CopyWithConverters(&cronJob, cronJobModel)

	var job v1.Job
	core.Copy(&job, cronJobModel.JobTemplate)
	cronJob.JobTemplate = &job

	return &cronJob
}

// CronJobV1ToCronJobM converts a CronJob object from the v1 API format
// to a CronJobM object in the internal model.
func CronJobV1ToCronJobM(cronJob *v1.CronJob) *model.CronJobM {
	var cronJobModel model.CronJobM
	_ = core.CopyWithConverters(&cronJobModel, cronJob)
	return &cronJobModel
}
