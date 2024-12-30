package model

import (
	"gorm.io/gorm"

	"github.com/onexstack/onex/internal/pkg/zid"
)

// AfterCreate runs after creating a CronJobM database record and updates the JobID field.
func (m *CronJobM) AfterCreate(tx *gorm.DB) (err error) {
	m.CronJobID = zid.CronJob.New(uint64(m.ID)) // Generate and set a new cronjob ID.

	return tx.Save(m).Error // Save the updated cronjob record to the database.
}

// AfterCreate runs after creating a JobM database record and updates the JobID field.
func (m *JobM) AfterCreate(tx *gorm.DB) (err error) {
	m.JobID = zid.Job.New(uint64(m.ID)) // Generate and set a new job ID.

	return tx.Save(m).Error // Save the updated job record to the database.
}
