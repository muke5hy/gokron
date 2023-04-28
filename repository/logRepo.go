package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yourusername/cron-job-product/db"
	"github.com/yourusername/cron-job-product/db/models"
)

type LogInterface interface {
	CreateLog(log *models.Log) error
	ListLogs(cronJobID uint, filters map[string]interface{}) ([]models.Log, error)
}

type logRepoImpl struct {
	db *gorm.DB
}

func NewlogRepo() (*logRepoImpl, error) {
	data := db.DB
	return &logRepoImpl{
		db: data,
	}, nil
}

func (l logRepoImpl) CreateLog(log *models.Log) error {
	return db.DB.Create(log).Error
}

func (l logRepoImpl) ListLogs(cronJobID uint, filters map[string]interface{}) ([]models.Log, error) {
	var logs []models.Log

	query := l.db.Where("cron_job_id = ?", cronJobID)

	// Apply filters if any
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key+" = ?", value)
		}
	}

	err := query.Order("created_at DESC").Find(&logs).Error
	return logs, err
}
