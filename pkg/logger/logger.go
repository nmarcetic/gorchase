package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
)

// Logger wraps logger lib
type Logger struct {
	*zap.Logger
}

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// New returns new instance of Logger
func New() (*Logger, error) {
	prodLogger, err := zap.NewProduction()
	if err != nil {
		return &Logger{}, fmt.Errorf("Cannot initialize logger. Error: %s", err)
	}

	defer prodLogger.Sync()

	var logger = &Logger{prodLogger}

	return logger, nil
}

// InvalidArg is a standard error message
func (l *Logger) InvalidArg(argumentName string) {
	l.Sugar().Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *Logger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Sugar().Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *Logger) MissingArg(argumentName string) {
	l.Sugar().Errorf(missingArgMessage.message, argumentName)
}
