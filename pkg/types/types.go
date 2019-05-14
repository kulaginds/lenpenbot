package types

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type BotAPIImplementation interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetChatMember(config tgbotapi.ChatConfigWithUser) (tgbotapi.ChatMember, error)
}
