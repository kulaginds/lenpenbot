package lenpenbot

import (
	"github.com/kulaginds/lenpenbot/pkg/store"
	"github.com/kulaginds/lenpenbot/pkg/top"
	"github.com/kulaginds/lenpenbot/pkg/types"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type LenPenBot struct {
	bot   types.BotAPIImplementation
	store store.Store
	top   top.TopImplementation
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

func NewLenPenBot(bot types.BotAPIImplementation, store store.Store, top top.TopImplementation) BotImplementation {
	return &LenPenBot{bot: bot, store: store, top: top}
}
