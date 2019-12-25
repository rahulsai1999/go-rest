package main

import (
	"fmt"

	controller "github.com/go-rest/controllers"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func postform(c *gin.Context) {
	name := c.PostForm("name")
	message := c.PostForm("message")
	c.JSON(200, gin.H{
		"name":    name,
		"message": message,
	})
}

func main() {
	fmt.Println("Hello")
	router := gin.Default()

	// route handling with anonymous functions
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	//route handling with named functions
	router.GET("/ping", ping)

	//route handling with refactoring
	router.GET("/external", controller.External)

	//post routes
	router.POST("/normal", postform)

	router.Run("0.0.0.0:5000")
}
