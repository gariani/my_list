package main

import (
	"fmt"
	"log"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/routers"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	if err := database.Connect(); err != nil {
		panic(1)
	}

	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to the DB", err)
	}

	defer database.DB.Close()

	r := routers.SetupRouter()

	if err := r.RunTLS(":8080", "cert.pem", "key.pem"); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

}
