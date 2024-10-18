package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/ARUMANDESU/decentrathon-ucms/backend/internal/config"
	"github.com/ARUMANDESU/decentrathon-ucms/backend/internal/tgbot"
	"github.com/ARUMANDESU/decentrathon-ucms/backend/pkg/logs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}

	cfg := config.MustLoad()

	log, teardown := logs.Setup(cfg.Env)
	defer teardown()

	log.Info("starting the telegram bot", slog.String("env", cfg.Env))

	bot, err := tgbot.NewBot(cfg.Telegram, log)
	if err != nil {
		log.Error("failed to create telegram bot", logs.Err(err))
		return
	}

	go func() {
		err := bot.Start()
		if err != nil {
			log.Error("failed to start telegram bot", logs.Err(err))
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	defer log.Info("telegram bot stopped", slog.String("signal", sign.String()))
	log.Info("stopping telegram bot", slog.String("signal", sign.String()))
}
