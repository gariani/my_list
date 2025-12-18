package auth

import (
	"context"
	"errors"
	"log"

	"github.com/gariani/my_list/src/internal/database"
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

	tx, err := database.DB.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.Serializable, AccessMode: pgx.ReadWrite})

	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `INSERT INTO users (id, email, pass_hash) VALUES ($1, $2, $3)`, uuid.New().String(), email, hashed)

	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return tx.Commit(context.Background())
}

func GetUserByEmail(email string) (*database.User, error) {

	tx, err := database.DB.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.Serializable, AccessMode: pgx.ReadOnly})

	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(context.Background(), `SELECT id, email, pass_hash, created_at FROM users WHERE email=$1`, email)

	var u database.User

	err = row.Scan(&u.ID, &u.Email, &u.PassHash, &u.CreatedAt)

	if err == pgx.ErrNoRows {
		tx.Rollback(context.Background())
		return nil, pgx.ErrNoRows
	}

	return &u, err

}
