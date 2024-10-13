package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *bot) CheckPrice() {
	id, err := b.service.Product.Update()

	b.logger.Info("check price", zap.Int("id", id))
	if err != nil {
		b.logger.Error("error while update products", zap.Error(err))
	}

	if id > 0 {
		uu, err := b.service.User.GetAll()
		if err != nil {
			b.logger.Error("error while get all users", zap.Error(err))
			return
		}

		p, err := b.service.Product.Get()

		if uu != nil {
			defaultGreeting := "Привет! Цена на товар изменилась:\n"
			for _, u := range uu {
				product := fmt.Sprintf("%s - %d ₽, <s>%d ₽</s>\n", p.Title, p.Price, p.OldPrice)
				reply := tgbotapi.NewMessage(u.ChatId, defaultGreeting+product)
				reply.ParseMode = "html"
				if err := b.apiRequest(reply); err != nil {
					b.logger.Error("Failed to send start message", zap.Error(err))
				}
			}

		}
	}
}
