package http

import (
	"log/slog"
	"net/http"
)

func internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONErorr(w, http.StatusInternalServerError, "the server encounterd a problem")
}

func forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	slog.Warn("forbidden", "method", r.Method, "path", r.URL.Path, "error")

	writeJSONErorr(w, http.StatusForbidden, "forbidden")
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	slog.Warn("bad request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONErorr(w, http.StatusBadRequest, err.Error())
}
func badRequestResponseWithTags(w http.ResponseWriter, r *http.Request, data []byte) {
	slog.Warn("bad request error", "method", r.Method, "path", r.URL.Path, "error", string(data))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)
}

func notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {

	slog.Warn("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONErorr(w, http.StatusNotFound, "not found")
}

func conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	slog.Warn("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONErorr(w, http.StatusConflict, err.Error())
}
func unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	slog.Warn("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONErorr(w, http.StatusUnauthorized, "unauthorized")
}

func unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	slog.Warn("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJSONErorr(w, http.StatusUnauthorized, "unauthorized")
}
