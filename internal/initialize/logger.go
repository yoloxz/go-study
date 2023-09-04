package initialize

import (
	"os"

	"github.com/rs/zerolog"
)

func LoggerInit() zerolog.Logger {
	// 创建一个 Zerolog logger
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	return logger
}
