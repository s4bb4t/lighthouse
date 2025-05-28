package hooks

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/core/levels"
	"github.com/s4bb4t/lighthouse/core/sp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap(err *sp.SPError, lvl levels.ErrorLevel) []zapcore.Field {
	var f []zapcore.Field

	f = append(f, zap.String("path", path))
	f = append(f, zap.String("desc", err.ReadDesc()))
	f = append(f, zap.String("hint", err.ReadHint()))
	f = append(f, zap.Any("time", err.ReadTime()))

	return f
}
