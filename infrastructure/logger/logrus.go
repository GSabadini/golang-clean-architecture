package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogrus() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return log
}
