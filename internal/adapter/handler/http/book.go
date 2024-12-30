package http

import (
	"net/http"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/Mazin-Ibrahim/book-store/internal/core/port"
)

type BookHandler struct {
	service port.BookService
}

func NewBookHandler(service port.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

type createBookRequest struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description" validate:"required,max=1000"`
	Author      string  `json:"author" validate:"required,max=50"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Cover       string  `json:"cover"`
}

func (bh *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var payload createBookRequest

	if err := readJSON(w, r, &payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		validationErrors, err := validationErrors(err)
		if err != nil {
			internalServerError(w, r, err)
			return
		}
		badRequestResponseWithTags(w, r, validationErrors)
		return
	}
	book := domain.Book{
		Name:        payload.Name,
		Author:      payload.Author,
		Description: payload.Description,
		Cover:       payload.Cover,
		Price:       payload.Price,
	}
	_, err := bh.service.CreateBook(r.Context(), &book)
	if err != nil {
		internalServerError(w, r, err)
		return
	}
	if err := jsonResponse(w, http.StatusCreated, newBookResponse(&book)); err != nil {
		internalServerError(w, r, err)
		return
	}
}

func (bh *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {

	books, err := bh.service.ListBooks(r.Context(), 0, 20)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	var booksList []bookResponse

	for _, book := range books {
		booksList = append(booksList, newBookResponse(&book))
	}

	if err := jsonResponse(w, http.StatusOK, booksList); err != nil {
		internalServerError(w, r, err)
		return
	}

}

func (bh *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r)
	if err != nil {
		internalServerError(w, r, err)
		return
	}
	book, err := bh.service.GetBook(r.Context(), id)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			badRequestResponse(w, r, err)
			return
		default:
			internalServerError(w, r, err)
			return
		}
	}
	if err := jsonResponse(w, http.StatusOK, bookResponse(*book)); err != nil {
		internalServerError(w, r, err)
	}
}
func (bh *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	book, err := bh.service.GetBook(r.Context(), id)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	if book == nil {
		badRequestResponse(w, r, nil)
	}

	if err = bh.service.DeleteBook(r.Context(), id); err != nil {
		internalServerError(w, r, err)
		return
	}
	if err = jsonResponse(w, http.StatusNoContent, nil); err != nil {
		internalServerError(w, r, err)
		return
	}
}

type updateBookRequest struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description" validate:"required,max=1000"`
	Author      string  `json:"author" validate:"required,max=50"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Cover       string  `json:"cover"`
}

func (bh *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var payload updateBookRequest
	if err := readJSON(w, r, &payload); err != nil {
		internalServerError(w, r, err)
	}

	if err := Validate.Struct(payload); err != nil {
		validationErrors, err := validationErrors(err)
		if err != nil {
			internalServerError(w, r, err)
			return
		}
		badRequestResponseWithTags(w, r, validationErrors)
		return
	}
	id, err := extractID(r)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	book := domain.Book{
		ID:          id,
		Name:        payload.Name,
		Author:      payload.Author,
		Description: payload.Description,
		Cover:       payload.Cover,
		Price:       payload.Price,
	}
	_, err = bh.service.UpdateBook(r.Context(), &book)
	if err != nil {
		switch err {
		case domain.ErrConflictingData:
			if err := jsonResponse(w, http.StatusConflict, err); err != nil {
				internalServerError(w, r, err)
				return
			}
			return
		default:
			internalServerError(w, r, err)
			return
		}
	}
	if err = jsonResponse(w, http.StatusOK, newBookResponse(&book)); err != nil {
		internalServerError(w, r, err)
		return
	}
}
