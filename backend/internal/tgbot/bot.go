package tgbot

import (
	"log/slog"
	"net/http"
	"strings"

	tb "gopkg.in/telebot.v3"

	"github.com/ARUMANDESU/decentrathon-ucms/backend/internal/config"
	"github.com/ARUMANDESU/decentrathon-ucms/backend/pkg/logs"
)

type Bot struct {
	cfg        config.Telegram
	log        *slog.Logger
	bot        *tb.Bot
	httpServer *http.Server
}

func NewBot(cfg config.Telegram, logger *slog.Logger) (*Bot, error) {
	const op = "tgbot.NewBot"
	log := logger.With("op", op)

	httpServer := NewHTTPServer(cfg)
	log.Info("http server started", slog.String("addr", cfg.WebhookURL))

	webhook := &tb.Webhook{
		Listen:   cfg.URL,
		Endpoint: &tb.WebhookEndpoint{PublicURL: cfg.WebhookURL},
	}

	log.Info(
		"webhook created",
		slog.String("webhook", cfg.WebhookURL),
		slog.String("url", cfg.URL),
	)

	spamProtected := tb.NewMiddlewarePoller(webhook, func(upd *tb.Update) bool {
		if upd.Message == nil {
			return true
		}

		if strings.Contains(upd.Message.Text, "spam") {
			return false
		}

		return true
	})

	b, err := tb.NewBot(tb.Settings{
		Token:  cfg.Token,
		Poller: spamProtected,
	})
	if err != nil {
		log.Error("failed to create new bot", logs.Err(err))
		return nil, err
	}
	log.Debug("bot created", slog.String("token", cfg.Token)) // @@TODO: remove later , this is for debugging

	return &Bot{
		bot:        b,
		cfg:        cfg,
		log:        logger,
		httpServer: httpServer,
	}, nil
}

func NewHTTPServer(cfg config.Telegram) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	httpServer := &http.Server{
		Addr:    cfg.WebhookURL,
		Handler: mux,
	}

	go httpServer.ListenAndServe()

	return httpServer
}

func (b *Bot) AddHandlers() {
	b.bot.Handle("/start", b.handleStartCommand)
	b.bot.Handle("/help", b.handleHelpCommand)
}

func (b *Bot) Start() error {
	b.AddHandlers()
	b.bot.Start()

	return nil
}

func (b *Bot) Stop() {
	b.bot.Stop()
	err := b.httpServer.Close()
	if err != nil {
		b.log.Error("failed to close http server", logs.Err(err))
		return
	}
}
