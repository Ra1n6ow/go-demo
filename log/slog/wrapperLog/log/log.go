package log

import (
	"log/slog"
	"os"
	"sync"
)

var (
	mu  sync.Mutex
	std = NewLogger(NewOptions())
)

type sLogger struct {
	s *slog.Logger
}

// 确保 sLogger 实现了 Logger 接口. 以下变量赋值，可以使错误在编译期被发现.
var _ Logger = &sLogger{}

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

// Init 使用指定的选项初始化 Logger.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

func NewLogger(opts *Options) *sLogger {
	hdOpts := &slog.HandlerOptions{
		AddSource: opts.AddSource, // 日志输出源
		Level:     opts.Level,
	}
	handler := slog.NewJSONHandler(os.Stdout, hdOpts)
	logger := slog.New(handler)
	return &sLogger{logger}
}

func Debug(msg string, keysAndValues ...interface{}) {
	std.s.Debug(msg, keysAndValues...)
}

func (l *sLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.s.Debug(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	std.s.Info(msg, keysAndValues...)
}

func (l *sLogger) Info(msg string, keysAndValues ...interface{}) {
	l.s.Info(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	std.s.Warn(msg, keysAndValues...)
}

func (l *sLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.s.Warn(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	std.s.Error(msg, keysAndValues...)
}

func (l *sLogger) Error(msg string, keysAndValues ...interface{}) {
	l.s.Error(msg, keysAndValues...)
}
