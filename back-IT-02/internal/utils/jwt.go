package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var SecretKey []byte

// func InitJWT() {
// 	SecretKey = []byte(os.Getenv("JWT_SECRET"))
// }

func GenerateToken(userID uint) (string, error) {

	fmt.Println("GENERATE SECRET =", os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	secret := []byte(os.Getenv("JWT_SECRET"))

	return token.SignedString(secret)
}

func ValidateToken(tokenString string) (uint, error) {
	fmt.Println("GENERATE SECRET =", os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("JWT_SECRET =", os.Getenv("JWT_SECRET"))
		fmt.Println("TOKEN =", tokenString)
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	if err != nil {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)

	userID := uint(claims["user_id"].(float64))

	return userID, nil
}
