package main

import (
	"fmt"
	"github.com/kulaginds/lenpenbot/internal/app/lenpenbot"
	"github.com/kulaginds/lenpenbot/pkg"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/kulaginds/lenpenbot/internal/config"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	cfg := config.MustInitConfig()

	httpClient := &http.Client{}

	if cfg.GetProxyHost() != "" {
		proxyUrl, err := url.Parse(fmt.Sprintf("%s://%s", cfg.GetProxyProtocol(), cfg.GetProxyHost()))
		if err != nil {
			log.Fatal(err)
		}

		proxyUrl.User = url.UserPassword("order@ruskyhost.ru", "login911##")

		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
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

	botClient := lenpenbot.NewLenPenBot(bot)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		var msgConf *tgbotapi.MessageConfig

		switch {
		case strings.HasPrefix(update.Message.Text, "/start"):
			msgConf, err = botClient.Start(update.Message)
			break;
		default:
			msgConf = pkg.Reply(update.Message, "Неизвестная команда")
		}

		if err != nil {
			log.Println(err)
			continue
		}

		if msgConf != nil {
			_, err := bot.Send(msgConf)
			if err != nil {
				log.Println(err)
			}
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
