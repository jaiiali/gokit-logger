# GoKit Multi Logger

## Install

```sh
go get "github.com/idpay/gokit-logger"
```

## Import

```sh
import "github.com/idpay/gokit-logger"
```

## Usage

```go
import (
	kitlog "github.com/go-kit/kit/log"
	loglog "github.com/idpay/gokit-logger"
)
```

```go
loggers := []kitlog.Logger{
	kitlog.NewLogfmtLogger(os.Stdout),
	kitlog.NewJSONLogger(os.Stdout),
}

l := loglog.NewLogger(loggers)

l.Log("msg", "message")

// Output:
// msg=message
// {"msg":"message"}
```
