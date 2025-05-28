package lighthouse

import (
	"github.com/s4bb4t/lighthouse/kibana"
	"github.com/s4bb4t/lighthouse/tg"
)

type Lighthouse struct {
	b *tg.Bot
	k *kibana.Hook
}

func NewLighthouse(bot *tg.Bot, kib *kibana.Hook) *Lighthouse {
	return &Lighthouse{
		b: nil,
		k: nil,
	}
}
