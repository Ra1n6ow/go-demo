package log

import (
	"log/slog"
)

const (
	LevelPanic = slog.Level(12)
	LevelFatal = slog.Level(16)
)

var LevelNames = map[slog.Leveler]string{
	LevelPanic: "PANIC",
	LevelFatal: "FATAL",
}

func levelReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	//fmt.Println("-->", a.Key, slog.LevelKey, "<--")
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		levelLabel, exists := LevelNames[level]
		if !exists {
			levelLabel = level.String()
		}
		a.Value = slog.StringValue(levelLabel)
	}
	return a
}
