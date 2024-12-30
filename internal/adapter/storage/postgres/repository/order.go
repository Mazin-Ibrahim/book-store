package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Mazin-Ibrahim/book-store/internal/adapter/storage/postgres"
	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

type OrderRepository struct {
	db *postgres.DB
}

func NewOrderReposiotory(db *postgres.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := or.db.QueryBuilder.Insert("orders").Columns("user_id", "book_id").Values(&order.UserId, &order.BookId).Suffix("RETURNING id,user_id,book_id")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = or.db.QueryRow(ctx, sql, args...).Scan(&order.ID, &order.UserId, &order.BookId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (or *OrderRepository) GetOrderById(ctx context.Context, id int64) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := or.db.QueryBuilder.Select("id,user_id,book_id").From("orders").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var order domain.Order
	err = or.db.QueryRow(ctx, sql, args...).Scan(&order.ID, &order.UserId, &order.BookId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}
	return &order, nil
}
func (or *OrderRepository) OrderLists(ctx context.Context, skip, limit int64) ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	query := or.db.QueryBuilder.Select("id,user_id,book_id").Options("id").Limit(uint64(limit))

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := or.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ordersList []domain.Order
	var order domain.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.UserId, &order.BookId); err != nil {
			return nil, err
		}
		ordersList = append(ordersList, order)
	}
	return ordersList, nil
}
