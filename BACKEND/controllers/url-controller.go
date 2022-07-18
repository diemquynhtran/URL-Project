package controllers

import (
	"fmt"
	"learning-go/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// /:short
func (repository *UrlRepo) GetURL(c *gin.Context)  {
	var url models.Url
	url = models.GetUrl(repository.Db, &url, c.Param("short"))
	c.Redirect(302, url.LongURL)
}
// /url
func (repository *UrlRepo) CreateUrl(c *gin.Context) {
	var url models.Url
	err := models.CreateUrl(repository.Db, &url, c.Query("long"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, url)
	}
}


