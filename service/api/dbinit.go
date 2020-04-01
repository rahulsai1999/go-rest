package api

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rahulsai1999/go-rest/service/db"
	"go.mongodb.org/mongo-driver/mongo"
)

// initializes single database collection object for all the api requests
var collectionBlogs *mongo.Collection
var collectionUsers *mongo.Collection
var jwtKey []byte
var costI int

func init() {
	godotenv.Load()
	jj := os.Getenv("JWT_KEY")
	cost := os.Getenv("COST")

	costI, _ = strconv.Atoi(cost)
	jwtKey = []byte(jj)
	collectionBlogs = db.ConnectClient().Database("golang-api").Collection("blogs")
	collectionUsers = db.ConnectClient().Database("golang-api").Collection("users")
}
