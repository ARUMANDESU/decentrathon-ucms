package tgbot

import tb "gopkg.in/telebot.v3"

var (
	EmptyButtonInline = tb.InlineButton{Text: " ", Data: ""}
	StartButtonInline = tb.InlineButton{Text: "/start", Data: "start"}
	HelpButtonInline  = tb.InlineButton{Text: "/help", Data: "help"}
)
