package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/yourusername/cron-job-product/db/models"
)

var cronManager *cron.Cron

func Init() {
	cronManager = cron.New(cron.WithSeconds())
	cronManager.Start()

	syncCronJobs()
}

func syncCronJobs() {
	cronJobs, err := models.ListCronJobs()
	if err != nil {
		return
	}

	for _, cronJob := range cronJobs {
		AddCronJob(&cronJob)
	}
}

func AddCronJob(cronJob *models.CronJob) (cron.EntryID, error) {
	job := NewCronJob(cronJob)
	return cronManager.AddJob(cronJob.Schedule, job)
}

func RemoveCronJob(entryID cron.EntryID) {
	cronManager.Remove(entryID)
}
