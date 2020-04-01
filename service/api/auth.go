package api

import (
	"log"

	"github.com/gin-gonic/gin"
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
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	password = generateHSPassword(password)

	ctx.JSON(200, gin.H{
		"username":    username,
		"hashed_pass": password,
	})
}

//Login -> route for login users
func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	exampleHash := "$2a$12$15C6HjXMRwYns0l.DgXGru56n69L5BmgYLmYYfGXmHT1VcH3ezlEK"
	cResult := compareHSPassword(exampleHash, password)

	ctx.JSON(200, gin.H{
		"username":     username,
		"login_status": cResult,
	})
}
