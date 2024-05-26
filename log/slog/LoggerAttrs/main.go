package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)

	logger.Warn(
		"incoming request",
		"method", "GET",
		"time_taken_ms", // 缺少key，会以 !BADKEY 替代
	)

	// 强类型上下文
	logger.Info(
		"incoming request",
		"method", "GET", //但这种写法也不会报错
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)

	// 最安全的上下文
	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"incoming request",
		//"method", "GET", // 这种写法会报错
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)

	// 分组上下文属性
	logger1 := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger1.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"image uploaded",
		slog.Int("id", 23123),
		//  properties.width=4000 properties.height=3000 properties.format=jpeg
		slog.Group("properties",
			slog.Int("width", 4000),
			slog.Int("height", 3000),
			slog.String("format", "jpeg"),
		),
	)

	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"image uploaded",
		slog.Int("id", 23123),
		// "properties":{"width":4000,"height":3000,"format":"jpeg"}
		slog.Group("properties",
			slog.Int("width", 4000),
			slog.Int("height", 3000),
			slog.String("format", "jpeg"),
		),
	)
}
