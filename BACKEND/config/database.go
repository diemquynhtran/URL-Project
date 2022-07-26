package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	e := godotenv.Load()
	if e != nil {
		fmt.Println("Error loading .env file")
	}
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, e := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		fmt.Printf("Error connecting to database : error=%v", e)
		return nil
	}
	fmt.Println("Connect database successful!!!")
	return db
}
func CloseDB(db *gorm.DB) {
	dbSQL, e := db.DB()
	if e != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
