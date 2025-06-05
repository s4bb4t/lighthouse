package telegram

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/internal/storage"
	"github.com/s4bb4t/lighthouse/internal/usecase"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"hash"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	wh = "https://webhook.site/cb526586-6f6c-4398-87b1-cc313a08d012"
	Dv = "Developer"
	Do = "DevOps"
	Bz = "Business"
)

var (
	sendErr         hash.Hash
	numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(Dv),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(Do),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(Bz),
		),
	)
)

func init() {
	var err error
	sendErr, err = sp.Registry.Reg(sp.New(sp.Err{
		Messages: map[string]string{
			sp.En: "failed to send message",
		},
		Desc:  "failed to send message",
		Hint:  "Check underlying error",
		Level: levels.LevelError,
	}).MustDone())
	if err != nil {
		panic(err)
	}
}

type Bot struct {
	wh      func(b *Bot, addr, port string) (error, chan error)
	storage usecase.Storage
	Api     *tgbotapi.BotAPI
	sync.RWMutex
}

func New(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	repo, err := storage.New()
	if err != nil {
		return nil, err
	}

	return &Bot{
		storage: repo,
		Api:     api,
	}, nil
}

func (b *Bot) SetWebHookHandler(h func(b *Bot, addr, port string) (error, chan error)) {
	b.wh = h
}

func (b *Bot) WebHook(url, port string) (error, chan error) {
	if b.wh != nil {
		return b.wh(b, url, port)
	}
	return b.StartDefaultWebHook(url, port)
}

func (b *Bot) handle(upd *tgbotapi.Update) {
	switch {
	case upd.Message == nil:
	case upd.Message.Command() == "start", upd.Message.Command() == "groups":
		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Choose your group")
		msg.ReplyMarkup = numericKeyboard
		b.Api.Send(msg)
	case upd.Message.Text == Dv, upd.Message.Text == Do, upd.Message.Text == Bz:
		b.Lock()
		err := b.storage.Put(upd.Message.Text, upd.Message.Chat.ID)
		if err != nil {
			sp.Wrap(sp.Ensure(err), sp.New(sp.Err{
				Messages: map[string]string{
					sp.En: "failed to subscribe user to alarm",
				},
				Desc:  "Failed to save user's to to storage",
				Hint:  "Check underlying Error",
				Level: levels.LevelError,
			}))
		}
		b.Unlock()
		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "✅ Subscribed to error notifications as *"+upd.Message.Text+"*")
		msg.ParseMode = "MarkdownV2"
		b.Api.Send(msg)
	default:
		b.Api.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Use /groups to subscribe to error notifications"))
	}
}

func (b *Bot) Info(msg string) error {
	b.RLock()
	defer b.RUnlock()

	subs, err := b.storage.Read(Dv)
	if err != nil {
		return sp.Wrap(sp.Ensure(err), sp.New(sp.Err{
			Messages: map[string]string{
				sp.En: "Failed to read users",
			},
			Desc:  "Failed to read subscribed users's ids",
			Hint:  "Check storage",
			Level: levels.LevelError,
		}))
	}

	for _, id := range subs {
		_, err = b.Api.Send(tgbotapi.NewMessage(id, "Info: "+msg))
		if err != nil {
			return sp.Wrap(sp.Ensure(err), sp.Registry.Get(sendErr))
		}
	}
	return nil
}

func (b *Bot) Error(e error) error {
	b.RLock()
	defer b.RUnlock()

	subs, err := b.storage.Read(Dv)
	if err != nil {
		return sp.Wrap(sp.Ensure(err), sp.New(sp.Err{
			Messages: map[string]string{
				sp.En: "Failed to read users",
			},
			Desc:  "Failed to read subscribed users's ids",
			Hint:  "Check storage",
			Level: levels.LevelError,
		}))
	}

	for _, id := range subs {
		msg := tgbotapi.NewMessage(id, prettify(e))
		msg.ParseMode = "MarkdownV2"
		_, err = b.Api.Send(msg)
		if err != nil {
			return sp.Wrap(sp.Ensure(err), sp.Registry.Get(sendErr))
		}
	}
	return nil
}

func (b *Bot) Warn(msg string) error {
	// TODO: implement
	_ = msg
	b.RLock()
	defer b.RUnlock()
	return fmt.Errorf("unmplemented")
}

func (b *Bot) Debug(msg string) error {
	// TODO: implement
	_ = msg
	b.RLock()
	defer b.RUnlock()
	return fmt.Errorf("unmplemented")
}

/*
func (b *Bot) listen() chan error {
	errCh := make(chan error)

	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := b.api.GetUpdatesChan(u)

		for upd := range updates {
			switch {
			case upd.Message == nil:
				continue
			case upd.Message.Command() == "start", upd.Message.Command() == "groups":
				msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Choose group")
				msg.ReplyMarkup = numericKeyboard
				_, err := b.api.Send(msg)
				if err != nil {
					errCh <- sp.Wrap(sp.Ensure(err), sp.Registry.Get(sendErr))
				}
				continue
			case upd.Message.Text == Dv, upd.Message.Text == Do, upd.Message.Text == Bz:
				b.Lock()
				err := b.storage.Put(upd.Message.Text, upd.Message.Chat.ID)
				if err != nil {
					errCh <- sp.Wrap(sp.Ensure(err), sp.New(sp.Err{
						Messages: map[string]string{
							sp.En: "failed to subscribe user to alarm",
						},
						Desc:  "Failed to save user's to to storage",
						Hint:  "Check underlying Error",
						Level: levels.LevelError,
					}))
				}
				b.Unlock()

				_, err = b.api.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "✅ Subscribed to error notifications"))
				if err != nil {
					errCh <- sp.Wrap(sp.Ensure(err), sp.Registry.Get(sendErr))
				}
			default:
				_, err := b.api.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Use /groups to subscribe to error notifications"))
				if err != nil {
					errCh <- sp.Wrap(sp.Ensure(err), sp.Registry.Get(sendErr))
				}
			}
		}
	}()

	return errCh
}
*/
