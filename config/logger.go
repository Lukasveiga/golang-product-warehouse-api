package config

import (
	"os"

	"golang.org/x/exp/slog"
)

func LoggerConfig(env string) {
	level := new(slog.LevelVar)

	switch env {
	case "prod":
		level.Set(slog.LevelInfo.Level())
	default:
		level.Set(slog.LevelDebug.Level())
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(l)
}