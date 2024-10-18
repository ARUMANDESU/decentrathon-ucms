package tgbot

import (
	"fmt"
	"strings"

	tb "gopkg.in/telebot.v3"

	"github.com/ARUMANDESU/decentrathon-ucms/backend/pkg/logs"
)

func (b *Bot) handleStartCommand(ctx tb.Context) error {
	const op = "tgbot.Bot.handler.handleStartCommand"
	log := b.log.With("op", op)

	err := ctx.Send(
		fmt.Sprintf("Salam %s, welcome to the bot!", ctx.Sender().FirstName),
		&tb.SendOptions{
			ParseMode: tb.ModeMarkdown,
		},
	)
	if err != nil {
		log.Error("failed to send message", logs.Err(err))
		return err
	}
	return nil
}

func (b *Bot) handleHelpCommand(ctx tb.Context) error {
	const op = "tgbot.Bot.handler.handleHelpCommand"
	log := b.log.With("op", op)

	log.Debug("Help command received")

	var helpMessage strings.Builder
	helpMessage.WriteString("Here are the available commands:\n")
	helpMessage.WriteString("/start - Start the bot\n")
	helpMessage.WriteString("/help - Show help message\n")

	err := ctx.Send(helpMessage.String(),
		&tb.SendOptions{
			ParseMode: tb.ModeMarkdown,
		},
	)
	if err != nil {
		log.Error("failed to send message", logs.Err(err))
		return err
	}

	return nil
}
