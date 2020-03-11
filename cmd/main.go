package main

import (
	"os"

	kitlog "github.com/go-kit/kit/log"
	loglog "github.com/idpay/gokit-logger"
)

func main() {
	loggers := []kitlog.Logger{
		kitlog.NewLogfmtLogger(os.Stdout),
		kitlog.NewJSONLogger(os.Stdout),
	}

	l := loglog.NewLogger(loggers)

	_ = l.Log("msg", "message")
}
