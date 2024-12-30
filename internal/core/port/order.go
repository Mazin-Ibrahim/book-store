package port

import (
	"context"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderById(ctx context.Context, id int64) (*domain.Order, error)
	OrderLists(ctx context.Context, skip, limit int64) ([]domain.Order, error)
}

type OrderService interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrder(ctx context.Context, id int64) (*domain.Order, error)
	OrderLists(ctx context.Context, skip, limit int64) ([]domain.Order, error)
}
