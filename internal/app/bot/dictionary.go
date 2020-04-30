package bot

import (
	"fmt"
	"github.com/drprykhodko/MathMusicBot/internal/app/model"
	"github.com/drprykhodko/MathMusicBot/internal/app/store/sqlstore"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"time"
)

func (b *bot) dictionary(message *tgbotapi.Message) {
	db := b.store.DB()

	args := message.CommandArguments()
	switch message.Command() {
	case Get:
		if args == "" {
			keys, err := b.store.Dictionary().GetKeys(db)
			if err != nil {
				panic(err)
			}
			b.replySelfDestruct(fmt.Sprintf("Используйте /%s _слово_\nДоступные слова: %s", Get, markdown(strings.Join(keys, ", "))), 30*time.Second, message)

			return
		}

		d, err := b.store.Dictionary().Get(db, args)
		if err != nil {
			if err == sqlstore.ErrorUnknownKey {
				b.replySelfDestruct("_Ничего не найдено_", 5*time.Second, message)

				return
			} else {
				panic(err)
			}
		}

		_, _ = b.reply(fmt.Sprintf("*%s*\n%v", args, markdown(d.Value)), message)

		return
	case Set:
		if args == "" {
			b.replySelfDestruct(fmt.Sprintf("Используйте /%s связку из _ключа_ и _значения_, а я попытаюсь всё правильно сделать", Set), 10*time.Second, message)

			return
		}
		
		d := &model.Dictionary{}
		
		if err := d.Parse(args); err != nil {
			switch err {
			case model.ErrorCantParseDictionary:
				b.replySelfDestruct("_У меня не получилось понять где ключ, а где значение_", 5*time.Second, message)
			case model.ErrorCantParseKey:
				b.replySelfDestruct("_У меня не получилось понять где ключ_", 5*time.Second, message)
			case model.ErrorCantParseValue:
				b.replySelfDestruct("_У меня не получилось понять где значение_", 5*time.Second, message)
			default:
				panic(err)
			}

			return
		}
		
		if err := b.store.Dictionary().Set(db, d); err != nil {
			b.replySelfDestruct("_Не удалось добавить новое слово_", 5*time.Second, message)

			return
		}
		
		_, _ = b.reply("_Новое слово успешно добавлено_", message)

		return
	case Delete:
		if args == "" {
			b.replySelfDestruct(fmt.Sprintf("Используйте /%s _слово_", Delete), 10*time.Second, message)

			return
		}
		_, err := b.store.Dictionary().Get(db, args)
		if err != nil {
			if err == sqlstore.ErrorUnknownKey {
				b.replySelfDestruct("_Такое слово отсутствует_", 5*time.Second, message)

				return
			} else {
				panic(err)
			}
		}
		
		err = b.store.Dictionary().Delete(db, args)
		if err != nil {
			if err == sqlstore.ErrorNothingToDelete {
				b.replySelfDestruct("_Такое слово отсутствует_", 5*time.Second, message)
			} else {
				b.replySelfDestruct("_Не удалось удалить слово_", 5*time.Second, message)
			}

			return
		}
		
		_, _ = b.reply("_Слово успешно удалено_", message)

		return
	}
}
