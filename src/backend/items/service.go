package items

import (
	"context"

	"github.com/gariani/my_list/internal/database"
	"github.com/gariani/my_list/tags"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	pool  *pgxpool.Pool
	query *database.Queries
}

func NewService(q *database.Queries, p *pgxpool.Pool) *Service {
	return &Service{
		pool:  p,
		query: q,
	}
}

func (s *Service) GetAllItemsByList(userID pgtype.UUID, id string) ([]ItemResponse, error) {

	var listId pgtype.UUID

	err := listId.Scan(id)
	if err != nil {
		return nil, err
	}

	param := database.GetItemsWithTypeByListParams{UserID: userID, ID: listId}

	itemsWithType, err := s.query.GetItemsWithTypeByList(context.Background(), param)

	if err != nil {
		return nil, err
	}

	if len(itemsWithType) <= 0 {
		return []ItemResponse{}, nil
	}

	itemsId := make([]pgtype.UUID, 0, len(itemsWithType))
	for _, itemId := range itemsWithType {
		itemsId = append(itemsId, itemId.ID)
	}

	tagsData, err := s.query.GetTagsByItems(context.Background(), itemsId)

	itemTagsMap := make(map[string][]tags.ResponseTag)
	if err == nil {
		for _, tagData := range tagsData {
			itemsIdStr := tagData.ItemID.String()
			tag := tags.ResponseTag{
				Id:   tagData.TagID.String(),
				Name: tagData.TagName,
			}

			itemTagsMap[itemsIdStr] = append(itemTagsMap[itemsIdStr], tag)
		}
	}

	responses := make([]ItemResponse, 0, len(itemsWithType))

	for _, item := range itemsWithType {
		resp := ToItemResponseByListRow(item)

		if item.Title.Valid {
			resp.Title = item.Title.String

		}

		if item.Content.Valid {
			resp.Content = item.Content.String
		}

		if item.Url.Valid {
			resp.Url = item.Url.String
		}

		if item.Thumbnail.Valid {
			resp.Thumbnail = item.Thumbnail.String
		}

		if item.TypeName.Valid {
			resp.Type = &ItemTypeInfo{
				ID:          item.TypeID.String(),
				Name:        item.TypeName.String,
				Description: item.TypeDescription.String,
			}
		}

		if itemTags, exists := itemTagsMap[item.ID.String()]; exists {
			resp.Tags = itemTags
		} else {
			resp.Tags = []tags.ResponseTag{}
		}

		responses = append(responses, resp)
	}

	return responses, nil
}

func (s *Service) CreateItem(req CreateItemRequest) (ItemResponse, error) {

	var listID, typeID pgtype.UUID
	if err := listID.Scan(req.ListID); err != nil {
		return ItemResponse{}, err
	}
	if err := typeID.Scan(req.TypeID); err != nil {
		return ItemResponse{}, err
	}

	tx, err := s.pool.Begin(context.Background())

	defer func() {

		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	title := pgtype.Text{}
	title.Scan(req.Title)

	content := pgtype.Text{}
	content.Scan(req.Content)

	thumbnail := pgtype.Text{}
	thumbnail.Scan(req.Thumbnail)

	url := pgtype.Text{}
	url.Scan(req.URL)

	param := database.CreateItemParams{
		ListID:    listID,
		TypeID:    typeID,
		Title:     title,
		Content:   content,
		Url:       url,
		Thumbnail: thumbnail,
	}

	qtx := s.query.WithTx(tx)

	databaseItem, err := qtx.CreateItem(context.Background(), param)

	if err != nil {
		return ItemResponse{}, err
	}

	return ToItemResponse(databaseItem), tx.Commit(context.Background())
}
