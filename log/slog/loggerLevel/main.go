package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	logLevel := &slog.LevelVar{} // 默认为0，即 Info
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	logLevel.Set(slog.LevelDebug) // 设置默认为 Debug
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	/*
		创建自定义日志级别
	*/
	const (
		// Level 类型实现了 Leveler 接口
		LevelTrace = slog.Level(-8)
		LevelFatal = slog.Level(12)
	)

	opts1 := &slog.HandlerOptions{
		Level: LevelTrace,
	}
	logger1 := slog.New(slog.NewJSONHandler(os.Stdout, opts1))
	ctx := context.Background()
	logger1.Log(ctx, LevelTrace, "Trace message") // 这里的等级为 DEBUG-4
	logger1.Log(ctx, LevelFatal, "Fatal level")   // 这里的等级为 ERROR+4

	var LevelNames = map[slog.Leveler]string{
		LevelTrace: "TRACE",
		LevelFatal: "FATAL",
	}

	opts2 := &slog.HandlerOptions{
		Level: LevelTrace,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}
	logger2 := slog.New(slog.NewJSONHandler(os.Stdout, opts2))
	logger2.Log(ctx, LevelTrace, "Trace message") // 这里的等级为 DEBUG-4
	logger2.Log(ctx, LevelFatal, "Fatal level")   // 这里的等级为 ERROR+4
}
