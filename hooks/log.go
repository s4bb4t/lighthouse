package hooks

import (
	"github.com/s4bb4t/lighthouse/core/levels"
	"github.com/s4bb4t/lighthouse/core/sp"
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
		String: err.ReadPath(),
	})
	f = append(f, zapcore.Field{
		Key:     "time",
		Type:    11,
		Integer: err.ReadTime().Unix(),
	})

	return f
}

func Slog(err *sp.Error) []slog.Attr {
	var f []slog.Attr

	f = append(f, slog.Attr{
		Key:   "desc",
		Value: slog.StringValue(err.ReadDesc()),
	})
	f = append(f, slog.Attr{
		Key:   "hint",
		Value: slog.StringValue(err.ReadHint()),
	})
	f = append(f, slog.Attr{
		Key:   "path",
		Value: slog.StringValue(err.ReadPath()),
	})
	f = append(f, slog.Attr{
		Key:   "time",
		Value: slog.AnyValue(err.ReadTime().Unix()),
	})

	return f
}
