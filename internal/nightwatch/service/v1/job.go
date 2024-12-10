package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/superproj/onex/pkg/core"
	nwv1 "github.com/superproj/onex/pkg/api/nightwatch/v1"
	"github.com/superproj/onex/pkg/api/zerrors"
)

// CreateJob handles the creation of a new Job.
func (s *NightWatchService) CreateJob(c *gin.Context) {
	var rq nwv1.CreateJobRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.Respond(c, zerrors.ErrorBindFailed(err.Error()), nil)
		return
	}

	if err := s.valid.ValidateCreateJobRequest(c, &rq); err != nil {
		core.Respond(c, zerrors.ErrorInvalidParameter(err.Error()), nil)
		return
	}

	resp, err := s.biz.Jobs().Create(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, resp)
}

// UpdateJob handles the update of an existing Job.
func (s *NightWatchService) UpdateJob(c *gin.Context) {
	var rq nwv1.UpdateJobRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.Respond(c, zerrors.ErrorInvalidParameter(err.Error()), nil)
		return
	}
	rq.JobID = c.Param("jobID")

	resp, err := s.biz.Jobs().Update(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, resp)
}

// DeleteJob handles the deletion of a specified Job.
func (s *NightWatchService) DeleteJob(c *gin.Context) {
	rq := nwv1.DeleteJobRequest{
		JobIDs: []string{c.Param("jobID")},
	}
	resp, err := s.biz.Jobs().Delete(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, resp)
}

// GetJob retrieves a specified Job.
func (s *NightWatchService) GetJob(c *gin.Context) {
	rq := nwv1.GetJobRequest{
		JobID: c.Param("jobID"),
	}
	job, err := s.biz.Jobs().Get(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, job)
}

// ListJob retrieves all Jobs.
func (s *NightWatchService) ListJob(c *gin.Context) {
	var rq nwv1.ListJobRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.Respond(c, zerrors.ErrorInvalidParameter(err.Error()), nil)
		return
	}

	resp, err := s.biz.Jobs().List(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, resp)
}
