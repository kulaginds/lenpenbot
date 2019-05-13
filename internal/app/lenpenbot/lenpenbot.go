package lenpenbot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kulaginds/lenpenbot/pkg/store"
)

type BotAPIImplementation interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type LenPenBot struct {
	bot   BotAPIImplementation
	store store.Store
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

func NewLenPenBot(bot BotAPIImplementation, store store.Store) BotImplementation {
	return &LenPenBot{bot: bot, store: store}
}
