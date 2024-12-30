package onex

import (
	"github.com/onexstack/onex/pkg/log"
	"github.com/onexstack/onex/pkg/logger"
)

// onexLogger provides an implementation of the logger.Logger interface.
type onexLogger struct{}

// Ensure that onexLogger implements the logger.Logger interface.
var _ logger.Logger = (*onexLogger)(nil)

// NewLogger creates a new instance of onexLogger.
func NewLogger() *onexLogger {
	return &onexLogger{}
}

// Debug logs a debug message with any additional key-value pairs.
func (l *onexLogger) Debug(msg string, kvs ...any) {
	log.Debugw(msg, kvs...)
}

// Warn logs a warning message with any additional key-value pairs.
func (l *onexLogger) Warn(msg string, kvs ...any) {
	log.Warnw(msg, kvs...)
}

// Info logs an informational message with any additional key-value pairs.
func (l *onexLogger) Info(msg string, kvs ...any) {
	log.Infow(msg, kvs...)
}

// Error logs an error message with any additional key-value pairs.
func (l *onexLogger) Error(msg string, kvs ...any) {
	log.Errorw(nil, msg, kvs...)
}
