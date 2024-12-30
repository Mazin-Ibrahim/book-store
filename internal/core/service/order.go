package service

import (
	"context"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/Mazin-Ibrahim/book-store/internal/core/port"
)

type OrderService struct {
	repo port.OrderRepository
}

func NewOrderService(repo port.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (os *OrderService) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	return nil, nil
}

func (os *OrderService) GetOrder(ctx context.Context, id int64) (*domain.Order, error) {
	return nil, nil
}
func (os *OrderService) ListsOrder(ctx context.Context, skip, limit int64) ([]domain.Order, error) {
	return nil, nil
}
