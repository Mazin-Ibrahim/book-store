package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Mazin-Ibrahim/book-store/internal/adapter/storage/postgres"
	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

var QueryTimeOutDuration = time.Second * 5

type BookRepository struct {
	db *postgres.DB
}

func NewBookRepository(db *postgres.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := br.db.QueryBuilder.Insert("books").
		Columns("name", "author", "price", "description", "cover").
		Values(book.Name, book.Author, book.Price, book.Description, book.Cover).
		Suffix("RETURNING id,name,author,price,description,cover")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = br.db.QueryRow(ctx, sql, args...).Scan(
		&book.ID,
		&book.Name,
		&book.Author,
		&book.Price,
		&book.Description,
		&book.Cover,
	)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (br *BookRepository) GetBookById(ctx context.Context, id int64) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	var book domain.Book

	query := br.db.QueryBuilder.Select("id,name,author,price,description,cover").From("books").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = br.db.QueryRow(ctx, sql, args...).Scan(
		&book.ID,
		&book.Name,
		&book.Author,
		&book.Price,
		&book.Description,
		&book.Cover,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}
	return &book, nil
}

func (br *BookRepository) ListBooks(ctx context.Context, skip, limit int64) ([]domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := br.db.QueryBuilder.Select(" id,name,author,price,description,cover").From("books").OrderBy("id").Limit(uint64(limit))
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := br.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var book domain.Book
	var books []domain.Book
	for rows.Next() {
		err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&book.Price,
			&book.Description,
			&book.Cover,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (br *BookRepository) UpdateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := br.db.QueryBuilder.Update("books").
		Set("name", book.Name).
		Set("author", book.Author).
		Set("price", book.Price).
		Set("description", book.Description).
		Set("cover", book.Cover).
		Where(sq.Eq{"id": book.ID}).
		Suffix("RETURNING id,name,author,price,description,cover")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = br.db.QueryRow(ctx, sql, args...).Scan(
		&book.ID,
		&book.Name,
		&book.Author,
		&book.Price,
		&book.Description,
		&book.Cover,
	)
	if err != nil {
		if errCode := br.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}
	return book, nil
}

func (br *BookRepository) DeleteBook(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := br.db.QueryBuilder.Delete("books").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = br.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
