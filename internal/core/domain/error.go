package domain

import "errors"

var (
	ErrDataNotFound       = errors.New("data not found")
	ErrConflictingData    = errors.New("data conflicts with existing data in unique column")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInternal           = errors.New("internal error")
	ErrTokenCreation      = errors.New("error creating token")
)
