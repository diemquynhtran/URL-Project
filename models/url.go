package models

import (
	"learning-go/helpers"

	"gorm.io/gorm"
)

func GetUrl(db *gorm.DB, url *Url, id string) Url {
	db.First(&url, helpers.IdBase10(id))
	return *url
}

//CreateUrl ... Insert New data
func CreateUrl(db *gorm.DB,url *Url, longUrl string) (err error) {
	url.SnowflakeID = helpers.CreateID()
	url.LongURL = longUrl
	url.ShortURL = "http://localhost:8080/" + helpers.IdBase62(url.SnowflakeID)
	if err = db.Create(url).Error; err != nil {
		return err
	}
	return nil




	// if err = db.Create(url).Error; err != nil {
	// 	return err
	// }
	// return nil
}
