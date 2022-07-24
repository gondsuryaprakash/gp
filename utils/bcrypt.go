package utils

import (
	"github.com/gondsuryaprakash/gondpariwar/logger"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	funcName := "utils.HashPassword"
	logger.I(funcName)

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.D("string(bytePassword)", string(bytePassword))
	return string(bytePassword), err
}
func CheckPasswordHash(hashedPassword, password string) bool {
	funcName := "utils.CheckPasswordHash"
	logger.D("hashedPassword", hashedPassword)
	logger.D("password", password)
	logger.I(funcName)
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		logger.E(funcName, err)
		return false
	}
	return true
}
