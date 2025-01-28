package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/onexstack/onexstack/pkg/core"
)

// CreateJob handles the creation of a new job.
func (h *Handler) CreateJob(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.JobV1().Create, h.val.ValidateCreateJobRequest)
}

// UpdateJob handles updating an existing job's details.
func (h *Handler) UpdateJob(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.JobV1().Update, h.val.ValidateUpdateJobRequest)
}

// DeleteJob handles the deletion of one or more jobs.
func (h *Handler) DeleteJob(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.JobV1().Delete, h.val.ValidateDeleteJobRequest)
}

// GetJob retrieves information about a specific job.
func (h *Handler) GetJob(c *gin.Context) {
	core.HandleUriRequest(c, h.biz.JobV1().Get, h.val.ValidateGetJobRequest)
}

// ListJob retrieves a list of jobs based on query parameters.
func (h *Handler) ListJob(c *gin.Context) {
	core.HandleQueryRequest(c, h.biz.JobV1().List, h.val.ValidateListJobRequest)
}
