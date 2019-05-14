package lenpenbot

import (
	"github.com/kulaginds/lenpenbot/pkg"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (i *LenPenBot) Start(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	helloMessage := `*Бот для измерений*

Доступные команды:
/reg - зарегистрироваться в измерениях
/enlarge - измерить
/top - топ за все время
/today - топ за сегодня
/credit - сантиметры в кредит`
	return pkg.Reply(msg, helloMessage), nil
}
