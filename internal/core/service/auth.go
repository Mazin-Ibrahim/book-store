package service

import (
	"context"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/Mazin-Ibrahim/book-store/internal/core/port"
	"github.com/Mazin-Ibrahim/book-store/internal/core/util"
)

type AuthService struct {
	repo         port.UserRepository
	tokenService port.TokenService
}

func NewAuthService(repo port.UserRepository, tokenService port.TokenService) *AuthService {
	return &AuthService{
		repo:         repo,
		tokenService: tokenService,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return "", domain.ErrInvalidCredentials
		}
		return "", domain.ErrInternal
	}

	err = util.ComparePassword(password, user.Password)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	accessToken, err := as.tokenService.CreateToken(user)
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	return accessToken, nil
}
