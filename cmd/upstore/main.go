//go:build !cgo

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ivlay/upstore/pkg/bot"
	htmlParser "github.com/Ivlay/upstore/pkg/parser"
	"github.com/Ivlay/upstore/pkg/repository"
	"github.com/Ivlay/upstore/pkg/service"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load()

	logger := prepareLogger()

	defer logger.Sync()

	cron := cron.New(cron.WithLogger(cron.DefaultLogger))

	logger.Info("App Started", zap.String("mode", os.Getenv("APP_ENV")))

	db, dbErr := repository.NewSqLiteDB(logger, os.Getenv("DB_PATH"))
	if dbErr != nil {
		log.Fatal(dbErr.Error())
	}

	repository := repository.New(logger, db)

	parser := htmlParser.New("https://upstore24.ru/product/macbookair15_m3_10gpu_8-512gb_midnight")

	fmt.Printf("%+v\n", parser.PrepareProduct())

	service := service.New(logger, repository, parser)

	bot, err := bot.New(logger, service, os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err.Error())
	}

	service.Product.Prepare()

	go bot.Run()

	bot.CheckPrice()

	cron.AddFunc("@every 1h", bot.CheckPrice)

	cron.Start()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	<-stopCh

	logger.Info("Bot stopped")
}

func prepareLogger() *zap.Logger {
	logger := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}

	return logger
}
