package api

import "github.com/gin-gonic/gin"

//Ping custom message
func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
