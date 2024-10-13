package bot

import (
	"github.com/Ivlay/upstore/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type bot struct {
	*tgbotapi.BotAPI
	logger   *zap.Logger
	service  *service.Service
	commands map[commandKey]commandEntity
}

func New(logger *zap.Logger, service *service.Service, token string) (*bot, error) {
	logger = logger.Named("bot")

	api, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return nil, err
	}

	b := &bot{
		BotAPI:  api,
		service: service,
		logger:  logger,
	}

	if err := b.initCommands(); err != nil {
		return nil, err
	}

	b.logger.Info("bot created", zap.String("name", b.Self.UserName))

	return b, nil
}

func (b *bot) apiRequest(c tgbotapi.Chattable) error {
	_, err := b.Request(c)

	return err
}
