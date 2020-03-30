package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service"
)

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()

	// route handling basic
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to API",
		})
	})

	// route handling for named functions
	router.GET("/ping", ping)

	router.POST("/add", service.PostForm)

	router.Run(":8000")
}
