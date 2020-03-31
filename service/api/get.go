package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Ping custom message
func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

//GetAllBlogs -> gets all blogs in db
func GetAllBlogs(ctx *gin.Context) {
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(5)

	var results []models.Blog
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Blog
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"result": results,
		})
	}
}

//GetBlogs -> gets blogs given id
func GetBlogs(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Blog{}
	filter := bson.M{"_id": docID}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"result": result,
		})
	}
}
