package log_test

import (
	"os"
	"testing"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
	loglog "github.com/idpay/gokit-logger"
)

func TestNewLogger(t *testing.T) {
	loggers := getLoggers()
	l := loglog.NewLogger(loggers)

	_ = l.Log("msg", "message")
}

func TestNewLoggerWithLevel_Info(t *testing.T) {
	loggers := getLoggers()
	l := loglog.NewLogger(loggers)

	_ = kitlevel.Info(l).Log("msg", "message")
}

func TestNewLoggerWithLevel_Error(t *testing.T) {
	loggers := getLoggers()
	l := loglog.NewLogger(loggers)

	_ = kitlevel.Error(l).Log("err", "error")
}

func getLoggers() []kitlog.Logger {
	fmtLogger := kitlog.NewLogfmtLogger(os.Stdout)
	fmtLogger = kitlog.With(fmtLogger, "ts", kitlog.DefaultTimestampUTC)

	jsonLogger := kitlog.NewJSONLogger(os.Stdout)
	jsonLogger = kitlog.With(jsonLogger, "ts", kitlog.DefaultTimestampUTC)

	return []kitlog.Logger{
		fmtLogger,
		jsonLogger,
	}
}
