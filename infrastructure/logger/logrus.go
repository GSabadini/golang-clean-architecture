package logger

import (
	"github.com/GSabadini/go-challenge/adapter/logger"
	logrusLib "github.com/sirupsen/logrus"
)

// NewLogrus returns the instance of logrus logger
func NewLogrus() *Logrus {
	log := logrusLib.New()
	log.SetFormatter(&logrusLib.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &Logrus{
		log: log,
	}
}

type Logrus struct {
	log *logrusLib.Logger
}

func NewLogrusLogger(log *logrusLib.Logger) logger.Logger {
	return &Logrus{log: log}
}

func (l *Logrus) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *Logrus) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *Logrus) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *Logrus) WithFields(fields logger.Fields) logger.Logger {
	return &LogrusEntry{
		entry: l.log.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *Logrus) WithError(err error) logger.Logger {
	return &LogrusEntry{
		entry: l.log.WithError(err),
	}
}

type LogrusEntry struct {
	entry *logrusLib.Entry
}

func (l *LogrusEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *LogrusEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *LogrusEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *LogrusEntry) WithFields(fields logger.Fields) logger.Logger {
	return &LogrusEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *LogrusEntry) WithError(err error) logger.Logger {
	return &LogrusEntry{
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
