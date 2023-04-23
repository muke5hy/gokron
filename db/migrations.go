package db

import "github.com/yourusername/cron-job-product/db/models"

func Migrate() {
	DB.AutoMigrate(&models.CronJob{}, &models.Log{})
}
