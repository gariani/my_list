package routers

import (
	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/middleware"
	"github.com/gariani/my_list/src/tags"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TagRouters(r *gin.Engine, pool *pgxpool.Pool, queries *database.Queries) {
	api := r.Group("/api")

	v1 := api.Group(("/v1"))

	v1.Use(middleware.AuthRequired(), middleware.VerifyCSRF())

	tag := v1.Group("/tags")

	tagService := tags.NewService(pool, queries)
	tag.GET("/", tags.GetAllTagHandler(tagService))

	tag.GET("/tag", tags.GetTagHandler(tagService))

	// tag.POST("/items", tags.CreateItemHandler(tagService))

}
