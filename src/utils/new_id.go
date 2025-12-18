package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetNewId() pgtype.UUID {
	id := uuid.New()

	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

}
