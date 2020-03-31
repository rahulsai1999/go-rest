package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Blog -> model for Blog object
type Blog struct {
	ID     primitive.ObjectID `bson:"_id,omitempty`
	Title  string
	Author string
	Body   string
}
