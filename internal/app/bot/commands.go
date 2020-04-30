package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	// Get key from the dictionary.
	Get = "get"

	// Set key to the dictionary.
	Set = "set"

	// Delete key from the dictionary.
	Delete = "delete"
)

func (b *bot) command(message *tgbotapi.Message) {
	switch message.Command() {
	case Get, Set, Delete:
		b.dictionary(message)
	}
}