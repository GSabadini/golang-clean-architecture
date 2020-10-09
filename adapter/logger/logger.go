package logger

// Logger
type Logger interface {
	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger

	WithError(err error) Logger
}

// Fields
type Fields map[string]interface{}
