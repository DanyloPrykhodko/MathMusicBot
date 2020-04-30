package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func (b *bot) send(text string, chatID int64) (*tgbotapi.Message, error) {
	config := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatID,
		},
		Text:      text,
		ParseMode: tgbotapi.ModeMarkdown,
	}

	m, err := b.Send(config)
	if  err != nil {
		return nil, err
	}

	return &m, err
}

func (b *bot) sendSelfDestruct(text string, duration time.Duration, chatID int64)  {
	m, err := b.send(text, chatID)
	if err == nil {
		go func() {
			time.Sleep(duration)
			_, _ = b.delete(m)
		}()
	}
}

func (b *bot) reply(text string, message *tgbotapi.Message) (*tgbotapi.Message, error) {
	config := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           message.Chat.ID,
			ReplyToMessageID: message.MessageID,
		},
		Text:      text,
		ParseMode: tgbotapi.ModeMarkdown,
	}

	m, err := b.Send(config)
	if err != nil {
		return nil, err
	}

	return &m, err
}

func (b *bot) replySelfDestruct(text string, duration time.Duration, message *tgbotapi.Message)  {
	m, err := b.reply(text, message)
	if err == nil {
		go func() {
			time.Sleep(duration)
			_, _ = b.delete(m)
		}()
	}
}

func (b *bot) delete(message *tgbotapi.Message) (*tgbotapi.Message, error) {
	config := tgbotapi.DeleteMessageConfig{
		ChatID:    message.Chat.ID,
		MessageID: message.MessageID,
	}

	m, err := b.Send(config)
	if err != nil {
		return nil, err
	}

	return &m, err
}

