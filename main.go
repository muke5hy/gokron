package main

import (
	"github.com/yourusername/cron-job-product/api"
	"github.com/yourusername/cron-job-product/cron"
	"github.com/yourusername/cron-job-product/db"
)

func main() {
	// Initialize the database
	db.Init()
	// Initialize the cron manager
	cron.Init()

	// Start the API
	router := api.Init()

	// Register the UI routes
	registerUIRoutes(router)

	// Run the server
	router.Run(":8080")
}
