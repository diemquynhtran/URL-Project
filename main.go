package main

import (
	"learning-go/routes"

	//"gorm.io/gorm"

)

func main() {
	r := routes.InitRouter()
	_ = r.Run(":8080")


}
