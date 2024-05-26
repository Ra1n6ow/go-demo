package main

import (
	"log/slog"
	"os"
	"runtime/debug"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	buildInfo, _ := debug.ReadBuildInfo()
	logger := slog.New(handler)
	child := logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)
	child.Info("image upload successful", slog.String("image_id", "39ud88"))
	child.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 mb"),
	)

	logger1 := slog.New(handler).WithGroup("program_info")
	child1 := logger1.With(
		slog.Int("pid", os.Getpid()),
		slog.String("go_version", buildInfo.GoVersion),
	)
	child1.Warn(
		"storage is 95% full",
		slog.String("available_space1111", "11.1 MB"),
	)
}
