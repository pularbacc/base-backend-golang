package main

import (
	"log"
	"os"

	"pular.server/database"
	"pular.server/server"
)

func main() {
	os.Chdir("/home")

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	server.Init(db)
}
