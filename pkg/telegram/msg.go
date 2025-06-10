package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sperror"
)

func (b *Bot) readIds(group string) ([]int64, error) {
	subs, err := b.storage.Read(group)
	if err != nil {
		return nil, sperror.Wrap(sperror.Ensure(err), sperror.New(sperror.Sample{
			Messages: map[string]string{
				sperror.En: "Failed to read users",
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
			return sperror.New(sperror.Sample{
				Messages: map[string]string{
					sperror.En: "Failed to send message",
				},
				Desc:     "Some error happened while sending message to user",
				Hint:     "Check underlying error",
				HttpCode: 500,
				Level:    levels.LevelError,
				Cause:    err,
			})
		}
	}
	return nil
}

// Info sends an informational message to all subscribed users and returns an error if the operation fails.
func (b *Bot) Info(msg string) error {
	b.RLock()
	defer b.RUnlock()

	subs, err := b.readIds("")
	if err != nil {
		return err
	}

	m := tgbotapi.NewMessage(0, "Info: "+msg)
	return b.sendTo(subs, &m)
}

// Error sends a formatted error message to the specified groups or all subscribed users and returns an error if it fails.
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
