package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

//UpdateBlog -> update one blog
func UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.PostForm("title")
	author := ctx.PostForm("author")
	body := ctx.PostForm("body")

	blog := models.Blog{
		Title:  title,
		Author: author,
		Body:   body,
	}

	update := bson.M{
		"$set": blog,
	}

	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Blog{}
	err := collection.FindOneAndUpdate(context.Background(), bson.M{"_id": docID}, update).Decode(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"updated_id": result.ID,
		})
	}

}
