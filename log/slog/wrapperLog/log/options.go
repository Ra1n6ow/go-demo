package log

import "log/slog"

// Options 包含与日志相关的配置项.
type Options struct {
	// 如果开启会在日志中显示调用日志所在的文件和行号
	AddSource bool
	Level     slog.Level
	// 指定日志显示格式，可选值：text, json
	Format string
	// 指定日志输出位置
	OutputPaths []string
}

// NewOptions 创建一个带有默认参数的 Options 对象.
func NewOptions() *Options {
	return &Options{
		AddSource:   false,
		Level:       slog.LevelInfo,
		Format:      "text",
		OutputPaths: []string{"stdout"},
	}
}
