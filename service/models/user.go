package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User -> model for User object
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Email string             `bson:"email,omitempty"`
	Hash  string             `bson:"hash,omitempty"`
}
