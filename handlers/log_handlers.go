package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/cron-job-product/db/models"
)

func ListLogs(c *gin.Context) {
	cronJobID, err := strconv.Atoi(c.Param("cronJobID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Cron Job ID"})
		return
	}

	logs, err := models.ListLogs(uint(cronJobID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
