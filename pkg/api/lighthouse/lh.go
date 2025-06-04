package lighthouse

import (
	"github.com/s4bb4t/lighthouse/internal/usecase"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"github.com/s4bb4t/lighthouse/pkg/logger"
	"github.com/s4bb4t/lighthouse/pkg/telegram"
)

type Lighthouse struct {
	log    usecase.Logger
	notify usecase.Notify
}

func New(stage, apikey string) *Lighthouse {
	b, _ := telegram.New(apikey)
	return &Lighthouse{
		log:    logger.New(stage, sp.En, nil),
		notify: b,
	}
}

// ManualNew manually creates Lighthouse from provided telegram.Bot, kiba
func ManualNew(log usecase.Logger, notify usecase.Notify) *Lighthouse {
	return &Lighthouse{
		log:    log,
		notify: notify,
	}
}

func (l *Lighthouse) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}
func (l *Lighthouse) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}
func (l *Lighthouse) Error(e error, lvl levels.Level) {
	l.log.Error(e, lvl)
}

func (l *Lighthouse) AlertWarn(msg string) error {
	return l.notify.Warn(msg)
}
func (l *Lighthouse) AlertInfo(msg string) error {
	return l.notify.Info(msg)
}
func (l *Lighthouse) AlertError(e error) error {
	return l.notify.Error(e)
}
func (l *Lighthouse) AlertDebug(msg string) error {
	return l.notify.Debug(msg)
}
