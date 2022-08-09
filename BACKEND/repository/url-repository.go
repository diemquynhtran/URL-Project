package repository

import (
	"learning-go/entity"

	"gorm.io/gorm"
)

type UrlRepository interface {
	CreateUrl(u entity.Url) entity.Url
	DeleteUrl(u entity.Url) bool
	GetUrlByUser(userId uint64) []entity.Url
	GetUrlById(urlId uint64) *entity.Url
	EditUrl(urlId uint64, userId uint64) *entity.Url

	CreateFreeUrl(u entity.FreeUrl) entity.FreeUrl
	GetFreeUrl(urlId uint64) entity.FreeUrl
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
	db.connection.Preload("User").Find(&u)
	return u
}
/////
func (db *urlConnection) GetUrlById(urlId uint64) *entity.Url {
	var url *entity.Url
	if err := db.connection.Where("snowflake_id = ?", urlId).Take(&url).Error; err != nil {
		return nil
	}
	url.NumberClick += 1
	db.connection.Save(url)
	return url
}

func (db *urlConnection) GetUrlByUser(userId uint64) []entity.Url {
	var urls []entity.Url
	db.connection.Preload("User").Where("user_id = ?", userId).Find(&urls)
	return urls
}

func (db *urlConnection) DeleteUrl(u entity.Url) bool{
	if err:=db.connection.Delete(&u).Error; err!=nil {
		return false
	}
	return true
}

func (db *urlConnection) EditUrl(urlId uint64,userId uint64) *entity.Url {
	var url *entity.Url
	if err := db.connection.Where("snowflake_id = ? AND user_id = ?", urlId, userId).Take(&url).Error; err != nil{
		return nil
	}
	url.Status = !url.Status
	db.connection.Preload("User").Save(url)
	return url
}



func (db *urlConnection) CreateFreeUrl(u entity.FreeUrl) entity.FreeUrl {
	db.connection.Save(u)
	return u
}
func (db *urlConnection) GetFreeUrl(urlId uint64) entity.FreeUrl {
	var url entity.FreeUrl
	db.connection.Where("snowflake_id = ?", urlId).Take(&url)
	return url
}

