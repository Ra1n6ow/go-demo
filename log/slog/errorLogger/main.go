package main

import (
	"context"
	"github.com/mdobak/go-xerrors"
	"log/slog"
	"os"
	"path/filepath"
)

type stackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	switch a.Value.Kind() {
	case slog.KindAny:
		switch v := a.Value.Any().(type) {
		case error:
			a.Value = fmtErr(v)
		}
	}
	return a
}

// marshalStack extracts stack frames from the error
func marshalStack(err error) []stackFrame {
	trace := xerrors.StackTrace(err)
	if len(trace) == 0 {
		return nil
	}
	frames := trace.Frames()
	s := make([]stackFrame, len(frames))
	for i, v := range frames {
		f := stackFrame{
			Source: filepath.Join(
				filepath.Base(filepath.Dir(v.File)),
				filepath.Base(v.File),
			),
			Func: filepath.Base(v.Function),
			Line: v.Line,
		}
		s[i] = f
	}
	return s
}

// fmtErr returns a slog.Value with keys `msg` and `trace`. If the error
// does not implement interface { StackTrace() errors.StackTrace }, the `trace`
// key is omitted.
func fmtErr(err error) slog.Value {
	var groupValues []slog.Attr
	groupValues = append(groupValues, slog.String("msg", err.Error()))
	frames := marshalStack(err)
	if frames != nil {
		groupValues = append(groupValues,
			slog.Any("trace", frames),
		)
	}
	return slog.GroupValue(groupValues...)
}

func main() {
	/*  打印错误信息
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	err := errors.New("something happened")
	ctx := context.Background()
	logger.ErrorContext(ctx, "upload failed", slog.Any("error", err))
	*/

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
	})
	logger := slog.New(h)
	ctx := context.Background()
	err := xerrors.New("something happened")
	logger.ErrorContext(ctx, "image uploaded", slog.Any("error", err))
}
