package routers

import (
	"github.com/gariani/my_list/src/internal/database"
	"github.com/gariani/my_list/src/lists"
	"github.com/gariani/my_list/src/middleware"
	"github.com/gin-gonic/gin"
)

func ListRouters(r *gin.Engine) {

	api := r.Group("/api")

	v1 := api.Group(("/v1"))

	v1.Use(middleware.AuthRequired(), middleware.VerifyCSRF())

	listService := lists.NewService(database.New(database.DB))

	v1.GET("/list/:id", lists.GetList(listService))
	v1.GET("/lists", lists.GetListsHandler(listService))
	v1.POST("/list", lists.CreateUserList(listService))
	v1.DELETE("/list/:id", lists.DeleteList(listService))

}
