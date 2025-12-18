package lists

import (
	"net/http"

	"github.com/gariani/my_list/src/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetListsHandler(svc *Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		userId := c.GetString("userId")

		if userId == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id not informed"})
		}

		allLists, err := svc.GetAllLists(userId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error retrieving user's lists", "errorMsg": err})
		}

		if allLists == nil || len(*allLists) <= 0 {
			allLists = &[]database.List{}
		}

		c.JSON(http.StatusOK, gin.H{
			"data": allLists,
		})
	}
}

func CreateUserList(svc *Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		userId := c.GetString("userId")

		if userId == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id not informed"})
		}

		var req UserList

		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		}

		newReq, err := svc.CreateUserList(userId, &req)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		}

		c.JSON(http.StatusCreated, newReq)
	}
}

func DeleteList(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		listId := c.Param("id")

		var id pgtype.UUID

		if err := id.Scan(listId); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid list id"})
		}

		err := svc.DeleteList(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "list id not found"})
		}

		c.Status(http.StatusNoContent)

	}
}

func GetList(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		listId := c.Param("id")

		var id pgtype.UUID

		if err := id.Scan(listId); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid list id"})
		}

		list, err := svc.GetList(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "list not found"})
		}

		c.JSON(http.StatusOK, list)

	}
}
