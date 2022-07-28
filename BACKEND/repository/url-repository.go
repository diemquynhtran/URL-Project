package repository

import (
	"learning-go/entity"

	"gorm.io/gorm"
)

type UrlRepository interface {
	CreateUrl(u entity.Url) entity.Url
	GetUrlByUser(userId uint64) []entity.Url
	GetUrlByName(urlId uint64) entity.Url
}

type urlConnection struct {
	connection *gorm.DB
}

func NewUrlRepository(connection *gorm.DB) UrlRepository {
	return &urlConnection{
		connection: connection,
	}
}
func (db *urlConnection) CreateUrl(u entity.Url) entity.Url {
	db.connection.Save(&u)
	return u
}

func (db *urlConnection) GetUrlByName(urlId uint64) entity.Url {
	var url entity.Url
	db.connection.Where("snowflake_id = ?", urlId).Take(&url)
	return url
}

func (db *urlConnection) GetUrlByUser(userId uint64) []entity.Url {
	var urls []entity.Url
	return urls
}
