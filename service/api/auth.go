package api

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rahulsai1999/go-rest/service/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func generateHSPassword(pwd string) string {
	BBpwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(BBpwd, costI)
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
	} else if cResult == true {
		expirationTime := time.Now().Add(10 * time.Minute)
		claims := &models.Claims{
			Email: result.Email,
			ID:    result.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error. Try Again",
			})
		} else {
			ctx.JSON(200, gin.H{
				"email": result.Email,
				"token": tokenString,
			})
		}
	} else {
		ctx.JSON(200, gin.H{
			"message": "Wrong Password",
		})
	}
}

//Refresh -> to refresh the jwt tokens (background task)
func Refresh(ctx *gin.Context) {
	h := models.Header{}
	err := ctx.ShouldBindHeader(&h)
	if err != nil {
		ctx.JSON(200, err)
	}

	tknStr := strings.Split(h.Authorization, " ")[1]
	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		ctx.JSON(200, gin.H{
			"message": "Unauthorized",
		})
	} else if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{})
		}
		ctx.JSON(http.StatusBadRequest, gin.H{})

	} else {
		expirationTime := time.Now().Add(10 * time.Minute)
		claims.ExpiresAt = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		} else {
			ctx.JSON(200, gin.H{
				"old_token": tknStr,
				"new_token": tokenString,
			})
		}
	}

}

// LoginMiddleware -> check if the user is logged in
func LoginMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := models.Header{}
		err := ctx.ShouldBindHeader(&h)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "err",
			})
		}

		if len(h.Authorization) == 0 {
			ctx.JSON(200, gin.H{
				"message": "NO TOKEN || Login to access this route",
			})
			ctx.Abort()
		} else {
			tknStr := strings.Split(h.Authorization, " ")[1]
			claims := &models.Claims{}
			tkn, _ := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if !tkn.Valid {
				ctx.JSON(200, gin.H{
					"message": "Token Invalid, Login to access this route",
				})
				ctx.Abort()
			} else {
				ctx.Next()
			}
		}

	}
}
