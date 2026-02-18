package items

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func GeAllItemsByListHandler(svc *Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		listId := c.Param("id")
		userIdStr := c.GetString("userId")

		if listId == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "list id not informed"})
			return
		}

		if userIdStr == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		var userId pgtype.UUID

		if err := userId.Scan(userIdStr); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":    "error retrieving the list of items",
				"errorMsg": err.Error(),
			})
			return
		}

		allItems, err := svc.GetAllItemsByList(userId, listId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error retrieving the list's item",
				"errorMsg": err,
			})
			return
		}

		if len(allItems) <= 0 {
			allItems = []ItemResponse{}
		}

		c.JSON(http.StatusOK, gin.H{
			"data": allItems,
		})
	}
}

func CreateItemHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItemRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item, err := svc.CreateItem(req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, item)

	}
}
