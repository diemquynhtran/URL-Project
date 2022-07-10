package controllers

import (
	"fmt"
	"learning-go/models"
	"net/http"
	"learning-go/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
type UrlRepo struct {
	Db *gorm.DB
}

func New() *UrlRepo {
	db := config.InitDb()
	db.AutoMigrate(&models.Url{})
	return &UrlRepo{Db: db}
}


func (repository *UrlRepo) GetURL(c *gin.Context)  {
	var url []models.Url
	err := models.GetAllUrl(repository.Db, &url)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, url)
	}	
}

func (repository *UrlRepo) CreateUrl(c *gin.Context) {
	var url models.Url
	c.BindJSON(&url)
	err := models.CreateUrl(repository.Db, &url)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, url)
	}
}


