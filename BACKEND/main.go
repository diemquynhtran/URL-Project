package main

import (
	"learning-go/routes"

)

func main() {
	r := routes.InitRouter()
	//r.Use(cors.Default())
	_ = r.Run(":8080")
}
