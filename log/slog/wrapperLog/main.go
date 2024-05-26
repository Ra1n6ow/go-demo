package main

import (
	"github.com/ra1n6ow/go-demo/log/slog/wrapperLog/log"
	"log/slog"
)

func logOptions() *log.Options {
	return &log.Options{
		AddSource: false,
		Level:     slog.LevelDebug,
	}
}

func main() {
	log.Init(logOptions())
	log.Info("Info message", slog.String("path", "/hello/world?q=search"))
	log.Error("Error message")
}
