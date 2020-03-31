package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Ping custom message
func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

//GetBlogs -> gets blogs given id
func GetBlogs(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Blog{}

	err := collection.FindOne(context.Background(), bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"result": result,
		})
	}
}
