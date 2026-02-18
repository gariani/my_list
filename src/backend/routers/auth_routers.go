package routers

import (
	"net/http"

	"github.com/gariani/my_list/auth"
	"github.com/gariani/my_list/internal/database"
	"github.com/gariani/my_list/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRouter(r *gin.Engine, pool *pgxpool.Pool, queries *database.Queries) {

	service := auth.NewService(pool, queries)

	r.POST("/register", auth.Register(service))
	r.POST("/login", auth.Login(service))

	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthRequired(), middleware.VerifyCSRF())
	authGroup.POST("/refresh", auth.Refresh(service))
	authGroup.POST("/logout", auth.Logout(service))

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.Use(middleware.AuthRequired(), middleware.VerifyCSRF())

	v1.GET("/profile", func(c *gin.Context) {
		userId := c.GetString("userId")
		if userId == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No userId in context"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Authenticated", "userId": userId})
	})

}
