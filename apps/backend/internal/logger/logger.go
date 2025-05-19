package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func New() {
	opts := &tint.Options{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, opts),
	))
}
