package http

import (
	"net/http"
	"time"

	"github.com/Mazin-Ibrahim/book-store/internal/adapter/config"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	*chi.Mux
}

func NewRouter(config *config.HTTP, bookHandler BookHandler, userHandler UserHandler, authHandler AuthHandler) (*Router, error) {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Route("/v1", func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Post("/create", bookHandler.CreateBook)
			r.Get("/", bookHandler.ListBooks)
			r.Get("/{id}", bookHandler.GetBookById)
			r.Delete("/{id}", bookHandler.DeleteBook)
			r.Put("/{id}", bookHandler.UpdateBook)
		})
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", userHandler.RegisterUser)
			r.Put("/{id}/update", userHandler.UpdateUser)
			r.Get("/", userHandler.ListUsers)
			r.Get("/{id}", userHandler.GetUser)
		})
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authHandler.Login)
		})
	})

	return &Router{
		Mux: router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	srv := http.Server{
		Addr:         listenAddr,
		Handler:      r,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	return srv.ListenAndServe()
}
