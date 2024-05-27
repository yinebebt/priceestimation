package utils

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// IsPasswordValid validate whether the password is enough strong
func IsPasswordValid(password string) bool {
	if !strings.ContainsAny(password, "0123456789") || !strings.ContainsAny(password, "!@#$%^&*_/~]") {
		return false
	}

	re, err := regexp.Compile("[A-Za-z]")
	if err != nil {
		fmt.Println("error occurred ", err)
		return false
	}
	if !re.Match([]byte(password)) {
		return false
	}
	return true
}
