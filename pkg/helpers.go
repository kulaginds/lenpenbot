package pkg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func Reply(msg *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	msgConf := tgbotapi.NewMessage(msg.Chat.ID, text)
	msgConf.ReplyToMessageID = msg.MessageID
	msgConf.ParseMode = "markdown"

	return &msgConf
}
