package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var jwtAccessSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))
var jwtRefreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

func GenerateAccessToken(userId pgtype.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtAccessSecret)
}

func GenerateRefreshToken(userId pgtype.UUID) (string, error) {
	claims := jwt.MapClaims{"user_id": userId, "exp": time.Now().Add(7 * 24 * time.Hour).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtRefreshSecret)
}

func ValidateAccessToken(tokenString string) (jwt.MapClaims, bool) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtAccessSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, false
	}

	return token.Claims.(jwt.MapClaims), true
}

func ValidateRefreshToken(tokenString string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtRefreshSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, false
	}

	return token.Claims.(jwt.MapClaims), true

}
