package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

// Init инициализирует логгер
func Init(level string) {
	zerolog.SetGlobalLevel(parseLevel(level))
	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func parseLevel(level string) zerolog.Level {
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}
