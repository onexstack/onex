package conversion

import (
	"github.com/onexstack/onexstack/pkg/core"

	"github.com/onexstack/onex/internal/nightwatch/model"
	v1 "github.com/onexstack/onex/pkg/api/nightwatch/v1"
)

// JobMToJobV1 converts a JobM object from the internal model
// to a Job object in the v1 API format.
func JobMToJobV1(jobModel *model.JobM) *v1.Job {
	var job v1.Job
	_ = core.CopyWithConverters(&job, jobModel)
	return &job
}

// JobV1ToJobM converts a Job object from the v1 API format
// to a JobM object in the internal model.
func JobV1ToJobM(job *v1.Job) *model.JobM {
	var jobModel model.JobM
	_ = core.CopyWithConverters(&jobModel, job)
	return &jobModel
}
