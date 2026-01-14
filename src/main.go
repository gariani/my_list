package main

import (
	"context"
	"fmt"
	"log"
	"os"

	connection "github.com/gariani/my_list/src/database"
	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/routers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func main() {

	godotenv.Load()

	dsnUrl := os.Getenv("DATABASE_URL")

	pool, err := connection.Connect(context.Background(), dsnUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	queries := database.New(pool) // ✅ create once

	defer Pool.Close()

	r := routers.SetupRouter(pool, queries)

	if err := r.RunTLS(":8080", "cert.pem", "key.pem"); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

}
