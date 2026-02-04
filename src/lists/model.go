package lists

import (
	"time"

	"github.com/gariani/my_list/src/internal/database"
)

type CreateListRequest struct {
	Name string `json:"name" binding:"required"`
}

type ListResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func ToListResponse(dbList database.List) ListResponse {
	return ListResponse{
		ID:        dbList.ID.String(),
		UserID:    dbList.UserID.String(),
		Name:      dbList.Name,
		CreatedAt: dbList.CreatedAt.Time.Format(time.RFC3339),
	}
}
