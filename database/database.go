package database

import (
	"fmt"
	"learning-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "quynh"
	DB_NAME     = "shortenId"
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDb()
	return Db
}

func connectDb() *gorm.DB {
	var err error
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database, error: %v", err)
		return nil
	}

	db.AutoMigrate(&models.Url{})
	return db
}
