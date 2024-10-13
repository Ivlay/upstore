package bot

import (
	"fmt"
	"log"

	"github.com/Ivlay/upstore"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *bot) CmdStart(upd tgbotapi.Update) {
	name := upd.Message.From.UserName

	if name == "" {
		name = upd.Message.From.FirstName
	}

	u := upstore.User{
		ChatId:    upd.Message.Chat.ID,
		FirstName: upd.Message.Chat.FirstName,
		UserId:    int(upd.Message.From.ID),
		UserName:  name,
	}

	_, err := b.service.User.FindOrCreate(u)
	if err != nil {
		log.Fatal(err.Error())
	}

	p, pErr := b.service.Product.Get()
	if pErr != nil {
		b.logger.Error("Filed to get product", zap.Error(err))
	}

	message := "Добро пожаловать в <b>UpStore price check</b>, %s!\n\nТекущая цена на <em>%s</em>:<strong>\n\n%d ₽</strong>"

	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message, name, p.Title, p.Price))
	reply.ParseMode = "html"

	// keyboard := tgbotapi.NewReplyKeyboard(
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton(string(ReplyProducts)),
	// 	),
	// )

	reply.DisableWebPagePreview = true

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("Failed to send start message", zap.Error(err))
	}
}
