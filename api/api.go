package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/cron-job-product/handlers"
)

func Init() *gin.Engine {
	router := gin.Default()

	CronJobHandler, err := handlers.NewCronJobHandelerImpl()
	if err != nil {
		log.Fatal(err)
	}

	api := router.Group("/api")
	{
		api.POST("/", CronJobHandler.CreateCronJob)
		api.GET("/", CronJobHandler.ListCronJobs)
		api.GET("/:id", CronJobHandler.GetCronJob)
		api.PUT("/:id", CronJobHandler.UpdateCronJob)
		api.DELETE("/:id", CronJobHandler.DeleteCronJob)
	}

	logHandler, err := handlers.NewLogHandelerImpl()
	if err != nil {
		log.Fatal(err)
	}
	logs := router.Group("/api/logs")
	{
		logs.GET("/:cronJobID", logHandler.ListLogs)
	}

	return router
}
