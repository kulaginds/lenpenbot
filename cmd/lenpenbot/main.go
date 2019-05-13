package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/kulaginds/lenpenbot/internal/config"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"
)

func main() {
	cfg := config.MustInitConfig()

	httpClient := &http.Client{}

	if cfg.GetProxyURL() != "" {
		proxyUrl, err := url.Parse(cfg.GetProxyURL())
		if err != nil {
			log.Fatal(err)
		}

		password, _ := proxyUrl.User.Password()
		proxyDialer, err := proxy.SOCKS5("tcp", proxyUrl.Host, &proxy.Auth{
			User: proxyUrl.User.Username(),
			Password: password,
		}, proxy.Direct)
		if err != nil {
			log.Fatal(err)
		}

		httpClient = &http.Client{
			Transport: &http.Transport{
				DialTLS: proxyDialer.Dial,
			},
		}
	}

	bot, err := tgbotapi.NewBotAPIWithClient(cfg.GetToken(), httpClient)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = cfg.GetDebug()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	var updates tgbotapi.UpdatesChannel

	if cfg.GetMode() == config.BotModeWebhook {
		updates, err = webhookMode(cfg, bot)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		updates = longPollMode(cfg, bot)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func webhookMode(cfg *config.Config, bot *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	_, err := bot.SetWebhook(
		tgbotapi.NewWebhookWithCert(
			fmt.Sprintf("%s%s", cfg.GetBotURL(), bot.Token),
			cfg.GetSSLCert(),
		),
	)
	if err != nil {
		return nil, err
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		return nil, err
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook(fmt.Sprintf("/%s", bot.Token))
	go http.ListenAndServeTLS(cfg.GetAddress(), cfg.GetSSLCert(), cfg.GetSSLKey(), nil)

	return updates, nil
}

func longPollMode(cfg *config.Config, bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = cfg.GetUpdateTimeout()

	updates, _ := bot.GetUpdatesChan(u)

	return updates
}
