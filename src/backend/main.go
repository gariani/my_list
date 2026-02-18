package main

// @title My List API
// @version 1.0
// @description API for managing lists
// @host localhost:8080
// @BasePath /
// @schemes https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey CSRF
// @in header
// @name X-CSRF-Token
import (
	"context"
	"fmt"
	"log"
	"os"

	connection "github.com/gariani/my_list/database"
	"github.com/gariani/my_list/internal/database"
	"github.com/gariani/my_list/routers"
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
