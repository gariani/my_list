package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	AuthRouter(r)
	ListRouters(r)
	ItemsRouters(r)

	return r
}
