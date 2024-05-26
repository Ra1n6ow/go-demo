package main

import (
	"context"
	"log/slog"
	"os"
)

type ctxKey string

const (
	slogFields ctxKey = "slog_fields"
)

type ContextHandler struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying
// handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}
	return h.Handler.Handle(ctx, r)
}

// AppendCtx adds an slog attribute to the provided context so that it will be
// included in any Record created with such context
func AppendCtx(parent context.Context, attr slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	if v, ok := parent.Value(slogFields).([]slog.Attr); ok {
		v = append(v, attr)
		return context.WithValue(parent, slogFields, v)
	}
	v := []slog.Attr{}
	v = append(v, attr)
	return context.WithValue(parent, slogFields, v)
}

func main() {
	/* 这样不会在日志中输出 context 信息
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx := context.WithValue(context.Background(), "request_id", "req-123")
	logger.InfoContext(ctx, "image uploaded", slog.String("image_id", "img-998"))
	*/
	h := &ContextHandler{slog.NewJSONHandler(os.Stdout, nil)}
	logger := slog.New(h)
	ctx := AppendCtx(context.Background(), slog.String("request_id", "req-123"))
	logger.InfoContext(ctx, "image uploaded", slog.String("image_id", "img-998"))
}
