package tags

import "github.com/gariani/my_list/internal/database"

type RequestTag struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ResponseTag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func TagToResponse(tag database.Tag) ResponseTag {
	return ResponseTag{
		Id:   tag.ID.String(),
		Name: tag.Name,
	}
}
