package logger

// Logger defines the methods for logging at different levels.
type Logger interface {
	// Debug logs a message at the debug level with optional key-value pairs.
	Debug(message string, keysAndValues ...any)

	// Warn logs a message at the warning level with optional key-value pairs.
	Warn(message string, keysAndValues ...any)

	// Info logs a message at the info level with optional key-value pairs.
	Info(message string, keysAndValues ...any)

	// Error logs a message at the error level with optional key-value pairs.
	Error(message string, keysAndValues ...any)
}
