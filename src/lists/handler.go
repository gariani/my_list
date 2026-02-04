package lists

import (
	"net/http"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetAllListsHandler(svc *Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		userId := c.GetString("userId")

		if userId == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id not informed"})
			return
		}

		allLists, err := svc.GetAllListByUserId(userId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error retrieving user's lists", "errorMsg": err})
			return
		}

		if len(allLists) <= 0 {
			allLists = []database.List{}
		}

		c.JSON(http.StatusOK, gin.H{
			"data": allLists,
		})
	}
}

func GetListHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		listId := c.Param("id")

		var id pgtype.UUID

		if err := id.Scan(listId); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid list id"})
			return
		}

		list, err := svc.GetList(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "list not found"})
			return
		}

		c.JSON(http.StatusOK, list)

	}
}

func CreateUserListHandler(svc *Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		userId := c.GetString("userId")

		if userId == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id not informed"})
			return
		}

		var req CreateListRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		newReq, err := svc.CreateUserList(userId, req)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusCreated, newReq)
	}
}

func DeleteListHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		listId := c.Param("id")

		var id pgtype.UUID

		if err := id.Scan(listId); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid list id"})
			return
		}

		err := svc.DeleteList(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "list id not found"})
			return
		}

		c.Status(http.StatusNoContent)

	}
}
