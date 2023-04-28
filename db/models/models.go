package models

import (
	"time"
)

type CronJob struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Command   string    `gorm:"type:text;not null" json:"command"`
	Schedule  string    `gorm:"not null" json:"schedule"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

type Log struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CronJobID uint      `gorm:"not null;index" json:"cron_job_id"`
	Output    string    `gorm:"type:text" json:"output"`
	Status    string    `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
