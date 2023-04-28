package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/cron-job-product/db/models"
	"github.com/yourusername/cron-job-product/repository"
)

type CronJobHandler interface {
	CreateCronJob(c *gin.Context)
	ListCronJobs(c *gin.Context)
	GetCronJob(c *gin.Context)
	UpdateCronJob(c *gin.Context)
	DeleteCronJob(c *gin.Context)
}

type CronJobHandelerImpl struct {
	repo repository.CronJobInterface
}

func NewCronJobHandelerImpl() (*CronJobHandelerImpl, error) {
	cronjobRep, err := repository.NewCronJobRepoStruct()
	if err != nil {
		return nil, err
	}
	return &CronJobHandelerImpl{
		repo: cronjobRep,
	}, nil
}

func (cj CronJobHandelerImpl) CreateCronJob(c *gin.Context) {
	var cronJob models.CronJob
	if err := c.ShouldBindJSON(&cronJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cj.repo.CreateCronJobRepo(&cronJob); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cronJob)
}

func (cj CronJobHandelerImpl) ListCronJobs(c *gin.Context) {
	cronJobs, err := cj.repo.ListCronJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cronJobs)
}

func (cj CronJobHandelerImpl) GetCronJob(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	cronJob, err := cj.repo.GetCronJob(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cronJob)
}

func (cj CronJobHandelerImpl) UpdateCronJob(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var cronJob models.CronJob
	if err := c.ShouldBindJSON(&cronJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cj.repo.UpdateCronJob(uint(id), &cronJob); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cronJob)
}

func (cj CronJobHandelerImpl) DeleteCronJob(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := cj.repo.DeleteCronJob(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
