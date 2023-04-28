package repository

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yourusername/cron-job-product/db"
	"github.com/yourusername/cron-job-product/db/models"
)

type CronJobInterface interface {
	CreateCronJobRepo(cronJob *models.CronJob) error
	ListCronJobs() ([]models.CronJob, error)
	GetCronJob(id uint) (models.CronJob, error)
	UpdateCronJob(id uint, updatedCronJob *models.CronJob) error
	DeleteCronJob(id uint) error
}

type CronJobRepoStruct struct {
	db *gorm.DB
}

func NewCronJobRepoStruct() (*CronJobRepoStruct, error) {
	data := db.DB
	return &CronJobRepoStruct{
		db: data,
	}, nil
}

func (c CronJobRepoStruct) CreateCronJobRepo(cronJob *models.CronJob) error {
	cronJob.CreatedAt = time.Now()
	cronJob.UpdatedAt = time.Now()
	return c.db.Create(cronJob).Error
}

func (c CronJobRepoStruct) ListCronJobs() ([]models.CronJob, error) {
	var cronJobs []models.CronJob
	err := c.db.Find(&cronJobs).Error
	return cronJobs, err
}

func (c CronJobRepoStruct) GetCronJob(id uint) (models.CronJob, error) {
	var cronJob models.CronJob
	err := db.DB.First(&cronJob, id).Error
	if err != nil {
		return models.CronJob{}, errors.New("Cron job not found")
	}
	return cronJob, nil
}

func (c CronJobRepoStruct) UpdateCronJob(id uint, updatedCronJob *models.CronJob) error {
	cronJob, err := c.GetCronJob(id)
	if err != nil {
		return err
	}

	cronJob.Command = updatedCronJob.Command
	cronJob.Schedule = updatedCronJob.Schedule
	cronJob.UpdatedAt = time.Now()

	return db.DB.Save(&cronJob).Error
}

func (c CronJobRepoStruct) DeleteCronJob(id uint) error {
	cronJob, err := c.GetCronJob(id)
	if err != nil {
		return err
	}

	return db.DB.Delete(&cronJob).Error
}
