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
	cur, err := collectionBlogs.Find(context.TODO(), filter, findOptions)
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
	err := collectionBlogs.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"result": result,
		})
	}
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

	result, err := collectionBlogs.InsertOne(context.Background(), blog)

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

	docID, _ := primitive.ObjectIDFromHex(id)
	blog := models.Blog{
		Title:  title,
		Author: author,
		Body:   body,
	}
	update := bson.M{
		"$set": blog,
	}

	filter := bson.M{"_id": docID}
	result := models.Blog{}
	err := collectionBlogs.FindOneAndUpdate(context.Background(), filter, update).Decode(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"updated_id": result.ID,
		})
	}

}

//DeleteBlog -> deletes blog based on id
func DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Blog{}
	filter := bson.M{"_id": docID}
	err := collectionBlogs.FindOneAndDelete(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"deleted_id": result.ID,
		})
	}
}
