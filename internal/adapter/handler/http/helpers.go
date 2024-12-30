package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func extractID(r *http.Request) (int64, error) {
	value := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, err
}
