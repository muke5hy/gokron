package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/cron-job-product/handlers"
)

func Init() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/", handlers.CreateCronJob)
		api.GET("/", handlers.ListCronJobs)
		api.GET("/:id", handlers.GetCronJob)
		api.PUT("/:id", handlers.UpdateCronJob)
		api.DELETE("/:id", handlers.DeleteCronJob)
	}

	logs := router.Group("/api/logs")
	{
		logs.GET("/:cronJobID", handlers.ListLogs)
	}

	return router
}
