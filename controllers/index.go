package controllers

import (
	"learning-go/models"
	"learning-go/config"
	"gorm.io/gorm"
)

type UrlRepo struct {
	Db *gorm.DB
}

func New() *UrlRepo {
	db := config.InitDb()
	db.AutoMigrate(&models.Url{})
	return &UrlRepo{Db: db}
}