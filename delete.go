package main

import (
	log "github.com/kirillDanshin/dlog" // Insert logs only in debug builds
	tg "github.com/toby3d/telegram"     // My Telegram bindings
)

func commandDelete(msg *tg.Message, pack bool) {
	bot.SendChatAction(msg.Chat.ID, tg.ActionTyping)

	T, err := switchLocale(msg.From.LanguageCode)
	errCheck(err)

	_, total, err := dbGetUserStickers(msg.From.ID, 0, "")
	errCheck(err)

	if total <= 0 {
		err = dbChangeUserState(msg.From.ID, stateNone)
		errCheck(err)

		reply := tg.NewMessage(msg.Chat.ID, T("error_empty_del"))
		_, err = bot.SendMessage(reply)
		errCheck(err)
		return
	}

	reply := tg.NewMessage(msg.Chat.ID, T("reply_del_sticker"))
	reply.ParseMode = tg.ModeMarkdown

	err = dbChangeUserState(msg.From.ID, stateDeleteSticker)
	errCheck(err)

	if pack {
		reply.Text = T("reply_del_pack")

		err = dbChangeUserState(msg.From.ID, stateDeletePack)
		errCheck(err)
	}

	markup := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonSwitchSelf(
				T("button_del"),
				" ",
			),
		),
	)
	reply.ReplyMarkup = &markup

	_, err = bot.SendMessage(reply)
	errCheck(err)
}

func actionDelete(msg *tg.Message, pack bool) {
	bot.SendChatAction(msg.Chat.ID, tg.ActionTyping)

	T, err := switchLocale(msg.From.LanguageCode)
	errCheck(err)

	reply := tg.NewMessage(msg.Chat.ID, T("success_del_sticker"))
	reply.ParseMode = tg.ModeMarkdown

	var notExist bool
	if pack {
		set, err := bot.GetStickerSet(msg.Sticker.SetName)
		errCheck(err)

		log.Ln("SetName:", set.Title)
		reply.Text = T("success_del_pack", map[string]interface{}{
			"SetTitle": set.Title,
		})

		notExist, err = dbDeletePack(msg.From.ID, msg.Sticker.SetName)
	} else {
		notExist, err = dbDeleteSticker(msg.From.ID, msg.Sticker.SetName, msg.Sticker.FileID)
	}
	errCheck(err)

	if notExist {
		reply.Text = T("error_already_del")
	}

	_, err = bot.SendMessage(reply)
	errCheck(err)
}
