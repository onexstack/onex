package bytelog

import (
	"code.byted.org/gopkg/logs/v2"
)

// cronLogger implement the cron.Logger interface.
type cronLogger struct{}

// NewLogger returns a cron logger.
func NewLogger() *cronLogger {
	return &cronLogger{}
}

// Debug logs routine messages about cron's operation.
func (l *cronLogger) Debug(msg string, keysAndValues ...any) {
	logs.V2.Debug().Str(msg).KVs(keysAndValues...).Emit()
}

// Info logs routine messages about cron's operation.
func (l *cronLogger) Info(msg string, keysAndValues ...any) {
	logs.V2.Info().Str(msg).KVs(keysAndValues...).Emit()
}

// Error logs an error condition.
func (l *cronLogger) Error(err error, msg string, keysAndValues ...any) {
	logs.V2.Error().Str(msg).Error(err).KVs(keysAndValues...).Emit()
}
