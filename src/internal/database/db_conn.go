package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	DB, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal("DATABASE ERROR: ", err)
		return err
	}

	return nil
}
