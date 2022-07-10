package models

import (
	"gorm.io/gorm"
)

func GetAllUrl(db *gorm.DB, url *[]Url) (err error) {
	if err = db.Find(url).Error; err != nil {
		return err
	}
	return nil
}

//CreateUrl ... Insert New data
func CreateUrl(db *gorm.DB,url *Url) (err error) {
	if err = db.Create(url).Error; err != nil {
		return err
	}
	return nil
}