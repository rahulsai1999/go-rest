package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rahulsai1999/go-rest/service"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	service.ExtRouter(port)
}
