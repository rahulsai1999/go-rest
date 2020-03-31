package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/models"
)

//PostForm  post route handler
func PostForm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	message := ctx.PostForm("message")
	ctx.JSON(200, gin.H{
		"name":    name,
		"message": message,
	})
}

// InsertBlog -> insert one blog
func InsertBlog(ctx *gin.Context) {
	title := ctx.PostForm("title")
	author := ctx.PostForm("author")
	body := ctx.PostForm("body")

	blog := models.Blog{
		Title:  title,
		Author: author,
		Body:   body,
	}

	result, err := collection.InsertOne(context.Background(), blog)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"inserted_id": result.InsertedID,
		})
	}
}
