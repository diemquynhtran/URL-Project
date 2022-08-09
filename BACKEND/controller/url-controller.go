package controller

import (
	"fmt"
	"learning-go/cache"
	"learning-go/dto"
	"learning-go/entity"
	"learning-go/helper"
	"learning-go/service"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type UrlController interface {
	GetUrlByName(ctx *gin.Context)
	CreateUrl(ctx *gin.Context)
	GetUrlByUser(ctx *gin.Context)
	//GetUrlCache(ctx *gin.Context)
	GetFreeUrl(ctx *gin.Context)

	CreateFreeUrl(ctx *gin.Context)
	EditUrl(ctx *gin.Context)
	DeleteUrl(ctx *gin.Context)
}

type urlController struct {
	urlService service.UrlService
	jwtService service.JWTService
	urlCache   cache.UrlCache
}

func NewUrlController(urlService service.UrlService, jwtService service.JWTService, urlCache cache.UrlCache) UrlController {
	return &urlController{
		urlService: urlService,
		jwtService: jwtService,
		urlCache:   urlCache,
	}
}

func (c *urlController) CreateUrl(ctx *gin.Context) {
	var urlDTO dto.CreateUrl
	errDTO := ctx.BindQuery(&urlDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			urlDTO.UserID = convertedUserID
		}
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
		result := c.urlService.GetUrlById(id)
		if result == nil {
			res := helper.BuildErrorResponse("Not found url", errDTO.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
		}
		if result.Status == false {
			res := helper.BuildErrorResponse("Url was blocked", errDTO.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
		}

		ctx.Redirect(302, result.LongURL)
	}
}

func (c *urlController) GetUrlByUser(ctx *gin.Context) {
	var urlRepoDto []entity.Url
	authHeader := ctx.GetHeader("Authorization")
	id := c.getUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No userId", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	urlRepoDto = c.urlService.GetUrlByUser(convertedUserID)
	ctx.JSON(http.StatusOK, urlRepoDto)
}

func (c *urlController) EditUrl(ctx *gin.Context) {
	urlId := ctx.Param("urlid")
	convertedUrlID, err := strconv.ParseUint(urlId, 10, 64)
	authHeader := ctx.GetHeader("Authorization")
	id := c.getUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, "Id user invalid")
	}
	var urlDTo dto.UrlResponse
	url := c.urlService.EditUrl(convertedUrlID, convertedUserID)
	if url == nil {
		res := helper.BuildErrorResponse("You are not permission EDIT", "", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	e := smapping.FillStruct(&urlDTo, smapping.MapFields(&url))
	if e != nil {
		log.Fatalf("Failed map %v", e)
	}
	response := helper.BuildResponse(true, "Create url successfully", urlDTo)
	ctx.JSON(http.StatusOK, response)
}

func (c *urlController) DeleteUrl(ctx *gin.Context) {
	urlId := ctx.Param("urlid")
	convertedUrlID, _ := strconv.ParseUint(urlId, 10, 64)
	var url entity.Url
	url.SnowflakeId = convertedUrlID
	res := c.urlService.DeleteUrl(url)
	if !res {
		res := helper.BuildErrorResponse("Not found url", "", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	response := helper.BuildResponse(true, "", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *urlController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

/*func (c *urlController) GetUrlCache(ctx *gin.Context) {
		id := helper.IdBase10(ctx.Param("cache"))
		idString :=strconv.Itoa(int(id))
		var urlOnCache *entity.Url = c.urlCache.Get(idString)
		if urlOnCache == nil {
			result := c.urlService.GetUrlById(id)
			c.urlCache.Set(idString, result)
			fmt.Println("heee")
			ctx.Redirect(302, result.LongURL)
		} else {
			fmt.Println("ho")
		ctx.Redirect(302, urlOnCache.LongURL)
		}
}*/

func (c *urlController) GetFreeUrl(ctx *gin.Context) {
	var u entity.FreeUrl
	errDTO := ctx.ShouldBind(&u)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		id := helper.IdBase10(ctx.Param("short"))
		/*idString :=strconv.Itoa(int(id))
		var urlOnCache *entity.FreeUrl = c.urlCache.Get(idString)
		if urlOnCache == nil {
			result := c.urlService.GetFreeUrl(id)
			c.urlCache.Set(idString, result)
			fmt.Println("heee")
			ctx.Redirect(302, result.LongURL)
		} else {
			fmt.Println("ho")
		ctx.Redirect(302, urlOnCache.LongURL)
		}*/

		result := c.urlService.GetFreeUrl(id)
		ctx.Redirect(302, result.LongURL)
	}

}
func (c *urlController) CreateFreeUrl(ctx *gin.Context) {
	var newUrl entity.FreeUrl
	errDTO := ctx.ShouldBind(&newUrl)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {

		result := c.urlService.CreateFreeUrl(newUrl)
		response := helper.BuildResponse(true, "Create url successfully", result)
		ctx.JSON(http.StatusCreated, response)
	}
}
