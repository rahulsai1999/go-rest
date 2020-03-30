package service

import "github.com/gin-gonic/gin"

//PostForm  post route handler
func PostForm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	message := ctx.PostForm("message")
	ctx.JSON(200, gin.H{
		"name":    name,
		"message": message,
	})
}
