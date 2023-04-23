package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yourusername/cron-job-product/db/models"
)

var DB *gorm.DB

func Init() {
	var err error
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost", "5432", "username", "dbname", "password", "disable")

	DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		panic("failed to connect to database")
	}

	// Run migrations
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&models.CronJob{}, &models.Log{})
}
