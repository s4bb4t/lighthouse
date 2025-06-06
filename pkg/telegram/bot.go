package telegram

import (
	"github.com/s4bb4t/lighthouse/internal/storage"
	"github.com/s4bb4t/lighthouse/pkg/core"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	DefaultDevGroup    = "Developer"
	DefaultDevOpsGroup = "DevOps"
	DefaultBusGroup    = "Business"
)

type Bot struct {
	kb      *tgbotapi.ReplyKeyboardMarkup
	wh      func(b *Bot, addr, port string) (error, chan error)
	storage core.Storage
	Api     *tgbotapi.BotAPI
	sync.RWMutex
}

func New(token string, groups []string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	repo, err := storage.New()
	if err != nil {
		return nil, err
	}

	k := tgbotapi.NewReplyKeyboard()
	switch {
	case groups == nil || len(groups) == 0:
		k.Keyboard = append(k.Keyboard, []tgbotapi.KeyboardButton{{Text: DefaultDevGroup}, {Text: DefaultDevOpsGroup}, {Text: DefaultBusGroup}})
	default:
		for _, group := range groups {
			if group != "" {
				k.Keyboard = append(k.Keyboard, []tgbotapi.KeyboardButton{{Text: group}})
			}
		}
	}

	return &Bot{
		kb:      &k,
		storage: repo,
		Api:     api,
	}, nil
}
