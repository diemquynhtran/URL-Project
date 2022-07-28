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
	GetUrlById(urlId uint64) entity.Url
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

func (service *urlService) GetUrlById(urlId uint64) entity.Url {
	return service.urlRepository.GetUrlByName(urlId)
}
