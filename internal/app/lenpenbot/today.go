package lenpenbot

import (
	"errors"
	"fmt"
	"time"

	"github.com/kulaginds/lenpenbot/pkg"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (i *LenPenBot) Today(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	updatedMin := time.Now().UTC().Truncate(24 * time.Hour)
	updatedMax := updatedMin.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	hasToday, err := i.store.HasToday(msg.Chat.ID, updatedMin, updatedMax)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Today: cannot check has today: chatID=%d; updatedMin=%s; updatedMax=%s: %s", msg.Chat.ID, updatedMin, updatedMax, err))
	}

	if !hasToday {
		return pkg.Reply(msg, "Не было измерений"), nil
	}

	today, err := i.store.GetToday(msg.Chat.ID, updatedMin)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Today: cannot get today top: chatID=%d; today=%s: %s", msg.Chat.ID, updatedMin, err))
	}

	return pkg.Reply(msg, today.Message), nil
}
