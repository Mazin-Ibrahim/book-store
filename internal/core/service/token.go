package service

import (
	"fmt"
	"os"
	"time"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
}

// CreateToken creates a new token for a given user
func (ts *TokenService) CreateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(3 * 24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": os.Getenv("JWT_ISS"),
		"aud": os.Getenv("JWT_AUD"),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken verifies the token and returns the payload
func (ts *TokenService) VerifyToken(token string) (*domain.TokenPayload, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(os.Getenv("JWT_AUD")),
		jwt.WithIssuer(os.Getenv("JWT_ISS")),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		return nil, err
	}
	tokenPayload := domain.TokenPayload{
		Token: jwtToken,
	}

	return &tokenPayload, nil
}
