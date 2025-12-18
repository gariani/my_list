package lists

import (
	"context"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/utils"
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

func (s *Service) GetAllLists(id string) (*[]database.List, error) {

	var userId pgtype.UUID
	err := userId.Scan(id)
	if err != nil {
		return nil, err
	}

	allLists, err := s.query.ListUserLists(context.Background(), userId)

	if err != nil {
		return nil, err
	}

	return &allLists, nil

}

func (s *Service) CreateUserList(id string, list *UserList) (*UserList, error) {

	var userId pgtype.UUID

	err := userId.Scan(id)
	if err != nil {
		return nil, err
	}

	listParam := database.CreateListParams{}

	listParam.ID = utils.GetNewId()
	listParam.Name = list.Name
	listParam.UserID = userId

	tx, err := database.DB.Begin(context.Background())

	if err != nil {
		return nil, err
	}

	tx.Rollback(context.Background())

	qtx := s.query.WithTx(tx)

	query, err := qtx.CreateList(context.Background(), listParam)

	if err != nil {
		return nil, err
	}

	newList := &UserList{}
	newList.Id = query.ID.String()
	newList.Name = query.Name
	newList.UserId = query.UserID.String()

	return newList, tx.Commit(context.Background())

}

func (s *Service) DeleteList(id pgtype.UUID) error {

	tx, err := database.DB.Begin(context.Background())

	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	qtx := s.query.WithTx(tx)

	if err := qtx.DeleteList(context.Background(), id); err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (s *Service) GetList(id pgtype.UUID) (*database.List, error) {
	tx, err := database.DB.Begin(context.Background())

	if err != nil {
		return nil, err
	}

	defer tx.Rollback(context.Background())

	qtx := s.query.WithTx(tx)

	query, err := qtx.GetList(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &query, tx.Commit(context.Background())

}
