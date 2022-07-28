package controller

import (
	"learning-go/dto"
	"learning-go/helper"
	"learning-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlController interface {
	GetUrlByName(ctx *gin.Context)
	CreateUrl(ctx *gin.Context)
	GetUrlByUser(ctx *gin.Context)
}

type urlController struct {
	urlService service.UrlService
}

func NewUrlController(urlService service.UrlService) UrlController {
	return &urlController{
		urlService: urlService,
	}
}

func (c *urlController) CreateUrl(ctx *gin.Context) {
	var urlDTO dto.CreateUrl
	errDTO := ctx.ShouldBind(&urlDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result := c.urlService.CreateUrl(urlDTO)
		response := helper.BuildResponse(true, "Create ok", result)
		ctx.JSON(http.StatusCreated, response)
	}
}
func (c *urlController) GetUrlByName(ctx *gin.Context) {
	var urlDTO dto.GetUrlByName
	errDTO := ctx.ShouldBind(&urlDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		id := helper.IdBase10(ctx.Param("short"))
		log.Println(id)
		result := c.urlService.GetUrlById(id)
		ctx.Redirect(302, result.LongURL)
	}
}

func (c *urlController) GetUrlByUser(ctx *gin.Context) {

}

/*
// /:short
func (repository *UrlRepo) GetURL(c *gin.Context) {
	var url entity.Url
	url = entity.GetUrl(repository.Db, &url, c.Param("short"))
	c.Redirect(302, url.LongURL)
}

// /url
func (repository *UrlRepo) CreateUrl(c *gin.Context) {
	var url entity.Url
	err := entity.CreateUrl(repository.Db, &url, c.Query("long"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, url)
	}
}
*/
