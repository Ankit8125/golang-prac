package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims // Hover to see all it's properties
	Role string `json:"role"`
}

func CreateToken (jwtSecret string, userID string, role string) (string, error) {
	now := time.Now().UTC()
	exp := now.Add(7*24*time.Hour)

	claims := Claims {
		// It is important to pass subject (userID) and role (can be user/admin/superadmin).
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userID, // unique identifier and this will become encrypted
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Role: role,
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // NewWithClaims creates a new Token with the specified signing method and claims.

	signed, err := tok.SignedString([] byte (jwtSecret))
	if err != nil {
		return "", fmt.Errorf("Sign in token failed: %w", err)
	}

	return signed, nil
}