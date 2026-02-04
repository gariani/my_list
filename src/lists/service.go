package lists

import (
	"context"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	pool  *pgxpool.Pool
	query *database.Queries
}

func NewService(p *pgxpool.Pool, q *database.Queries) *Service {
	return &Service{
		pool:  p,
		query: q,
	}
}

func (s *Service) GetAllListByUserId(id string) ([]database.List, error) {

	var userId pgtype.UUID
	err := userId.Scan(id)
	if err != nil {
		return nil, err
	}

	allLists, err := s.query.ListUserLists(context.Background(), userId)

	if err != nil {
		return nil, err
	}

	return allLists, nil
}

func (s *Service) GetList(id pgtype.UUID) (*database.List, error) {

	query, err := s.query.GetList(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &query, nil
}

func (s *Service) CreateUserList(id string, list CreateListRequest) (ListResponse, error) {

	var userId pgtype.UUID

	err := userId.Scan(id)
	if err != nil {
		return ListResponse{}, err
	}

	listParam := database.CreateListParams{}

	listParam.ID = utils.GetNewId()
	listParam.Name = list.Name
	listParam.UserID = userId

	tx, err := s.pool.Begin(context.Background())

	if err != nil {
		return ListResponse{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	qtx := s.query.WithTx(tx)

	query, err := qtx.CreateList(context.Background(), listParam)

	if err != nil {
		return ListResponse{}, err
	}

	return ToListResponse(query), tx.Commit(context.Background())

}

func (s *Service) DeleteList(id pgtype.UUID) error {

	tx, err := s.pool.Begin(context.Background())

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	qtx := s.query.WithTx(tx)

	if err := qtx.DeleteList(context.Background(), id); err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
