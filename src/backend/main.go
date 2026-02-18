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

	"github.com/gariani/my_list/ai"
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

	ollamaUrl := "http://127.0.0.1:11434"

	model := "llama3:latest"
	aiService := ai.NewOllamaService(ollamaUrl, model)

	r := routers.SetupRouter(pool, queries, aiService)

	if err := r.RunTLS(":8080", "cert.pem", "key.pem"); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

}
