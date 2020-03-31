package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteBlog -> deletes blog based on id
func DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Blog{}
	filter := bson.M{"_id": docID}
	err := collection.FindOneAndDelete(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"deleted_id": result.ID,
		})
	}
}
