package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/api"
)

//ExtRouter Custom Router
func ExtRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	router := gin.Default()

	// route handling basic
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to API",
		})
	})

	// route handling for external
	router.GET("/ping", api.Ping)
	router.GET("/blog", api.GetAllBlogs)
	router.GET("/blog/:id", api.GetBlogs)

	// routes only accesiible if logged in
	authonly := router.Group("/")
	authonly.Use(api.LoginMiddleware())
	{
		authonly.POST("/blog", api.InsertBlog)
		authonly.PUT("/blog/:id", api.UpdateBlog)
		authonly.DELETE("/blog/:id", api.DeleteBlog)
	}

	//auth routes
	authGroup := router.Group("/auth")
	authGroup.POST("/signup", api.Signup)
	authGroup.POST("/login", api.Login)
	authGroup.GET("/refresh", api.Refresh)

	return router
}
