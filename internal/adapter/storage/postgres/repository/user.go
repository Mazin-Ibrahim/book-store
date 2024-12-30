package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Mazin-Ibrahim/book-store/internal/adapter/storage/postgres"
	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *postgres.DB
}

// GetUserById implements port.UserRepository.
func (ur *UserRepository) GetUserById(ctx context.Context, id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	var user domain.User

	query := ur.db.QueryBuilder.Select("id,name,email").From("users").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := ur.db.QueryBuilder.Insert("users").
		Columns("name", "email", "password").
		Values(user.Name, user.Email, user.Password).Suffix("RETURNING id,name,email")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errCode := ur.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return user, nil
}

// GetUserByEmailAndPassword gets a user by email from the database
func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()
	query := ur.db.QueryBuilder.Select("id,name,email,password").From("users").Where(sq.Eq{"email": email})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}
	var user domain.User
	err = ur.db.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &user, nil
}

// ListUsers lists all users from the database
func (ur *UserRepository) ListUsers(ctx context.Context, skip, limit int64) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()
	query := ur.db.QueryBuilder.Select("id,name,email").From("users").OrderBy("id").Limit(uint64(limit))

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := ur.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var usersList []domain.User
	var user domain.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}
	return usersList, nil
}

// UpdateUser updates a user by ID in the database
func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := ur.db.QueryBuilder.Update("users").
		Set("name", user.Name).
		Set("email", user.Email).
		Set("password", user.Password).
		Where(sq.Eq{"id": user.ID}).
		Suffix("RETURNING id,name,email")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errCode := ur.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID from the database
func (ur *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := ur.db.QueryBuilder.Delete("users").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = ur.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
