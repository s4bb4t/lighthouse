package lighthouse

import (
	"github.com/s4bb4t/lighthouse/kibana"
	"github.com/s4bb4t/lighthouse/logger"
	"github.com/s4bb4t/lighthouse/tg"
)

type Lighthouse struct {
	*logger.Logger
	*tg.Bot
	*kibana.Hook
}

func New() {

}

func ManualNew(bot *tg.Bot, kib *kibana.Hook, log *logger.Logger) *Lighthouse {
	return &Lighthouse{
		Logger: log,
		Bot:    bot,
		Hook:   kib,
	}
}
