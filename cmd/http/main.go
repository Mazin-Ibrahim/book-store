package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Mazin-Ibrahim/book-store/internal/adapter/config"
	"github.com/Mazin-Ibrahim/book-store/internal/adapter/handler/http"
	"github.com/Mazin-Ibrahim/book-store/internal/adapter/logger"
	"github.com/Mazin-Ibrahim/book-store/internal/adapter/storage/postgres"
	"github.com/Mazin-Ibrahim/book-store/internal/adapter/storage/postgres/repository"
	"github.com/Mazin-Ibrahim/book-store/internal/core/service"
)

func main() {

	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}
	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := http.NewBookHandler(bookService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	tokenService := service.TokenService{}
	authService := service.NewAuthService(userRepo, &tokenService)
	authHandler := http.NewAuthHandler(authService)

	router, err := http.NewRouter(config.HTTP, *bookHandler, *userHandler, *authHandler)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}

}
