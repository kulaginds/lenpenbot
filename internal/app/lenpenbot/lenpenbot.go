package lenpenbot

import "github.com/go-telegram-bot-api/telegram-bot-api"

type LenPenBot struct {
	bot *tgbotapi.BotAPI
}

type BotImplementation interface {
	Start(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	Reg(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	Enlarge(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	Shit(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	Top(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	Today(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	Credit(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
}

func NewLenPenBot(bot *tgbotapi.BotAPI) BotImplementation {
	return &LenPenBot{bot: bot}
}
