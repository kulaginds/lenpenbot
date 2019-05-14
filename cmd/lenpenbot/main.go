package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/kulaginds/lenpenbot/internal/app/lenpenbot"
	"github.com/kulaginds/lenpenbot/internal/config"
	"github.com/kulaginds/lenpenbot/pkg/store/pgstore"
	top2 "github.com/kulaginds/lenpenbot/pkg/top"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.MustInitConfig()

	httpClient := &http.Client{}

	if cfg.GetProxyHost() != "" {
		proxyUrl, err := url.Parse(fmt.Sprintf("%s://%s", cfg.GetProxyProtocol(), cfg.GetProxyHost()))
		if err != nil {
			log.Fatal(err)
		}

		proxyUrl.User = url.UserPassword(cfg.GetProxyUser(), cfg.GetProxyPassword())

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

	log.Printf("Authorized on account %s (%d)", bot.Self.UserName, bot.Self.ID)

	var updates tgbotapi.UpdatesChannel

	if cfg.GetMode() == config.BotModeWebhook {
		updates, err = webhookMode(cfg, bot)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		updates = longPollMode(cfg, bot)
	}

	dbConn, err := sql.Open("postgres", cfg.GetDatabaseDSN())
	if err != nil {
		log.Fatal(err)
	}

	store := pgstore.NewPGStore(dbConn)
	top := top2.NewTop(bot, store)

	botClient := lenpenbot.NewLenPenBot(bot, store, top)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		var msgConf *tgbotapi.MessageConfig

		switch {
		case strings.HasPrefix(update.Message.Text, "/start"):
			msgConf, err = botClient.Start(update.Message)
			break;
		case strings.HasPrefix(update.Message.Text, "/reg"):
			msgConf, err = botClient.Reg(update.Message)
			break;
		case strings.HasPrefix(update.Message.Text, "/enlarge"):
			msgConf, err = botClient.Enlarge(update.Message)
			break;
		case strings.HasPrefix(update.Message.Text, "/shit"):
			msgConf, err = botClient.Shit(update.Message)
			break;
		case strings.HasPrefix(update.Message.Text, "/today"):
			msgConf, err = botClient.Today(update.Message)
			break;
		case strings.HasPrefix(update.Message.Text, "/top"):
			msgConf, err = botClient.Top(update.Message)
			break;
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
	whCfg := tgbotapi.NewWebhook(fmt.Sprintf("%s%s", cfg.GetBotURL(), bot.Token))

	if cfg.GetSSLCert() != "" {
		whCfg = tgbotapi.NewWebhookWithCert(
			fmt.Sprintf("%s%s", cfg.GetBotURL(), bot.Token),
			cfg.GetSSLCert(),
		)
	}

	_, err := bot.SetWebhook(whCfg)
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

	address := cfg.GetAddress()

	if cfg.GetUseEnvPort() {
		address = fmt.Sprintf(":%d", cfg.GetPort())
	}

	if cfg.GetSSLCert() != "" {
		go http.ListenAndServeTLS(address, cfg.GetSSLCert(), cfg.GetSSLKey(), nil)
	} else {
		go http.ListenAndServe(address, nil)
	}

	return updates, nil
}

func longPollMode(cfg *config.Config, bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = cfg.GetUpdateTimeout()

	updates, _ := bot.GetUpdatesChan(u)

	return updates
}
