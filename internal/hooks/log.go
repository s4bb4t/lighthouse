package hooks

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"go.uber.org/zap/zapcore"
	"log/slog"
)

func Zap(e *sp.Error, lvl levels.Level) []zapcore.Field {
	err := e.Spin(lvl)
	var f []zapcore.Field

	f = append(f, zapcore.Field{
		Key:    "desc",
		Type:   15,
		String: err.Desc(),
	})
	f = append(f, zapcore.Field{
		Key:    "hint",
		Type:   15,
		String: err.Hint(),
	})
	f = append(f, zapcore.Field{
		Key:    "path",
		Type:   15,
		String: err.Source(),
	})

	return f
}

func Slog(e *sp.Error, lvl levels.Level) []any {
	err := e.Spin(lvl)
	var f []any

	f = append(f, slog.Attr{
		Key:   "desc",
		Value: slog.StringValue(err.Desc()),
	})
	f = append(f, slog.Attr{
		Key:   "hint",
		Value: slog.StringValue(err.Hint()),
	})
	if err.Caused() != nil {
		f = append(f, slog.Attr{
			Key:   "error",
			Value: slog.AnyValue(err.Caused()),
		})
	}
	f = append(f, slog.Attr{
		Key:   "source",
		Value: slog.StringValue(err.Source()),
	})

	return f
}
