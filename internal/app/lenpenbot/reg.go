package lenpenbot

import (
	"errors"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kulaginds/lenpenbot/pkg"
)

func (i *LenPenBot) Reg(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	if msg.Chat == nil {
		return pkg.Reply(msg, "Регистрироваться можно только в чате"), nil
	}

	isUserRegistered, err := i.store.IsUserRegistered(msg.From.ID, msg.Chat.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Reg: cannot check is user registered: userID=%d; chatID=%d", msg.From.ID, msg.Chat.ID))
	}

	if isUserRegistered {
		return pkg.Reply(msg, "Вы уже зарегистрированы"), nil
	}

	err = i.store.RegisterUser(msg.From.ID, msg.Chat.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Reg: cannot register user: userID=%d; chatID=%d", msg.From.ID, msg.Chat.ID))
	}

	return pkg.Reply(msg, "Вы успешно зарегистрированы"), nil
}
