Bot required env:
- TELEGRAM_TOKEN
- DATABASE_DSN
- BOT_URL (for webhook mode)

TELEGRAM_BOT_MODE:
- LONGPOLL
- WEBHOOK

Long poll env:
- TELEGRAM_LONGPOLL_UPDATE_TIMEOUT (default: 60 seconds)

Webhook env:
- TELEGRAM_WEBHOOK_ADDRESS (default: 0.0.0.0:80)

Addicted env:
- DEBUG (default: false)
- SSL_CERT (for webhook mode)
- SSL_KEY (for webhook mode)
- PROXY_PROTOCOL (example: http)
- PROXY_HOST (example: google.com:80)
- PROXY_USER (example: anonymous)
- PROXY_PASSWORD (example: password)
