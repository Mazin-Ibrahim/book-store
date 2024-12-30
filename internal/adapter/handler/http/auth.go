package http

import (
	"net/http"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/Mazin-Ibrahim/book-store/internal/core/port"
)

type AuthHandler struct {
	authService port.AuthService
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type loginRequestPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=100,min=3"`
}

func (as *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload loginRequestPayload
	if err := readJSON(w, r, &payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	err := Validate.Struct(payload)

	if err != nil {
		messages, err := validationErrors(err)
		if err != nil {
			internalServerError(w, r, err)
			return
		}
		badRequestResponseWithTags(w, r, messages)
		return
	}

	token, err := as.authService.Login(r.Context(), payload.Email, payload.Password)
	if err != nil {
		switch err {
		case domain.ErrInvalidCredentials:
			if err := jsonResponse(w, http.StatusNotFound, err); err != nil {
				internalServerError(w, r, err)
				return
			}
			return
		case domain.ErrInternal:
			internalServerError(w, r, err)
			return
		case domain.ErrTokenCreation:
			internalServerError(w, r, err)
		}
	}

	if err := jsonResponse(w, http.StatusOK, token); err != nil {
		internalServerError(w, r, err)
		return
	}

}
