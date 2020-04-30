package bot

import (
	"github.com/drprykhodko/MathMusicBot/internal/app/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

type bot struct {
	token string
	*tgbotapi.BotAPI
	store store.Storer
}

func newBot(token string, store store.Storer) *bot {
	s := &bot{
		token: token,
		store: store,
	}

	return s
}

func (b *bot) launch() {
	defer func() {
		if r := recover(); r != nil {
			time.Sleep(3 * time.Second)
			b.launch()
		}
	}()

	botApi, err := tgbotapi.NewBotAPI(b.token)
	if err != nil {
		panic("error creating newBot telegram bot client")
	}
	b.BotAPI = botApi

	updates, _ := b.GetUpdatesChan(tgbotapi.UpdateConfig{
		Timeout: 60,
	})

	for update := range updates {
		if update.Message != nil {
			b.message(update.Message)
		}
	}
}

func (b *bot) message(message *tgbotapi.Message) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			m, err := b.send("_Что-то пошло не так_", message.Chat.ID)
			if err == nil {
				time.Sleep(5*time.Second)
				_, _ = b.delete(m)
			}
		}
	}()

	switch {
	case message.IsCommand():
		b.command(message)
	case message.NewChatMembers != nil && len(*message.NewChatMembers) > 0:
		b.newChatMembers(message)
	}
}
