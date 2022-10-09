package log

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"os"
	"time"
)

type Config struct {
	Level  string
	Format string
}

func New(config *Config) log.Logger {
	var (
		l  log.Logger
		le level.Option
	)

	if config.Format == "json" {
		l = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	} else {
		l = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	}

	if len(config.Level) != 0 {
		l = log.With(l, "ts", log.TimestampFormat(
			func() time.Time { return time.Now().Local() },
			"2006-01-02 15:04:05.000 ",
		), "caller", log.Caller(5))
		switch config.Level {
		case "debug":
			le = level.AllowDebug()
		case "info":
			le = level.AllowInfo()
		case "warn":
			le = level.AllowWarn()
		case "error":
			le = level.AllowError()
		default:
			le = level.AllowInfo()
		}
	} else {
		l = log.With(l, "ts", log.TimestampFormat(
			func() time.Time { return time.Now().Local() },
			"2006-01-02 15:04:05.000 ",
		), "caller", log.DefaultCaller)
	}

	l = level.NewFilter(l, le)
	return l
}
