package logwrapper

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      string
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{"1", "Invalid arg: %s for userid-%s"}
	invalidArgValueMessage = Event{"2", "Invalid value for argument: %s: %s for userid-%s"}
	missingArgMessage      = Event{"3", "Missing arg: %s for userid- %s"}
)

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string, userId string) {
	l.Errorf(invalidArgMessage.message, argumentName, userId)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string, userId string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue, userId)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string, userId string) {
	l.Errorf(missingArgMessage.message, argumentName, userId)
}
