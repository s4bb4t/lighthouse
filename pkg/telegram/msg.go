package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
)

func (b *Bot) Info(msg string) error {
	b.RLock()
	defer b.RUnlock()

	subs, err := b.storage.Read("")
	if err != nil {
		return sp.Wrap(sp.Ensure(err), sp.New(sp.Sample{
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

func (b *Bot) readIds(group string) ([]int64, error) {
	subs, err := b.storage.Read(group)
	if err != nil {
		return nil, sp.Wrap(sp.Ensure(err), sp.New(sp.Sample{
			Messages: map[string]string{
				sp.En: "Failed to read users",
			},
			Desc:  "Failed to read subscribed users's ids",
			Hint:  "Check storage",
			Level: levels.LevelError,
		}))
	}
	return subs, nil
}

func (b *Bot) sendTo(ids []int64, msg *tgbotapi.MessageConfig) error {
	for _, id := range ids {
		msg.ChatID = id
		_, err := b.Api.Send(*msg)
		if err != nil {
			return sp.Registry.Get(sendErr).SetCaused(err)
		}
	}
	return nil
}

func (b *Bot) Error(e error, groups ...string) error {
	b.RLock()
	defer b.RUnlock()

	msg := tgbotapi.NewMessage(0, prettify(e))
	msg.ParseMode = "MarkdownV2"

	var subs []int64
	var err error

	if groups == nil || len(groups) == 0 {
		subs, err = b.readIds("")
		if err != nil {
			return err
		}
	}

	for _, group := range groups {
		ids, err := b.readIds(group)
		subs = append(subs, ids...)
		if err != nil {
			return err
		}
	}

	return b.sendTo(subs, &msg)
}
