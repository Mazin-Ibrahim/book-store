package port

import (
	"context"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
)

type UserRepository interface {
	// CreateUser inserts a new user into the database
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUserById selects a User by id
	GetUserById(ctx context.Context, id int64) (*domain.User, error)
	// GetUserByEmai selects a User by email
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	// ListUsers selects a list of Users with pagination
	ListUsers(ctx context.Context, skip, limit int64) ([]domain.User, error)
	// UpdateUser updates a User
	UpdateUser(ctx context.Context, User *domain.User) (*domain.User, error)
	// DeleteUser deletes a User
	DeleteUser(ctx context.Context, id int64) error
}

type UserService interface {
	// Register registers a new user
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUser returns a user by id
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	// ListUsers returns a list of users with pagination
	ListUsers(ctx context.Context, skip, limit int64) ([]domain.User, error)
	// UpdateUser updates a user
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser deletes a user
	DeleteUser(ctx context.Context, id int64) error
}
