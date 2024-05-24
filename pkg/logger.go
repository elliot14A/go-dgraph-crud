package pkg

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func GetLogger() *slog.Logger {
	return logger
}

func SetupLogger() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))
}

