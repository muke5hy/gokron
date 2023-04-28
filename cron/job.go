package cron

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/yourusername/cron-job-product/db/models"
	"github.com/yourusername/cron-job-product/repository"
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
	logRepo, err := repository.NewlogRepo()
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command("sh", "-c", j.CronJobModel.Command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()

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

	err = logRepo.CreateLog(&logEntry)
	if err != nil {
		log.Printf("Error saving log: %v", err)
	}
}
