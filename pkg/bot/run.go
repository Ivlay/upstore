package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

const (
	productPrefix   = "Товары"
	analyticsPrefix = "Аналитика"
)

func (b *bot) Run() {
	updetesCfg := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 10,
	}

	for upd := range b.GetUpdatesChan(updetesCfg) {
		go b.processUpdate(upd)
	}
}

func (b *bot) processUpdate(upd tgbotapi.Update) {
	if upd.Message != nil {
		if upd.Message.IsCommand() {
			key := upd.Message.Command()
			if cmd, ok := b.commands[commandKey(key)]; ok {
				cmd.action(upd)
			} else {
				b.logger.Error("command handler not found", zap.String("cmd", key))
			}
			return
		}
	}
}
