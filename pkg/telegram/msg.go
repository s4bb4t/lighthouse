package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
)

func (b *Bot) Info(msg string) error {
	b.RLock()
	defer b.RUnlock()

	subs, err := b.storage.Read("")
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

func (b *Bot) Error(e error, group string) error {
	b.RLock()
	defer b.RUnlock()

	subs, err := b.storage.Read(group)
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
