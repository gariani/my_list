package items

import (
	"context"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	query *database.Queries
}

func NewService(q *database.Queries) *Service {
	return &Service{
		query: q,
	}
}

func (s *Service) GetAllItemsByList(userID pgtype.UUID, id string) ([]database.Item, error) {

	var listId pgtype.UUID

	err := listId.Scan(id)
	if err != nil {
		return nil, err
	}

	param := database.GetAllItemsByListParams{UserID: userID, ID: listId}

	items, err := s.query.GetAllItemsByList(context.Background(), param)

	if err != nil {
		return nil, err
	}

	return items, nil
}
