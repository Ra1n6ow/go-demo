package log

import (
	"context"
	"log/slog"
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
	Panic(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}

// Init 使用指定的选项初始化 Logger.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

func NewLogger(opts *Options) *sLogger {

	//hops := &slog.HandlerOptions{
	//	Level:       opts.Level,
	//	AddSource:   opts.AddSource,
	//	ReplaceAttr: levelReplaceAttr,
	//}
	//h := slog.NewJSONHandler(os.Stdout, hops)
	h := NewHandler(&slog.HandlerOptions{
		Level:       opts.Level,
		AddSource:   opts.AddSource,
		ReplaceAttr: nil,
	})

	s := slog.New(h)

	logger := &sLogger{s}
	return logger
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

func Panic(msg string, keysAndValues ...interface{}) {
	std.s.Log(context.Background(), LevelPanic, msg, keysAndValues...)
}

func (l *sLogger) Panic(msg string, keysAndValues ...interface{}) {
	l.s.Log(context.Background(), LevelPanic, msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	std.s.Log(context.Background(), LevelFatal, msg, keysAndValues...)
}

func (l *sLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.s.Log(context.Background(), LevelFatal, msg, keysAndValues...)
}
