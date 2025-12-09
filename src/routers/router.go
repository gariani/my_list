package routers

import (
	"github.com/gariani/my_list/src/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.POST("/refresh", auth.Refresh)
	r.POST("/logout", auth.Logout)

	api := r.Group("/api")
	api.Use(auth.AuthRequired())

	api.GET("/profile", func(c *gin.Context) {
		userId := c.GetString("userId")
		c.JSON(200, gin.H{"message": "Authenticated", "userId": userId})
	})

	return r

}
