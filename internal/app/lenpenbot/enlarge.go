package lenpenbot

import (
	"errors"
	"fmt"
	"time"

	"github.com/kulaginds/lenpenbot/pkg"
	"github.com/kulaginds/lenpenbot/pkg/penis"
	"github.com/kulaginds/lenpenbot/pkg/randLength"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	enlargeMin = 1
	enlargeMax = 27
)

func (i *LenPenBot) Enlarge(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	isUserRegistered, err := i.store.IsUserRegistered(msg.From.ID, msg.Chat.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Enlarge: cannot check is user registered: userID=%d; chatID=%d: %s", msg.From.ID, msg.Chat.ID, err))
	}

	if !isUserRegistered {
		return pkg.Reply(msg, "Чтобы измерить, нужно зарегистрироваться"), nil
	}

	today := time.Now().UTC()
	todayDate := today.Truncate(24 * time.Hour)

	isEnlarge, err := i.store.IsEnlarge(msg.From.ID, msg.Chat.ID, todayDate)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Enlarge: cannot check is enlarge: userID=%d; chatID=%d; today=%s: %s", msg.From.ID, msg.Chat.ID, today, err))
	}

	if isEnlarge {
		return pkg.Reply(msg, "Вы уже измеряли сегодня"), nil
	}

	length := randLength.Generate(enlargeMin, enlargeMax)

	err = i.store.Enlarge(msg.From.ID, msg.Chat.ID, length, today)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Enlarge: cannot enlarge: userID=%d; chatID=%d; length=%d; today=%s: %s", msg.From.ID, msg.Chat.ID, length, today, err))
	}

	todayTop, err := i.top.GenerateToday(msg.Chat.ID, todayDate)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Enlarge: cannot generate today: chatID=%d; today=%s: %s", msg.Chat.ID, today, err))
	}

	err = i.store.SetToday(msg.Chat.ID, todayDate, todayTop)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Enlarge: cannot set today: chatID=%d; today=%s: %s", msg.Chat.ID, today, err))
	}

	topTop, err := i.top.GenerateTop(msg.Chat.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Enlarge: cannot generate top: chatID=%d: %s", msg.Chat.ID, err))
	}

	err = i.store.SetTop(msg.Chat.ID, topTop)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Enlarge: cannot set top: chatID=%d: %s", msg.Chat.ID, today, err))
	}

	resp := `Так-так, посмотрим...
Померили снова и оказалось: %d см.
%s`

	return pkg.Reply(msg, fmt.Sprintf(resp, length, penis.Generate(length))), nil
}
