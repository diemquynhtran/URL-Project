package main

import (
	"learning-go/routes"

)

func main() {
	r := routes.InitRouter()
	_ = r.Run(":8080")
}
