package main

import (
	"github.com/ra1n6ow/go-demo/log/slog/wrapperLog/log"
	"log/slog"
)

func logOptions() *log.Options {
	return &log.Options{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
}

func main() {
	log.Init(logOptions())
	log.Debug("Debug message", slog.String("KD", "VD"))
	log.Info("Info message", slog.String("KI", "VI"))
	log.Error("Error message")
	log.Panic("Panic message", slog.String("KP", "VP"))
	log.Fatal("Fatal message", slog.String("KF", "VF"))
}

//// prettylog
//func main() {
//	prettyHandler := prettylog.NewHandler(&slog.HandlerOptions{
//		Level:       slog.LevelInfo,
//		AddSource:   true,
//		ReplaceAttr: nil,
//	})
//	log := slog.New(prettyHandler)
//	log.Info("Info message", slog.String("KI", "VI"))
//}
