package lenpenbot

import (
	"errors"
	"fmt"

	"github.com/kulaginds/lenpenbot/pkg"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (i *LenPenBot) Top(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	hasTop, err := i.store.HasTop(msg.Chat.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Top: cannot check has top: chatID=%d: %s", msg.Chat.ID, err))
	}

	if !hasTop {
		return pkg.Reply(msg, "Не было измерений"), nil
	}

	top, err := i.store.GetTop(msg.Chat.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Top: cannot get top: chatID=%d: %s", msg.Chat.ID, err))
	}

	return pkg.Reply(msg, top.Message), nil
}
