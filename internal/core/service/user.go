package service

import (
	"context"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/Mazin-Ibrahim/book-store/internal/core/port"
	"github.com/Mazin-Ibrahim/book-store/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

// Register creates a new user
func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser gets a user by ID
func (us *UserService) GetUser(ctx context.Context, id int64) (*domain.User, error) {

	user, err := us.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil

}

// ListUsers lists all users
func (us *UserService) ListUsers(ctx context.Context, skip, limit int64) ([]domain.User, error) {

	users, err := us.repo.ListUsers(ctx, skip, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates a user's name, email, and password
func (us *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	user, err = us.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user by ID
func (us *UserService) DeleteUser(ctx context.Context, id int64) error {
	if err := us.repo.DeleteUser(ctx, id); err != nil {
		return err
	}
	return nil
}
