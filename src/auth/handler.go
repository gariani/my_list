package auth

import (
	"net/http"

	"github.com/gariani/my_list/src/utils"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := RegisterUser(input.Email, input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})

}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := GetUserByEmail(input.Email)
	if err != nil || !utils.CheckPassword(user.PassHash, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	accessToken, err := GenerateAccessToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token error"})
		return
	}

	refreshToken, err := GenerateRefreshToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token error"})
		return
	}
	c.SetCookie("access_token", accessToken, 900, "/", "", true, true)
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "logged in"})

}

func Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	claims, valid := ValidateRefreshToken(refreshToken)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token claims"})
		return
	}

	newAccess, err := GenerateAccessToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate access token"})
		return
	}

	c.SetCookie("access_token", newAccess, 900, "/", "", true, true)
}

func Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
