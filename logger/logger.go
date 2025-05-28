package logger

import (
	"log/slog"
	"os"
)

func Log() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Debug("test", "test", "test")
}
