package main

import (
	"bot/bot"
	"bot/http"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func abortOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	tgToken, exists := os.LookupEnv("TG_TOKEN")
	if !exists {
		abortOnError(fmt.Errorf("token cannot be empty"))
	}
	client, err := http.NewStdTelegramAPIClient(tgToken)
	abortOnError(err)
	bot := bot.NewBot(client)
	offset := int64(0)

	for {
		updates, err := bot.GetUpdates(offset)
		abortOnError(err)
		for _, update := range updates {
			// text := update.Message.Text
			// userFirstName := update.Message.From.FirstName
			// userId := update.Message.Chat.ID
			// response := "Hi " + userFirstName + "! You wrote: " + text
			// bot.SendMessage(userId, response)
			bot.ProcessUpdate(update)
			offset = update.ID + 1
		}
	}
}
