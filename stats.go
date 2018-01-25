package main

import tg "github.com/toby3d/telegram"

type statistics struct {
	Users int // Total number of users

	Files    int // Total number of WebP
	Stickers int // Total number of stickers
	Sets     int // Total number of sets
}

func commandStats(msg *tg.Message) {
	T, err := switchLocale(msg.From.LanguageCode)
	errCheck(err)

	err = dbChangeUserState(msg.From.ID, stateNone)
	errCheck(err)

	_, err = bot.SendChatAction(msg.Chat.ID, tg.ActionTyping)
	errCheck(err)

	stats, err := dbGetPackStats(msg.From.ID)
	errCheck(err)

	text := T("reply_stats", map[string]interface{}{
		"Users":    stats.Users,
		"Files":    stats.Files,
		"Stickers": stats.Stickers,
		"Sets":     stats.Sets,
	})

	reply := tg.NewMessage(msg.Chat.ID, text)
	reply.ParseMode = tg.ModeMarkdown
	reply.ReplyMarkup = getMenuKeyboard(T)

	_, err = bot.SendMessage(reply)
	errCheck(err)
}
