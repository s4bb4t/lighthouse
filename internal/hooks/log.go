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
		String: err.ReadDesc(),
	})
	f = append(f, zapcore.Field{
		Key:    "hint",
		Type:   15,
		String: err.ReadHint(),
	})
	f = append(f, zapcore.Field{
		Key:    "path",
		Type:   15,
		String: err.ReadSource(),
	})
	f = append(f, zapcore.Field{
		Key:    "time",
		Type:   15,
		String: err.ReadTime().String(),
	})

	return f
}

func Slog(e *sp.Error, lvl levels.Level) []any {
	err := e.Spin(lvl)
	var f []any

	f = append(f, slog.Attr{
		Key:   "desc",
		Value: slog.StringValue(err.ReadDesc()),
	})
	f = append(f, slog.Attr{
		Key:   "hint",
		Value: slog.StringValue(err.ReadHint()),
	})
	f = append(f, slog.Attr{
		Key:   "source",
		Value: slog.StringValue(err.ReadSource()),
	})
	f = append(f, slog.Attr{
		Key:   "err_time",
		Value: slog.StringValue(err.ReadTime().Format("2006.01.02 15:04:05")),
	})

	return f
}
