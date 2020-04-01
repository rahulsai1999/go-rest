package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User -> model for User object
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Email string             `bson:"email,omitempty"`
	Hash  string             `bson:"hash,omitempty"`
}

//Claims -> model for jwt claim
type Claims struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Email string             `json:"email"`
	jwt.StandardClaims
}

//Header -> model for decoding header
type Header struct {
	Authorization string `header:"Authorization"`
}
