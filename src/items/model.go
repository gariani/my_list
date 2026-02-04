package items

import (
	"time"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/tags"
)

type CreateItemRequest struct {
	ListID    string `json:"list_id" binding:"required,uuid"`
	TypeID    string `json:"type_id" binding:"required,uuid"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

type ItemResponse struct {
	ID        string             `json:"id"`
	ListID    string             `json:"list_id"`
	TypeID    string             `json:"type_id"`
	Title     string             `json:"title"`
	Content   string             `json:"content,omitempty"`
	Url       string             `json:"url,omitempty"`
	Thumbnail string             `json:"thumbnail,omitempty"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	Type      *ItemTypeInfo      `json:"type,omitempty"`
	Tags      []tags.ResponseTag `json:"tags,omitempty"`
}

type ItemTypeInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateItemRequest struct {
	Title     *string  `json:"title"`
	Content   *string  `json:"content"`
	URL       *string  `json:"url"`
	Thumbnail *string  `json:"thumbnail"`
	TagIDs    []string `json:"tag_ids"`
}

func ToItemResponseByListRow(dbItem database.GetItemsWithTypeByListRow) ItemResponse {
	return ItemResponse{
		ID:        dbItem.ID.String(),
		ListID:    dbItem.ListID.String(),
		TypeID:    dbItem.TypeID.String(),
		Title:     dbItem.Title.String,
		Content:   dbItem.Content.String,
		Url:       dbItem.Url.String,
		Thumbnail: dbItem.Thumbnail.String,
		CreatedAt: dbItem.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: dbItem.UpdatedAt.Time.Format(time.RFC3339),
	}
}

func ToItemResponse(dbItem database.Item) ItemResponse {
	return ItemResponse{
		ID:        dbItem.ID.String(),
		ListID:    dbItem.ListID.String(),
		TypeID:    dbItem.TypeID.String(),
		Title:     dbItem.Title.String,
		Content:   dbItem.Content.String,
		Url:       dbItem.Url.String,
		Thumbnail: dbItem.Thumbnail.String,
		CreatedAt: dbItem.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: dbItem.UpdatedAt.Time.Format(time.RFC3339),
	}
}
