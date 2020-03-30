package main

import (
	"github.com/rahulsai1999/go-rest/service"
	"github.com/rahulsai1999/go-rest/service/db"
)

func main() {
	service.ExtRouter(":5000")
	db.Ping()
}
