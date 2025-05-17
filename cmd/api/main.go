package main

import (
	"log/slog"
	"os"
)

const appVersion = "0.0.1-alpha"

func main() {
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	defaultLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(defaultLogger)

	slog.Info(
		"RedToGreen API iniciando...",
		slog.String("version", appVersion),
		slog.String("log_level", slog.LevelInfo.String()),
	)

	return nil
}
