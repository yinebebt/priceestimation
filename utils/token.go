package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
)

// GenerateToken generates token
func GenerateToken(userid uuid.UUID) string {
	extendedClaim := jwt.RegisteredClaims{
		Issuer:    "https://chipchip.social",
		Subject:   userid.String(),
		Audience:  []string{"https://chipchip.social"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        uuid.New().String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, extendedClaim)
	tkn, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("error occurred, ", err)
		return ""
	}
	return tkn
}

// ValidateToken validate the given token
func ValidateToken(token string) (*jwt.RegisteredClaims, bool) {
	claims := &jwt.RegisteredClaims{}
	segments := strings.Split(token, ".")
	if len(segments) < 3 {
		return claims, false
	}

	if _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("st-son"), nil
	}); err != nil {
		return claims, false
	}
	return claims, true
}
