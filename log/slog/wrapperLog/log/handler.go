package log

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

type Handler struct {
	h slog.Handler
	m *sync.Mutex
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h: h.h.WithAttrs(attrs), m: h.m}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h: h.h.WithGroup(name), m: h.m}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	r.Time.Format("[15:04:05.000]")
	h.h.Handle(ctx, r)
	return nil
}

func NewHandler(opts *slog.HandlerOptions) *Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &Handler{
		h: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: opts.ReplaceAttr,
		}),
		m: &sync.Mutex{},
	}
}
