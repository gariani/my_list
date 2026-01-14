package routers

import (
	"github.com/gariani/my_list/src/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(pool *pgxpool.Pool, queries *database.Queries) *gin.Engine {

	r := gin.Default()

	AuthRouter(r, pool, queries)
	ListRouters(r, pool, queries)
	ItemsRouters(r, pool, queries)

	return r
}
