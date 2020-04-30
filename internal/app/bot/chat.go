package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *bot) newChatMembers(message *tgbotapi.Message) {
	for _, member := range *message.NewChatMembers {
		_, _ = b.reply(fmt.Sprintf("Добро пожаловать, %s 👋", tag(member)), message)
	}
}

func tag(user tgbotapi.User) string {
	if user.UserName == "" {
		return fmt.Sprintf("[%s](tg://user?id=%d)", markdown(user.String()), user.ID)
	} else {
		return "@" + markdown(user.String())
	}
}
