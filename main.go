package main

import (
	"github.com/yourusername/cron-job-product/api"
	"github.com/yourusername/cron-job-product/cron"
	"github.com/yourusername/cron-job-product/db"
	"github.com/yourusername/cron-job-product/template"
)

func main() {
	// Initialize the database
	db.Init()
	// Initialize the cron manager
	cron.Init()

	// Start the API
	router := api.Init()

	// Register the UI routes
	template.RegisterUIRoutes(router)

	// Run the server
	router.Run(":8000")
}
