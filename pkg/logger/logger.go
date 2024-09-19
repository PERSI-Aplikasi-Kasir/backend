package logger

import (
	"backend/pkg/env"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var lumberjackLogger *lumberjack.Logger

func InitializeLogger(filename string) {
	fmt.Println("===== Initialize Logger =====")
	zerolog.TimeFieldFormat = "15:04:05, 2006-01-02"

	var logger zerolog.Logger

	lumberjackLogger = &lumberjack.Logger{
		MaxSize:  100, // megabytes
		MaxAge:   14,  // days
		Filename: filename,
	}

	writers := []io.Writer{zerolog.ConsoleWriter{Out: os.Stderr}, lumberjackLogger}
	mw := io.MultiWriter(writers...)

	if env.ENVIRONMENT == "production" {
		logger = zerolog.New(lumberjackLogger).With().Timestamp().Caller().Logger().Level(zerolog.InfoLevel)
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			trimmedPath := file
			if idx := strings.Index(file, "backend/"); idx != -1 {
				trimmedPath = file[idx:]
			}
			return fmt.Sprintf("%s:%d", trimmedPath, line)
		}
	} else {
		logger = zerolog.New(mw).With().Timestamp().Caller().Logger().Level(zerolog.DebugLevel)
	}

	log.Logger = logger
	fmt.Println("✓ Logger initialized")
}

func UnsyncLogger() {
	if lumberjackLogger != nil {
		if err := lumberjackLogger.Close(); err != nil {
			log.Error().Err(err).Msg("Error while closing the logger")
			return
		}

		lumberjackLogger = nil
	}

	fmt.Println("✓ Logger closed")
}

func RotateLogger() {
	if lumberjackLogger != nil {
		if err := lumberjackLogger.Rotate(); err != nil {
			log.Error().Err(err).Msg("Error while rotating logs")
			return
		}
	}

	fmt.Println("✓ Logger rotated")
}
