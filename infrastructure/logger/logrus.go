package logger

import (
	"github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/sirupsen/logrus"
)

// NewLogrus returns the instance of logrusWrapper logger
func NewLogrus() *logrusWrapper {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &logrusWrapper{
		log: log,
	}
}

type logrusWrapper struct {
	log *logrus.Logger
}

func NewLogrusLogger(log *logrus.Logger) logger.Logger {
	return &logrusWrapper{log: log}
}

func (l *logrusWrapper) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *logrusWrapper) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *logrusWrapper) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *logrusWrapper) WithFields(fields logger.Fields) logger.Logger {
	return &logrusEntry{
		entry: l.log.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusWrapper) WithError(err error) logger.Logger {
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

func (l *logrusEntry) WithFields(fields logger.Fields) logger.Logger {
	return &logrusEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusEntry) WithError(err error) logger.Logger {
	return &logrusEntry{
		entry: l.entry.WithError(err),
	}
}

func convertToLogrusFields(fields logger.Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, field := range fields {
		logrusFields[index] = field
	}

	return logrusFields
}
