package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":          email,
		"UserID":         userId,
		"ExpirationTime": 1000,
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return 0, err
	}

	isValid := parsedToken.Valid

	if !isValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token")
	}

	// email := claims["Email"].(string)
	userId := int64(claims["UserID"].(float64))
	// userId := claims["UserID"].(int64)

	return userId, nil
}
