package http

import (
	"net/http"

	"github.com/Mazin-Ibrahim/book-store/internal/core/domain"
	"github.com/Mazin-Ibrahim/book-store/internal/core/port"
)

type UserHandler struct {
	service port.UserService
}

func NewUserHandler(service port.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

type registerRequestUser struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100,min=3"`
}

func (us *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload registerRequestUser
	if err := readJSON(w, r, &payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}
	err := Validate.Struct(payload)
	if err != nil {
		validationsMessages, err := validationErrors(err)
		if err != nil {
			internalServerError(w, r, err)
			return
		}
		badRequestResponseWithTags(w, r, validationsMessages)
		return
	}
	user := domain.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	_, err = us.service.Register(r.Context(), &user)
	if err != nil {
		switch err {
		case domain.ErrConflictingData:
			jsonResponse(w, http.StatusConflict, "email used before")
			return
		default:
			internalServerError(w, r, err)
			return
		}
	}
	if err = jsonResponse(w, http.StatusOK, newUserResponse(&user)); err != nil {
		internalServerError(w, r, err)
		return
	}
}

type updateRequestUser struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100,min=3"`
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var payload updateRequestUser
	if err := readJSON(w, r, &payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		messages, err := validationErrors(err)
		if err != nil {
			internalServerError(w, r, err)
			return
		}
		badRequestResponseWithTags(w, r, messages)
		return
	}

	id, err := extractID(r)
	if err != nil {
		internalServerError(w, r, err)
		return
	}
	user := &domain.User{
		ID:       id,
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	_, err = uh.service.UpdateUser(r.Context(), user)
	if err != nil {
		switch err {
		case domain.ErrConflictingData:
			err := jsonResponse(w, http.StatusConflict, "email used before")
			if err != nil {
				internalServerError(w, r, err)
				return
			}
			return
		default:
			internalServerError(w, r, err)
			return
		}
	}
	if err := jsonResponse(w, http.StatusOK, newUserResponse(user)); err != nil {
		internalServerError(w, r, err)
		return
	}
}

func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {

	users, err := uh.service.ListUsers(r.Context(), 0, 10)

	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			if err := jsonResponse(w, http.StatusOK, err); err != nil {
				internalServerError(w, r, err)
				return
			}
			return
		default:
			internalServerError(w, r, err)
			return
		}
	}

	var usersList []userResponse
	for _, user := range users {
		usersList = append(usersList, newUserResponse(&user))
	}
	if err := jsonResponse(w, http.StatusOK, usersList); err != nil {
		internalServerError(w, r, err)
		return
	}
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := extractID(r)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	user, err := uh.service.GetUser(r.Context(), id)
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

	if err := jsonResponse(w, http.StatusOK, newUserResponse(user)); err != nil {
		internalServerError(w, r, err)
		return
	}
}
