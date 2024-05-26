package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// 设置默认logger
	slog.SetDefault(logger)
	slog.Info("Info message")
	// 还会改变 log 包默认使用的 log.Logger
	log.Println("Hello from old logger")

	/*
		将 slog.Logger 转换为 log.Logger
	*/
	handler1 := slog.NewJSONHandler(os.Stdout, nil)
	logger1 := slog.NewLogLogger(handler1, slog.LevelError)
	_ = http.Server{
		// this API only accepts `log.Logger`
		ErrorLog: logger1,
	}
}
