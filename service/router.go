package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/api"
)

//ExtRouter Custom Router
func ExtRouter(port string) {
	router := gin.Default()

	// route handling basic
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to API",
		})
	})

	// route handling for external
	router.GET("/ping", api.Ping)
	router.POST("/add", api.PostForm)

	router.Run(port)
}
