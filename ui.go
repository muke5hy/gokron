package main

import (
	"html/template"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/cron-job-product/db/models"
)

func registerUIRoutes(router *gin.Engine) {
	// Serve the HTML template
	router.GET("/", func(c *gin.Context) {
		cronJobs, _ := models.ListCronJobs()
		tmplPath := filepath.Join("html", "index.html")
		tmpl, _ := template.ParseFiles(tmplPath)
		tmpl.Execute(c.Writer, gin.H{"CronJobs": cronJobs})
	})
}
