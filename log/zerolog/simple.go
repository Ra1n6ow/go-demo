package zerolog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// 基础示例
func Sample() {
	// 默认时间格式：{"level":"debug","time":"2024-03-18T15:44:06+08:00","message":"hello world"}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Note: By default log writes to os.Stderr
	// Note: The default log level for log.Print is trace
	log.Print("hello world")
	// Output: {"time":1516134303,"level":"debug","message":"hello world"}
}

// 添加上下文信息
func Contextual() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
	// {"level":"debug","Scale":"833 cents","Interval":833.09,"time":1710747743,"message":"Fibonacci is everywhere"}

	// 日志等级
	log.Info().
		Str("Name", "Tom").
		Send()
	// {"level":"info","Name":"Tom","time":1710747743}
}
