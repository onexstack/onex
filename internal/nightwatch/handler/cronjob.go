package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/onexstack/onexstack/pkg/core"
)

// CreateCronJob handles the creation of a new cronjob.
func (h *Handler) CreateCronJob(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.CronJobV1().Create, h.val.ValidateCreateCronJobRequest)
}

// UpdateCronJob handles updating an existing cronjob's details.
func (h *Handler) UpdateCronJob(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.CronJobV1().Update, h.val.ValidateUpdateCronJobRequest)
}

// DeleteCronJob handles the deletion of one or more cronjobs.
func (h *Handler) DeleteCronJob(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.CronJobV1().Delete, h.val.ValidateDeleteCronJobRequest)
}

// GetCronJob retrieves information about a specific cronjob.
func (h *Handler) GetCronJob(c *gin.Context) {
	core.HandleUriRequest(c, h.biz.CronJobV1().Get, h.val.ValidateGetCronJobRequest)
}

// ListCronJob retrieves a list of cronjobs based on query parameters.
func (h *Handler) ListCronJob(c *gin.Context) {
	core.HandleQueryRequest(c, h.biz.CronJobV1().List, h.val.ValidateListCronJobRequest)
}
