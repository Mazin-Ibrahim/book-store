package logger

import (
	"log/slog"
	"os"

	"github.com/Mazin-Ibrahim/book-store/internal/adapter/config"
	slogmulti "github.com/samber/slog-multi"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *slog.Logger

func Set(config *config.App) {

	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	if config.Env == "development" {
		logRotate := &lumberjack.Logger{
			Filename:   "log/app.log",
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
			Compress:   true,
		}

		logger = slog.New(
			slogmulti.Fanout(
				slog.NewJSONHandler(logRotate, nil),
				slog.NewTextHandler(os.Stderr, nil),
			),
		)
	}
	slog.SetDefault(logger)

}
