package main

import (
	"flag"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// 设置全局日志
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Debug().Msg("This message appears only when log level set to Debug")
	log.Info().Msg("This message appears when log level set to Debug or Info")

	if e := log.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}

	// Log()方法无视等级
	log.Log().Str("foo2", "bar2").Send()
}

/*
❯ go run main.go
{"level":"info","time":1710750651,"message":"This message appears when log level set to Debug or Info"}
{"foo2":"bar2","time":1710750651}

❯ go run main.go -debug
{"level":"debug","time":1710750660,"message":"This message appears only when log level set to Debug"}
{"level":"info","time":1710750660,"message":"This message appears when log level set to Debug or Info"}
{"level":"debug","foo":"bar","time":1710750660,"message":"some debug message"}
{"foo2":"bar2","time":1710750660}
*/
