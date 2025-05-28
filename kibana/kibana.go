package kibana

import (
	"github.com/s4bb4t/lighthouse/core/sp"
)

type Hook struct {
	client *Client
}

func NewHook(cfg Config) *Hook {
	return &Hook{
		client: NewClient(cfg),
	}
}

func (h *Hook) Fire(err *sp.Error) error {
	return h.client.LogError(err)
}

func LogToKibana(hook *Hook, err *sp.Error) error {
	return hook.Fire(err)
}
