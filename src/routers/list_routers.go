package routers

import (
	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/lists"
	"github.com/gariani/my_list/src/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ListRouters(r *gin.Engine, p *pgxpool.Pool, q *database.Queries) {

	api := r.Group("/api")

	v1 := api.Group(("/v1"))

	v1.Use(middleware.AuthRequired(), middleware.VerifyCSRF())

	listService := lists.NewService(p, q)

	v1.GET("/list/:id", lists.GetListHandler(listService))
	v1.GET("/lists", lists.GetAllListsHandler(listService))
	v1.POST("/list", lists.CreateUserListHandler(listService))
	v1.DELETE("/list/:id", lists.DeleteListHandler(listService))

}
