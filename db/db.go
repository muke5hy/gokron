package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/yourusername/cron-job-product/db/models"
)

var DB *gorm.DB

func Init() {
	// Load environment variables from .env.local or .env file
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			panic("failed to load environment variables")
		}
	}

	var dbErr error
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println(dbURL)
	DB, dbErr = gorm.Open("postgres", dbURL)
	if dbErr != nil {
		fmt.Println(dbErr)
		panic("failed to connect to database")
	}

	// Run migrations
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&models.CronJob{}, &models.Log{})
}
