package config

import (
	"expvar"
	"github.com/kelseyhightower/envconfig"
)

type BotMode string

const (
	BotModeLongPoll BotMode = "LONGPOLL"
	BotModeWebhook  BotMode = "WEBHOOK"
)

type Config struct {
	Token         string  `envconfig:"TELEGRAM_TOKEN"`
	Mode          BotMode `envconfig:"TELEGRAM_BOT_MODE"`
	UpdateTimeout int     `envconfig:"TELEGRAM_LONGPOLL_UPDATE_TIMEOUT"`
	Address       string  `envconfig:"TELEGRAM_WEBHOOK_ADDRESS"`
	Debug         bool    `envconfig:"DEBUG"`
	BotURL        string  `envconfig:"BOT_URL"`
	SSLCert       string  `envconfig:"SSL_CERT"`
	SSLKey        string  `envconfig:"SSL_KEY"`
	ProxyProtocol string  `envconfig:"PROXY_PROTOCOL"`
	ProxyHost     string  `envconfig:"PROXY_HOST"`
	ProxyUser     string  `envconfig:"PROXY_USER"`
	ProxyPassword string  `envconfig:"PROXY_PASSWORD"`
}

// MustInitConfig инициализирует и возвращает конфиг иначе при ошибке кидает панику
func MustInitConfig() *Config {
	conf, err := InitConfig()
	if err != nil {
		panic(err)
	}

	return conf
}

// InitConfig инициализирует и возвращает конфиг
func InitConfig() (*Config, error) {
	conf := &Config{}

	err := envconfig.Process("", conf)
	if err != nil {
		return nil, err
	}

	expvar.Publish("config", expvar.Func(func() interface{} {
		return conf
	}))

	return conf, nil
}

// GetToken возвращает telegram-token бота
func (c *Config) GetToken() string { return c.Token }

// GetMode возвращает режим работы бота: long-poll или webhook
func (c *Config) GetMode() BotMode {
	if c.Mode == "" {
		return BotModeLongPoll
	}

	return c.Mode
}

// GetUpdateTimeout возвращает количество секунд для обновления режима long-poll
func (c *Config) GetUpdateTimeout() int {
	if c.UpdateTimeout <= 0 {
		return 60
	}

	return c.UpdateTimeout
}

// GetAddress возвращает адрес для запуска webhook-сервера
func (c *Config) GetAddress() string {
	if c.Address == "" {
		return "0.0.0.0:443"
	}

	return c.Address
}

// GetDebug возвращает флаг включения режима отладки
func (c *Config) GetDebug() bool { return c.Debug }

// GetBotURL возвращает url бота для режима webhook
func (c *Config) GetBotURL() string { return c.BotURL }

// GetSSLCert возвращает текст SSL-сертификата
func (c *Config) GetSSLCert() string { return c.SSLCert }

// GetSSLKey возвращает приватный ключ SSL-сертификата
func (c *Config) GetSSLKey() string { return c.SSLKey }

// GetProxyProtocol возвращает протокол прокси
func (c *Config) GetProxyProtocol() string { return c.ProxyProtocol }

// GetProxyURL возвращает url прокси
func (c *Config) GetProxyHost() string { return c.ProxyHost }

// GetProxyUser возвращает пользователя прокси
func (c *Config) GetProxyUser() string { return c.ProxyUser }

// GetProxyPassword возвращает пароль пользователя прокси
func (c *Config) GetProxyPassword() string { return c.ProxyPassword }
