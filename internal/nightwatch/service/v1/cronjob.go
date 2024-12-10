package v1

import (
	"github.com/gin-gonic/gin"

	nwv1 "github.com/superproj/onex/pkg/api/nightwatch/v1"
	"github.com/superproj/onex/pkg/api/zerrors"
	"github.com/superproj/onex/pkg/core"
)

// CreateCronJob handles the creation of a new CronJob.
func (s *NightWatchService) CreateCronJob(c *gin.Context) {
	var rq nwv1.CreateCronJobRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.Respond(c, zerrors.ErrorInvalidParameter(err.Error()), nil)
		return
	}

	if err := s.valid.ValidateCreateCronJobRequest(c, &rq); err != nil {
		core.Respond(c, err, nil)
		return
	}

	resp, err := s.biz.CronJobs().Create(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, resp)
}

// UpdateCronJob handles the update of an existing CronJob.
func (s *NightWatchService) UpdateCronJob(c *gin.Context) {
	var rq nwv1.UpdateCronJobRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.Respond(c, zerrors.ErrorInvalidParameter(err.Error()), nil)
		return
	}
	rq.CronJobID = c.Param("cronJobID")

	resp, err := s.biz.CronJobs().Update(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, resp)
}

// DeleteCronJob handles the deletion of a specified CronJob.
func (s *NightWatchService) DeleteCronJob(c *gin.Context) {
	rq := nwv1.DeleteCronJobRequest{
		CronJobIDs: []string{c.Param("cronJobID")},
	}
	resp, err := s.biz.CronJobs().Delete(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, resp)
}

// GetCronJob retrieves a specified CronJob.
func (s *NightWatchService) GetCronJob(c *gin.Context) {
	rq := nwv1.GetCronJobRequest{
		CronJobID: c.Param("cronJobID"),
	}
	cronJob, err := s.biz.CronJobs().Get(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, cronJob)
}

// ListCronJob retrieves all CronJobs.
func (s *NightWatchService) ListCronJob(c *gin.Context) {
	var rq nwv1.ListCronJobRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.Respond(c, zerrors.ErrorInvalidParameter(err.Error()), nil)
		return
	}

	resp, err := s.biz.CronJobs().List(c, &rq)
	if err != nil {
		core.Respond(c, err, nil)
		return
	}

	core.Respond(c, nil, resp)
}
