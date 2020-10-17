package logger

import (
	"github.com/GSabadini/go-challenge/adapter/logger"
	logrusLib "github.com/sirupsen/logrus"
)

// NewLogrus returns the instance of logrus logger
func NewLogrus() *logrus {
	log := logrusLib.New()
	log.SetFormatter(&logrusLib.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &logrus{
		log: log,
	}
}

type logrus struct {
	log *logrusLib.Logger
}

func NewLogrusLogger(log *logrusLib.Logger) logger.Logger {
	return &logrus{log: log}
}

func (l *logrus) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *logrus) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *logrus) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *logrus) WithFields(fields logger.Fields) logger.Logger {
	return &logrusEntry{
		entry: l.log.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrus) WithError(err error) logger.Logger {
	return &logrusEntry{
		entry: l.log.WithError(err),
	}
}

type logrusEntry struct {
	entry *logrusLib.Entry
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

func convertToLogrusFields(fields logger.Fields) logrusLib.Fields {
	logrusFields := logrusLib.Fields{}
	for index, field := range fields {
		logrusFields[index] = field
	}

	return logrusFields
}
