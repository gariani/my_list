package models

import "time"

type User struct {
	Id        string    `db:"id"`
	Email     string    `db:"email"`
	PassHash  string    `db:"pass_hash"`
	CreatedAt time.Time `db:"created_at"`
}
