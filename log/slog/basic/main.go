package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Print("Info message")
	slog.Info("Info message")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Info message")

	logger1 := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger1.Info("Info message")
}
