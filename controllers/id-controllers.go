package controllers

import (
	"fmt"
	snowfake "learning-go/snowflake"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateID(ctx *gin.Context) {
	snowfake.SetEpoch(1577836800) // timestamp start from 01/01/2020 @ 12:00am (UTC)
	snowfake.SetNodeBits(10)       // reserve 2^10 machine numbers
	snowfake.SetSeqBits(12)        // 2^12 unique ID per second
	_ = snowfake.Init()           // must be called to instantiate new config
	nodeID := uint64(1)

	node, _ := snowfake.CreateNode(nodeID)
	id := node.GenerateID()
	fmt.Print(id)

	ctx.JSON(http.StatusOK, gin.H{"status":"OK","data": id})
}