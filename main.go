package main

import (
	"github.com/mmfalcao/go-basic-api/database"
	"github.com/mmfalcao/go-basic-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
