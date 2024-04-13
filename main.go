package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	traq "github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	payload "github.com/traPtitech/traq-ws-bot/payload"
)

func main() {
	// .env ファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数を取得
	accessToken := os.Getenv("TRAQ_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("TRAQ_ACCESS_TOKEN is not set")
	}

	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: accessToken,
	})
	if err != nil {
		panic(err)
	}

	bot.OnMessageCreated(func(p *payload.MessageCreated) {
		log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
		_, _, err := bot.API().
			MessageApi.
			PostMessage(context.Background(), p.Message.ChannelID).
			PostMessageRequest(traq.PostMessageRequest{
				Content: "oisu-",
			}).
			Execute()
		if err != nil {
			log.Println(err)
		}
	})

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
