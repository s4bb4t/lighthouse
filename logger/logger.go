package logger

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/core/levels"
	"github.com/s4bb4t/lighthouse/core/sp"
	"github.com/s4bb4t/lighthouse/hooks"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

const (
	Local = "local"
	Dev   = "dev"
	Prod  = "prod"
)

type Logger struct {
	pd    func(layers int) string
	log   *slog.Logger
	stage string
	lg    string
	noop  bool
}

// Noop - creates new Logger that does nothing
//
// # It's useful for testing purposes
//
// Example:
//
//	l := logger.Noop()
//	l.Error(errors.New("test error"), logger.LevelHighDebug)
func Noop() *Logger {
	l := &Logger{noop: true}
	return l
}

// New - creates new Logger
//
// stage - one of Local, Dev, Prod
// lg - language code
// out - io.Writer to write logs to
//
// Logger's language is used only to define sp.Error's message
func New(stage, lg string, out io.Writer) *Logger {
	if out == nil {
		out = os.Stdout
	}

	l := &Logger{lg: lg, stage: stage, pd: func(layers int) string {
		_, file, line, ok := runtime.Caller(layers + 1)
		if ok {
			absPath, err := filepath.Abs(file)
			if err != nil {
				panic(err)
			}
			return fmt.Sprintf("%s:%d", absPath, line)
		}
		return "unknown"
	}}
	switch stage {
	default:
		l.log = slog.New(newPrettyHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case Dev, Prod:
		l.log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	}
	return l
}

// With - adds fields to logger
//
// It's a shortcut for slog.With()
/*
func (l *Logger) With(args ...any) {
	if l.noop {
		return
	}
	l.log = l.log.With(args...)
}
*/

// Error - logs error
//
// If error is sp.Error - it prepares and appears at right way
// if e is not sp.Error, it will be logged at classic way
//
// if e is nil, it will be ignored
//
// lvl - error level
func (l *Logger) Error(e error, lvl levels.Level) {
	if l.noop || e == nil {
		return
	}
	// check if e is sp.Error
	err, ok := sp.Cast(e)
	if !ok {
		// this is not sp.Error
		// log it at classic way
		l.log.Error(e.Error(), slog.String("source", l.pd(1)))
		return
	}
	// spin-prepare and log error
	l.log.Error(err.ReadMsg(l.lg), hooks.Slog(err, lvl)...)
}

// Debug - prints additional debug log to Logger's out
func (l *Logger) Debug(msg string, args ...any) {
	if l.noop {
		return
	}
	args = append(args, slog.String("log_at", l.pd(1)))
	l.log.Debug(msg, args...)
}

// Info - prints additional info to Logger's out
func (l *Logger) Info(msg string, args ...any) {
	if l.noop {
		return
	}
	l.log.Info(msg, args...)
}
