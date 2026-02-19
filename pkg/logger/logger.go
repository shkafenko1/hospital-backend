package logger

import (
	"log/slog"
	"os"
)

func InitLogger() {
	h := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(h)
	slog.SetDefault(logger)
}
