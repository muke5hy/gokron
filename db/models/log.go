package models

import (
	"time"

	"github.com/yourusername/cron-job-product/store"
)

type Log struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CronJobID uint      `gorm:"not null;index" json:"cron_job_id"`
	Output    string    `gorm:"type:text" json:"output"`
	Status    string    `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func CreateLog(log *Log) error {
	return store.DB.Create(log).Error
}

func ListLogs(cronJobID uint, filters map[string]interface{}) ([]Log, error) {
	var logs []Log

	query := store.DB.Where("cron_job_id = ?", cronJobID)

	// Apply filters if any
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key+" = ?", value)
		}
	}

	err := query.Order("created_at DESC").Find(&logs).Error
	return logs, err
}
