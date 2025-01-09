package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), error
}

func CheckPasswordHash(password, hash string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return error == nil
}
