package tags

import (
	"context"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	p *pgxpool.Pool
	q *database.Queries
}

func NewService(pool *pgxpool.Pool, query *database.Queries) *Service {
	return &Service{
		p: pool,
		q: query,
	}
}

func (svc *Service) GetTag(id string) (ResponseTag, error) {

	var tagId pgtype.UUID

	responseTag := ResponseTag{}

	if err := tagId.Scan(id); err != nil {
		return responseTag, err
	}

	tag, err := svc.q.GetTag(context.Background(), tagId)

	if err != nil {
		return responseTag, err
	}

	tag.ID.Scan(responseTag)

	return responseTag, nil
}

func (svc *Service) GetAllTags() ([]ResponseTag, error) {

	tags, err := svc.q.ListTags(context.Background())

	if err != nil {
		return nil, err
	}

	respTag := make([]ResponseTag, 0, len(tags))

	for _, tag := range tags {
		respTag = append(respTag, TagToResponse(tag))
	}

	return respTag, nil

}
