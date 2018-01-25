package main

import (
	log "github.com/kirillDanshin/dlog"
	tg "github.com/toby3d/telegram"
)

func updateMessage(msg *tg.Message) {
	log.D(msg)
	if bot.IsMessageFromMe(msg) || bot.IsForwardFromMe(msg) {
		log.Ln("Ignore message update")
		return
	}

	switch {
	case bot.IsCommandToMe(msg):
		log.Ln("isCommandToMe")
		commands(msg)
	case msg.IsText():
		log.Ln("isCommandToMe")
		messages(msg)
	default:
		actions(msg)
	}
}
