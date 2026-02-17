package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(UserID uint) (string, error) {
	key := getJWTKey()

	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)

}

func ParseToken(tokenString string) error {
	key := getJWTKey()

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil || !token.Valid {
		return fmt.Errorf("некорректный токен")
	}
	return nil
}
