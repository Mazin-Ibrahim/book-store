package http

import "github.com/Mazin-Ibrahim/book-store/internal/core/domain"

type bookResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Cover       string  `json:"cover"`
}

func newBookResponse(book *domain.Book) bookResponse {
	return bookResponse{
		ID:          book.ID,
		Name:        book.Name,
		Author:      book.Author,
		Description: book.Description,
		Cover:       book.Cover,
		Price:       book.Price,
	}
}

type userResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func newUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
