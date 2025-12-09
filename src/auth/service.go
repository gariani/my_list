package auth

import (
	"context"
	"errors"
	"log"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/models"
	"github.com/gariani/my_list/src/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func RegisterUser(email, password string) error {
	hashed, err := utils.HashPassword(password)

	if err != nil {
		log.Fatal("saving user: ", err.Error())
		return err
	}

	user, err := GetUserByEmail(email)

	if err != nil && err != pgx.ErrNoRows {
		log.Fatal("saving user", err.Error())
		return err
	}

	if user != nil {
		return errors.New("user already exists")
	}

	_, err = database.DB.Exec(context.Background(), `INSERT INTO users (id, email, pass_hash) VALUES ($1, $2, $3)`, uuid.New().String(), email, hashed)

	return err

}

func GetUserByEmail(email string) (*models.User, error) {

	row := database.DB.QueryRow(context.Background(), `SELECT id, email, pass_hash, created_at FROM users WHERE email=$1`, email)

	var u models.User

	err := row.Scan(&u.Id, &u.Email, &u.PassHash, &u.CreatedAt)

	if err == pgx.ErrNoRows {
		return nil, pgx.ErrNoRows
	}

	return &u, err

}
