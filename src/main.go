package main

import (
	"fmt"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/routers"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	if err := database.Connect(); err != nil {
		panic(1)
	}

	database.Connect()

	r := routers.SetupRouter()
	r.RunTLS(":8080", "cert.pem", "key.pem")

	fmt.Println()

}
