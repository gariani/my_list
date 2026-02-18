package routers

import (
	"time"

	"github.com/gariani/my_list/ai"
	"github.com/gariani/my_list/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(pool *pgxpool.Pool, queries *database.Queries, aiService ai.Service) *gin.Engine {

	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1", "::1", "172.16.0.0/12", "192.168.0.0/16", "10.0.0.0/8"})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "https://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	AuthRouter(r, pool, queries)
	ListRouters(r, pool, queries, aiService)
	ItemsRouters(r, pool, queries, aiService)
	TagRouters(r, pool, queries)

	return r
}
