package models

import (
	"learning-go/helpers"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Url struct {
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeID uint64 `json:"snowflake" gorm:"primaryKey; UNIQUE"`
		
}

//Get a url
func GetUrl(db *gorm.DB, url *Url, id string) Url {
	db.First(&url, helpers.IdBase10(id))
	return *url
}

//Create new url 
func CreateUrl(db *gorm.DB,url *Url, longUrl string) (err error) {
	e:= godotenv.Load()
	if e != nil {
		panic("Open .env failed")
	}
	url.SnowflakeID = helpers.CreateID()
	url.LongURL = longUrl
	url.ShortURL = os.Getenv("HOST_NAME") + helpers.IdBase62(url.SnowflakeID)
	if err = db.Create(url).Error; err != nil {
		return err
	}
	return nil
}
