package models

//import "gorm.io/gorm"

type Url struct {
	LongURL     string `json:"longUrl"`
	ShortURL    string `json:"shortUrl" gorm:"UNIQUE"`
	SnowflakeID uint64 `json:"id" gorm:"primary_key;UNIQUE"`
}
