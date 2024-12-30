package empty

import "github.com/onexstack/onex/pkg/logger"

// emptyLogger is an implementation of the logger.Logger interface that performs no operations.
// This can be useful in contexts where a logger is required but logging output is not desired.
type emptyLogger struct{}

// Ensure that emptyLogger implements the logger.Logger interface.
var _ logger.Logger = (*emptyLogger)(nil)

// NewLogger returns a new instance of an empty logger.
func NewLogger() *emptyLogger {
	return &emptyLogger{}
}

// Debug logs a message at the Debug level. This implementation does nothing.
func (l *emptyLogger) Debug(msg string, keysAndValues ...any) {}

// Warn logs a message at the Warn level. This implementation does nothing.
func (l *emptyLogger) Warn(msg string, keysAndValues ...any) {}

// Info logs a message at the Info level. This implementation does nothing.
func (l *emptyLogger) Info(msg string, keysAndValues ...any) {}

// Error logs a message at the Error level. This implementation does nothing.
func (l *emptyLogger) Error(msg string, keysAndValues ...any) {}
