package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/yourusername/cron-job-product/db/models"
	"github.com/yourusername/cron-job-product/repository"
)

var cronManager *cron.Cron

func Init() {
	cronManager = cron.New(cron.WithSeconds())
	cronManager.Start()

	syncCronJobs()
}

func syncCronJobs() {
	CronJobRepo, err := repository.NewCronJobRepoStruct()
	if err != nil {
		fmt.Println(err)
	}
	cronJobs, err := CronJobRepo.ListCronJobs()
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
