package api

import (
	"github.com/rahulsai1999/go-rest/service/db"
	"go.mongodb.org/mongo-driver/mongo"
)

// initializes single database collection object for all the api requests
var collection *mongo.Collection

func init() {
	collection = db.ConnectClient().Database("golang-api").Collection("blogs")
}
