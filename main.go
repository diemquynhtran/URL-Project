package main

import (
	"learning-go/routes"

	//"gorm.io/gorm"

)

var err error

func main() {
	r := routes.InitRouter()
	_ = r.Run(":8080")


}


// func main() {
// 	dsn := "root:quynh@tcp(localhost:3306)/urldb?charset=utf8mb4&parseTime=True&loc=Local"
//   	db, _ := gorm.Open(mysql.Open(dsn), &gorm.config{})
// 	db.AutoMigrate(&models.Url{})
// 	r := routes.InitRouter()
// 	 //running
// 	 r.Run()

// }