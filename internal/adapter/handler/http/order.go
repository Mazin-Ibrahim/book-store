package http

import (
	"net/http"

	"github.com/Mazin-Ibrahim/book-store/internal/core/port"
)

type OrderHandler struct {
	service port.OrderService
}

func NewOrderService(service port.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

type createOrderRequest struct {
	BookId int64 `json:"book_id"`
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var payload createBookRequest
	err := readJSON(w, r, &payload)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
}
func (oh *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request)   {}
func (oh *OrderHandler) ListsOrder(w http.ResponseWriter, r *http.Request) {}
