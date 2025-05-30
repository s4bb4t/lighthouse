package logger

import (
	"context"
	"fmt"
	"io"
	stdLog "log"

	"github.com/fatih/color"
	"log/slog"
)

type PrettyHandler struct {
	opts *slog.HandlerOptions
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

func newPrettyHandler(out io.Writer, opts *slog.HandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	var b string
	cnt := r.NumAttrs()
	if cnt != 0 {
		b = "\n"
	}
	r.Attrs(func(a slog.Attr) bool {
		cnt--
		if cnt == 0 {
			b += fmt.Sprintf("\t%s = %s", a.Key, a.Value)

		} else {
			b += fmt.Sprintf("\t%s = %s\n", a.Key, a.Value)
		}
		return true
	})

	timeStr := r.Time.Format("[Jan 02 - 15:04:05]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(b),
	)

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
