package domain

import "github.com/golang-jwt/jwt/v5"

type TokenPayload struct {
	*jwt.Token
}
