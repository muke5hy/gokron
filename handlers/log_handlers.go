package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/cron-job-product/repository"
)

type LogHandler interface {
	ListLogs(c *gin.Context)
	CreateLog(c *gin.Context)
}

type LogHandelerImpl struct {
	repo repository.LogInterface
}

func NewLogHandelerImpl() (*LogHandelerImpl, error) {
	logRep, err := repository.NewlogRepo()
	if err != nil {
		return nil, err
	}
	return &LogHandelerImpl{
		repo: logRep,
	}, nil
}

func (l LogHandelerImpl) ListLogs(c *gin.Context) {
	cronJobID, err := strconv.Atoi(c.Param("cronJobID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Cron Job ID"})
		return
	}

	filter := make(map[string]interface{})

	logs, err := l.repo.ListLogs(uint(cronJobID), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (l LogHandelerImpl) CreateLog(c *gin.Context) {

}
