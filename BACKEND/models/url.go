package models

import (
	"learning-go/helpers"

	"gorm.io/gorm"
)

type Url struct {
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeID uint64 `json:"snowflake" gorm:"primaryKey; UNIQUE"`
}

var HOST_NAME =  "http://localhost:8080/"

func GetUrl(db *gorm.DB, url *Url, id string) Url {
	db.First(&url, helpers.IdBase10(id))
	return *url
}

//CreateUrl ... Insert New data
func CreateUrl(db *gorm.DB,url *Url, longUrl string) (err error) {
	url.SnowflakeID = helpers.CreateID()
	url.LongURL = longUrl
	url.ShortURL = HOST_NAME + helpers.IdBase62(url.SnowflakeID)
	if err = db.Create(url).Error; err != nil {
		return err
	}
	return nil
}
