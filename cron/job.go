package cron

import (
	"bytes"
	"log"
	"os/exec"
	"time"

	"github.com/yourusername/cron-job-product/db/models"
)

type CronJob struct {
	CronJobModel *models.CronJob
}

func NewCronJob(cronJobModel *models.CronJob) *CronJob {
	return &CronJob{
		CronJobModel: cronJobModel,
	}
}

func (j *CronJob) Run() {
	cmd := exec.Command("sh", "-c", j.CronJobModel.Command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	logEntry := models.Log{
		CronJobID: j.CronJobModel.ID,
		CreatedAt: time.Now(),
	}

	if err != nil {
		logEntry.Status = "Failed"
		logEntry.Output = err.Error()
	} else {
		logEntry.Status = "Success"
		logEntry.Output = out.String()
	}

	err = models.CreateLog(&logEntry)
	if err != nil {
		log.Printf("Error saving log: %v", err)
	}
}
