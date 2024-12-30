package service

import (
	"context"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/Mazin-Ibrahim/book-store/internal/core/port"
)

type BookService struct {
	repo port.BookRepository
}

func NewBookService(repo port.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (bs *BookService) CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {

	book, err := bs.repo.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (bs *BookService) GetBook(ctx context.Context, id int64) (*domain.Book, error) {
	book, err := bs.repo.GetBookById(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (bs *BookService) ListBooks(ctx context.Context, skip, limt int64) ([]domain.Book, error) {

	books, err := bs.repo.ListBooks(ctx, skip, limt)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (bs *BookService) UpdateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	book, err := bs.repo.UpdateBook(ctx, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (bs *BookService) DeleteBook(ctx context.Context, id int64) error {
	err := bs.repo.DeleteBook(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
