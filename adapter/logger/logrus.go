package logger

import (
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(log *logrus.Logger) Logger {
	return &logrusLogger{log: log}
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusEntry{
		entry: l.log.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogger) WithError(err error) Logger {
	return &logrusEntry{
		entry: l.log.WithError(err),
	}
}

type logrusEntry struct {
	entry *logrus.Entry
}

func (l *logrusEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusEntry) WithFields(fields Fields) Logger {
	return &logrusEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusEntry) WithError(err error) Logger {
	return &logrusEntry{
		entry: l.entry.WithError(err),
	}
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, field := range fields {
		logrusFields[index] = field
	}

	return logrusFields
}
