// main.go

package main

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type BotContext struct {
	SelectedCity string
}

type LocationInfo struct {
	Cities []string `json:"cities"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	apiToken := os.Getenv("RAPID_API_KEY")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	contexts := make(map[int64]*BotContext)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Привет! Я твой бот."
			case "lowprice":
				context, exists := contexts[update.Message.Chat.ID]
				if !exists {
					context = &BotContext{}
					contexts[update.Message.Chat.ID] = context
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В какой город вы хотите поехать?")
				_, err := bot.Send(msg)
				if err != nil {
					return
				}

			case "help":
				msg.Text = "Доступные команды: \n/start - Начать работу\n/help - помощь"
			default:
				msg.Text = "Неизвестная команда!"
			}
			_, err := bot.Send(msg)
			if err != nil {
				return
			}
		} else {
			context, exists := contexts[update.Message.Chat.ID]
			if exists {
				context.SelectedCity = update.Message.Text

				locationInfo, err := getLocation(context.SelectedCity, apiToken)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "error")
					_, err := bot.Send(msg)
					if err != nil {
						return
					}
					log.Fatal(err)
				}

				locationResponse := LocationInfo{
					Cities: locationInfo,
				}

				responseJSON, err := json.Marshal(locationResponse)
				if err != nil {
					log.Fatal(err)
				}

				messageText := string(responseJSON)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
				_, err = bot.Send(msg)
				if err != nil {
					return
				}

				delete(contexts, update.Message.Chat.ID)
			}
		}
	}
}
