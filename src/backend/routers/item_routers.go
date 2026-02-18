package routers

import (
	"github.com/gariani/my_list/ai"
	"github.com/gariani/my_list/internal/database"
	"github.com/gariani/my_list/items"
	"github.com/gariani/my_list/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ItemsRouters(r *gin.Engine, pool *pgxpool.Pool, queries *database.Queries, aiService ai.Service) {
	api := r.Group("/api")

	v1 := api.Group(("/v1"))

	v1.Use(middleware.AuthRequired(), middleware.VerifyCSRF())

	list := v1.Group("/lists/:id")

	itemService := items.NewService(queries, pool, aiService)

	list.GET("/items", items.GeAllItemsByListHandler(itemService))
	// v1.GET("/lists", lists.GetListsHandler(listService))
	list.POST("/items", items.CreateItemHandler(itemService))
	// v1.DELETE("/list/:id", lists.DeleteList(listService))

}
