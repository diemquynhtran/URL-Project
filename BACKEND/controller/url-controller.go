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
	errDTO := ctx.BindQuery(&urlDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result := c.urlService.CreateUrl(urlDTO)
		response := helper.BuildResponse(true, "Create url successfully", result)
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
