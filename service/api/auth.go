package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func generateHSPassword(pwd string) string {
	BBpwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(BBpwd, 8)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

func compareHSPassword(hash string, pwd string) bool {
	BBhash := []byte(hash)
	BBpwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(BBhash, BBpwd)
	if err != nil {
		return false
	}
	return true
}

//Signup -> route for creating new users
func Signup(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	password = generateHSPassword(password)

	user := models.User{
		Name:  name,
		Email: email,
		Hash:  password,
	}

	result, err := collectionUsers.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"_id":         result.InsertedID,
			"name":        name,
			"email":       email,
			"hashed_pass": password,
		})
	}
}

//Login -> route for login users
func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	filter := bson.M{"email": email}
	result := models.User{}

	err := collectionUsers.FindOne(context.Background(), filter).Decode(&result)
	cResult := compareHSPassword(result.Hash, password)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "User Not Found",
		})
	} else {
		ctx.JSON(200, gin.H{
			"email":        result.Email,
			"login_status": cResult,
		})
	}
}
