package main

import (
	"fmt"
	"time"

	tg "github.com/toby3d/telegram" // My Telegram bindings
)

func commandReset(msg *tg.Message) {
	T, err := switchLocale(msg.From.LanguageCode)
	errCheck(err)

	_, total, err := dbGetUserStickers(msg.From.ID, 0, "")
	errCheck(err)

	_, err = bot.SendChatAction(msg.Chat.ID, tg.ActionTyping)
	errCheck(err)

	if total <= 0 {
		err = dbChangeUserState(msg.From.ID, stateNone)
		errCheck(err)

		reply := tg.NewMessage(msg.Chat.ID, T("error_already_reset"))
		reply.ParseMode = tg.ModeMarkdown
		reply.ReplyMarkup = getMenuKeyboard(T)

		_, err = bot.SendMessage(reply)
		errCheck(err)
		return
	}

	err = dbChangeUserState(msg.From.ID, stateReset)
	errCheck(err)

	reply := tg.NewMessage(
		msg.Chat.ID,
		T("reply_reset", map[string]interface{}{
			"KeyPhrase":     T("meta_key_phrase"),
			"CancelCommand": cmdCancel,
		}))
	reply.ParseMode = tg.ModeMarkdown
	reply.ReplyMarkup = getCancelButton(T)

	_, err = bot.SendMessage(reply)
	errCheck(err)
}

func actionReset(msg *tg.Message) {
	T, err := switchLocale(msg.From.LanguageCode)
	errCheck(err)

	err = dbChangeUserState(msg.From.ID, stateNone)
	errCheck(err)

	_, err = bot.SendChatAction(msg.Chat.ID, tg.ActionTyping)
	errCheck(err)

	if msg.Text != T("meta_key_phrase") {
		reply := tg.NewMessage(msg.Chat.ID, T("error_reset_phrase"))
		reply.ParseMode = tg.ModeMarkdown
		reply.ReplyMarkup = getMenuKeyboard(T)

		_, err = bot.SendMessage(reply)
		errCheck(err)
		return
	}

	err = dbResetUserStickers(msg.From.ID)
	errCheck(err)

	reply := tg.NewMessage(msg.Chat.ID, T("success_reset"))
	reply.ParseMode = tg.ModeMarkdown
	reply.ReplyMarkup = tg.NewReplyKeyboardRemove(false)

	_, err = bot.SendMessage(reply)
	errCheck(err)

	for i := 1; i <= 3; i++ {
		_, err = bot.SendChatAction(msg.Chat.ID, tg.ActionTyping)
		errCheck(err)

		text := T(fmt.Sprint("meta_reset_", i))

		time.Sleep(time.Minute * time.Duration(len(text)) / 1000)

		reply = tg.NewMessage(msg.Chat.ID, text)
		reply.ParseMode = tg.ModeMarkdown

		_, err = bot.SendMessage(reply)
		errCheck(err)
	}
}
