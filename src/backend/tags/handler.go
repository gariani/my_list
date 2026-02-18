package tags

import (
	"net/http"

	"github.com/gariani/my_list/utils"
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

// @Summary Get all tags
// @Description Returns all tags
// @Tags tags
// @Produce json
// @Success 200 {array} tags.ResponseTag
// @Failure 401 {object} utils.ErrorResponse
// @Security BearerAuth
// @Security CSRF
// @Router /api/v1/tags [get]
func GetAllTagHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		respTag, err := svc.GetAllTags()

		if err != nil {
			c.JSON(http.StatusNotFound, utils.ErrorResponse{Message: "Failed to get tags", Code: http.StatusNotFound})
			return
		}

		c.JSON(http.StatusFound, respTag)
	}
}
