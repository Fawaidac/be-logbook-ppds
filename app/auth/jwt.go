package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	ProgramStudi string `json:"program_studi"`
	Role         string `json:"role"`
	Nim		  string `json:"nim"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, username, name, programStudi, nim, role, secret string) (string, error) {
	claims := JWTClaims{
		UserID:       userID,
		Username:     username,
		Name:         name,
		ProgramStudi: programStudi,
		Nim:          nim,
		Role:         role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenStr, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
