package models


//import "gorm.io/gorm"

type Url struct {
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeID uint64 `json:"id" gorm:"primary_key;UNIQUE"`
}

