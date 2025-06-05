package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
)

func (b *Bot) handle(upd *tgbotapi.Update) {
	switch {
	case upd.Message == nil:
	case upd.Message.Command() == "start", upd.Message.Command() == "groups":
		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Choose your group")
		msg.ReplyMarkup = b.kb
		b.Api.Send(msg)
	case b.checkGroup(upd.Message.Text):
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

func (b *Bot) checkGroup(g string) bool {
	for _, v := range b.kb.Keyboard {
		for _, v2 := range v {
			if v2.Text == g {
				return true
			}
		}
	}
	return false
}
