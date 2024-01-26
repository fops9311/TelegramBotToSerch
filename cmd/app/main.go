package main

import (
	hotelsfourprovider "TelegramBotToSerch/internal/hotels-four-provider"
	telegrambotrequester "TelegramBotToSerch/internal/telegram-bot-requester"
	service "TelegramBotToSerch/internal/telegram-bot-search-service"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	service.NewService(
		service.Opts{
			Logger:   logger,
			Provider: hotelsfourprovider.NewHotelFourProvider(),
		},
	).WithRequester(&telegrambotrequester.TelegramBotRequester{}).Run()
}
