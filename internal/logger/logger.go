package logger

import (
	"github.com/Imnarka/user-service/internal/config"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Entry
}

func InitLogger(cfg *config.Config) *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	entry := logger.WithFields(logrus.Fields{
		"service": "user-service",
	})
	return &Logger{Entry: entry}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{Entry: l.Entry.WithField(key, value)}
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{Entry: l.Entry.WithError(err)}
}
