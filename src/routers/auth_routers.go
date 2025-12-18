package routers

import (
	"net/http"

	"github.com/gariani/my_list/src/auth"
	"github.com/gariani/my_list/src/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthRequired(), middleware.VerifyCSRF())
	authGroup.POST("/refresh", auth.Refresh)
	authGroup.POST("/logout", auth.Logout)

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
