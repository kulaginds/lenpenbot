package lenpenbot

import (
	"errors"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kulaginds/lenpenbot/pkg"
	"github.com/kulaginds/lenpenbot/pkg/penis"
	"github.com/kulaginds/lenpenbot/pkg/randLength"
	"time"
)

const (
	enlargeMin = 1
	enlargeMax = 27
)

func (i *LenPenBot) Enlarge(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	isUserRegistered, err := i.store.IsUserRegistered(msg.From.ID, msg.Chat.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Reg: cannot check is user registered: userID=%d; chatID=%d", msg.From.ID, msg.Chat.ID))
	}

	if !isUserRegistered {
		return pkg.Reply(msg, "Чтобы измерить, нужно зарегистрироваться"), nil
	}

	today := time.Now().UTC()

	isEnlarge, err := i.store.IsEnlarge(msg.From.ID, msg.Chat.ID, today.Truncate(24 * time.Hour))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Reg: cannot check is enlarge: userID=%d; chatID=%d; today=%s", msg.From.ID, msg.Chat.ID, today))
	}

	if isEnlarge {
		return pkg.Reply(msg, "Вы уже измеряли сегодня"), nil
	}

	length := randLength.Generate(enlargeMin, enlargeMax)

	err = i.store.Enlarge(msg.From.ID, msg.Chat.ID, length, today)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Reg: cannot enlarge: userID=%d; chatID=%d; length=%d; today=%s", msg.From.ID, msg.Chat.ID, length, today))
	}

	resp := `Так-так, посмотрим...
Померили снова и оказалось: %d см.
%s`

	return pkg.Reply(msg, fmt.Sprintf(resp, length, penis.Generate(length))), nil
}
