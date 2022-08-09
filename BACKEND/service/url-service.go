package service

import (
	"learning-go/dto"
	"learning-go/entity"
	"learning-go/helper"
	"learning-go/repository"
	"log"
	"os"

	"github.com/mashingan/smapping"
)

type UrlService interface {
	CreateUrl(u dto.CreateUrl) entity.Url
	GetUrlById(urlId uint64) *entity.Url
	GetUrlByUser(id uint64) []entity.Url

	EditUrl(urlId uint64,userId uint64) *entity.Url
	DeleteUrl(u entity.Url) bool

	CreateFreeUrl(u entity.FreeUrl) entity.FreeUrl
	GetFreeUrl(urlId uint64) entity.FreeUrl
}

type urlService struct {
	urlRepository repository.UrlRepository
}

func NewUrlService(urlRepo repository.UrlRepository) UrlService {
	return &urlService{
		urlRepository: urlRepo,
	}
}

func (service *urlService) CreateUrl(urlDto dto.CreateUrl) entity.Url {
	url := entity.Url{}
	err := smapping.FillStruct(&url, smapping.MapFields(&urlDto))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	url.LongURL = urlDto.LongURL
	url.SnowflakeId = helper.CreateID()
	url.ShortURL = os.Getenv("HOST_NAME") + helper.IdBase62(url.SnowflakeId)
	res := service.urlRepository.CreateUrl(url)
	return res
}

func (service *urlService) GetUrlById(urlId uint64) *entity.Url {
	return service.urlRepository.GetUrlById(urlId)
}

func (service *urlService) GetUrlByUser(id uint64) []entity.Url{
	return service.urlRepository.GetUrlByUser(id)
}

func (service *urlService) EditUrl(urlId uint64,userId uint64) *entity.Url {
	return service.urlRepository.EditUrl(urlId,userId)
}

func (service *urlService) DeleteUrl(u entity.Url) bool {
	return service.urlRepository.DeleteUrl(u)
}




func (service *urlService) CreateFreeUrl(u entity.FreeUrl) entity.FreeUrl {
	url := entity.FreeUrl{}
	url.LongURL = u.LongURL
	url.SnowflakeId = helper.CreateID()
	url.ShortURL = os.Getenv("HOST_NAME") + helper.IdBase62(url.SnowflakeId)
	res := service.urlRepository.CreateFreeUrl(url)
	return res
}
func (service *urlService)	GetFreeUrl(urlId uint64) entity.FreeUrl {
	return service.urlRepository.GetFreeUrl(urlId)
}