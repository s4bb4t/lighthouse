package hooks

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"go.uber.org/zap/zapcore"
	"log/slog"
	"strconv"
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
	f = append(f, zapcore.Field{
		Key:    "time",
		Type:   15,
		String: err.Time().String(),
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
	f = append(f, slog.Attr{
		Key:   "err_time",
		Value: slog.StringValue(err.Time().Format("2006.01.02 15:04:05")),
	})
	f = append(f, slog.Attr{
		Key:   "source",
		Value: slog.StringValue(err.Source()),
	})
	for i, call := range err.Stack() {
		f = append(f, slog.Attr{
			Key:   "trace call " + strconv.Itoa(i+1),
			Value: slog.StringValue(call),
		})
	}

	return f
}
