package log

import (
	kitlog "github.com/go-kit/kit/log"
)

type Logger struct {
	loggers []kitlog.Logger
}

func (l *Logger) Log(keyvals ...interface{}) error {
	for _, logger := range l.loggers {
		_ = logger.Log(keyvals...)
	}

	return nil
}

func NewLogger(loggers []kitlog.Logger) *Logger {
	return &Logger{
		loggers: loggers,
	}
}
