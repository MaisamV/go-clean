package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger wraps zerolog.Logger
type Logger struct {
	zerolog.Logger
}

// New creates a new logger instance
func New() *Logger {
	// Configure zerolog
	zerolog.TimeFieldFormat = time.RFC3339
	
	// Create logger with console writer for development
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	
	logger := zerolog.New(consoleWriter).With().Timestamp().Logger()
	
	return &Logger{Logger: logger}
}

// NewWithLevel creates a new logger with specified level
func NewWithLevel(level string) *Logger {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}
	
	zerolog.SetGlobalLevel(logLevel)
	return New()
}