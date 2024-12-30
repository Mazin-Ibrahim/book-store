package port

import (
	"context"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
)

// BookRepository is an interface for interacting with book-related data
type BookRepository interface {
	// CreateBook inserts a new book into the database
	CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, error)
	// GetBookById selects a book by id
	GetBookById(ctx context.Context, id int64) (*domain.Book, error)

	// ListBooks selects a list of books with pagination
	ListBooks(ctx context.Context, skip, limit int64) ([]domain.Book, error)
	// UpdateBook updates a book
	UpdateBook(ctx context.Context, book *domain.Book) (*domain.Book, error)
	// DeleteBook deletes a book
	DeleteBook(ctx context.Context, id int64) error
}

// BookService is an interface for interacting with book-related business logic
type BookService interface {
	// CreateBook creates a new book
	CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, error)
	// GetBook returns a book by id
	GetBook(ctx context.Context, id int64) (*domain.Book, error)
	// ListBooks returns a list of books with pagination
	ListBooks(ctx context.Context, skip, limit int64) ([]domain.Book, error)
	// UpdateBook updates a book
	UpdateBook(ctx context.Context, book *domain.Book) (*domain.Book, error)
	// DeleteBook deletes a book
	DeleteBook(ctx context.Context, id int64) error
}
