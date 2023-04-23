package models

import (
	"errors"
	"time"

	"github.com/yourusername/cron-job-product/db"
)

type CronJob struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Command   string    `gorm:"type:text;not null" json:"command"`
	Schedule  string    `gorm:"not null" json:"schedule"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func CreateCronJob(cronJob *CronJob) error {
	cronJob.CreatedAt = time.Now()
	cronJob.UpdatedAt = time.Now()
	return db.DB.Create(cronJob).Error
}

func ListCronJobs() ([]CronJob, error) {
	var cronJobs []CronJob
	err := db.DB.Find(&cronJobs).Error
	return cronJobs, err
}

func GetCronJob(id uint) (CronJob, error) {
	var cronJob CronJob
	err := db.DB.First(&cronJob, id).Error
	if err != nil {
		return CronJob{}, errors.New("Cron job not found")
	}
	return cronJob, nil
}

func UpdateCronJob(id uint, updatedCronJob *CronJob) error {
	cronJob, err := GetCronJob(id)
	if err != nil {
		return err
	}

	cronJob.Command = updatedCronJob.Command
	cronJob.Schedule = updatedCronJob.Schedule
	cronJob.UpdatedAt = time.Now()

	return db.DB.Save(&cronJob).Error
}

func DeleteCronJob(id uint) error {
	cronJob, err := GetCronJob(id)
	if err != nil {
		return err
	}

	return db.DB.Delete(&cronJob).Error
}
