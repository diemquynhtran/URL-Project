package main

import (
	"learning-go/route"
)

func main() {
	r := routes.InitRouter()
	_ = r.Run(":8080")
}
