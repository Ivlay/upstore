package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type commandEntity struct {
	key         commandKey
	description string
	action      func(upd tgbotapi.Update)
}

const (
	StartCmdKey          = commandKey("start")
	MySubscriptionCmdKey = commandKey("subscription")
	ProductKey           = commandKey("products")
	OriginalSite         = commandKey("site")
)

type commandKey string

func (b *bot) initCommands() error {
	b.commands = make(map[commandKey]commandEntity)

	commands := []commandEntity{
		{
			key:         StartCmdKey,
			description: "Запустить бота",
			action:      b.CmdStart},
	}

	tgCommands := make([]tgbotapi.BotCommand, 0, len(commands))
	for _, cmd := range commands {
		b.commands[cmd.key] = cmd

		tgCommands = append(tgCommands, tgbotapi.BotCommand{
			Command:     "/" + string(cmd.key),
			Description: cmd.description,
		})

		config := tgbotapi.NewSetMyCommands(tgCommands...)

		return b.apiRequest(config)

	}

	return nil
}
