package tags

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTagHandler(svc *Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		id := c.Param("id")

		responseTag, err := svc.GetTag(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, responseTag)

	}

}

func GetAllTagHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		respTag, err := svc.GetAllTags()

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error})
			return
		}

		c.JSON(http.StatusFound, respTag)

	}
}
