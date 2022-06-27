package main

import (
	"fmt"
	"learning-go/snowflake"
)

func main() {
    node,_ := snowflake.CreateNewNode(1)
    id := node.Generate()
    fmt.Print(id)
}
