package auth

import (
	"fmt"
	"net/http"

	"github.com/gariani/my_list/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(s *Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		var input LoginInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		if err := s.RegisterUser(input.Email, input.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user created"})

	}
}

// @Summary Login user
// @Description Login with email/password and return JWT + CSRF
// @Tags auth
// @Accept json
// @Produce json
// @Param login body auth.LoginRequest true "Login info"
// @Success 200 {object} auth.LoginResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /login [post]
func Login(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input LoginInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse{Message: "invalid input", Code: http.StatusBadRequest})
			return
		}

		user, err := s.GetUserByEmail(input.Email)
		if err != nil || !utils.CheckPassword(user.PassHash, input.Password) {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse{Message: fmt.Sprintf("invalid credentials %s", err.Error()), Code: http.StatusBadRequest})
			return
		}

		accessToken, err := GenerateAccessToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse{Message: "token error", Code: http.StatusBadRequest})
			return
		}

		refreshToken, err := GenerateRefreshToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse{Message: "token error", Code: http.StatusBadRequest})
			return
		}

		csrfToken := uuid.New().String()
		c.SetCookie("csrf_token", csrfToken, 3600, "/", "", true, false)

		c.SetCookie("access_token", accessToken, 900, "/", "", true, true)
		c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", true, true)

		c.JSON(http.StatusOK, LoginResponse{AccessToken: accessToken, CSRFToken: csrfToken})

	}
}

func Refresh(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
			return
		}

		claims, valid := ValidateRefreshToken(refreshToken)
		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
			return
		}

		userIdStr, exists := claims["user_id"].(string)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token claims"})
			return
		}

		var userId pgtype.UUID
		err = userId.Scan(userIdStr)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
			return
		}

		newAccess, err := GenerateAccessToken(userId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not generate access token"})
			return
		}

		c.SetCookie("access_token", newAccess, 900, "/", "", true, true)
	}
}

func Logout(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("access_token", "", -1, "/", "", true, true)
		c.SetCookie("refresh_token", "", -1, "/", "", true, true)
		c.JSON(http.StatusOK, gin.H{"message": "logged out"})
	}
}
