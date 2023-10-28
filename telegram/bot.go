package telegram

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Println(update.Message.Text)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}

func BotInit() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("6811179663:AAFiezciE4I62ml4AplNcL6_jcSJa0Y-rio", opts...)
	if err != nil {
		panic(err)
	}
	log.Println("Starting telegram bot...")
	b.Start(ctx)
}
