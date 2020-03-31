package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectClient -> returns db client connected to the provided database
func ConnectClient() *mongo.Client {

	err := godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	finalDB := "mongodb://" + dbUser + ":" + dbPass + "@" + dbURL

	clientOptions := options.Client().ApplyURI(finalDB)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// Ping -> checks database connection and returns the client object
func Ping() *mongo.Client {
	client := ConnectClient()
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	return client
}
