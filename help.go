package main

import (
	log "github.com/kirillDanshin/dlog" // Insert logs only in debug builds
	"github.com/toby3d/go-telegram"     // My Telegram bindings
)

func commandHelp(msg *telegram.Message) {
	log.Ln("Received a /help command")
	bot.SendChatAction(msg.Chat.ID, telegram.ActionTyping)

	err := dbChangeUserState(msg.From.ID, stateNone)
	errCheck(err)

	T, err := switchLocale(msg.From.LanguageCode)
	errCheck(err)

	markup := telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonSwitch(
				T("button_share"),
				" ",
			),
		),
	)

	reply := telegram.NewMessage(
		msg.Chat.ID, T("reply_help", map[string]interface{}{
			"AddStickerCommand": cmdAddSticker,
			"AddPackCommand":    cmdAddPack,
			"DeleteCommand":     cmdDelete,
			"ResetCommand":      cmdReset,
			"CancelCommand":     cmdCancel,
			"Username":          bot.Self.Username,
		}),
	)
	reply.ParseMode = telegram.ModeMarkdown
	reply.ReplyMarkup = &markup

	_, err = bot.SendMessage(reply)
	errCheck(err)
}
