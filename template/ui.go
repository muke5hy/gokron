package template

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/cron-job-product/repository"
)

func RegisterUIRoutes(router *gin.Engine) {

	CronJobRepo, err := repository.NewCronJobRepoStruct()
	if err != nil {
		fmt.Println(err)
	}
	// Serve the HTML template
	router.GET("/", func(c *gin.Context) {
		cronJobs, _ := CronJobRepo.ListCronJobs()
		tmplPath := filepath.Join("html", "index.html")
		tmpl, _ := template.ParseFiles(tmplPath)
		tmpl.Execute(c.Writer, gin.H{"CronJobs": cronJobs})
	})
}
