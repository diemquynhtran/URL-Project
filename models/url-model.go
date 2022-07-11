package models

//import "gorm.io/gorm"

type Url struct {
	//Id          string `json:"id" gorm:"primary_key"`
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeID uint64 `json:"snowflake" gorm:"primaryKey; UNIQUE"`
}
