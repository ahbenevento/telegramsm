package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

//  //  //

const CONFIG_FILENAME string = "telegramsm.conf"

func main() {
	if len(os.Args) == 1 {
		showHelp()
		return
	}

	cfg, err := loadConfig(CONFIG_FILENAME)

	if err != nil {
		exitWithError(err.Error())
	}

	fmt.Println(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(func(ctx context.Context, b *bot.Bot, update *models.Update) {
			fmt.Println(update.Message.Chat.ID)
			cancel()
		}),
	}

	b, err := bot.New(
		"5483339458:AAGa9EFtIpOWleiKCJ-c7FsJzzxrB-twzII",
		opts...,
	)

	b.Start(ctx)
}

func showHelp() {
	fmt.Print(`telegramsm bot-name chat-id message

Very simple example to send messages using Telegram bots.

`)
}

func exitWithError(error string) {
	fmt.Fprintf(os.Stderr, "Error: %s.\n", error)
	os.Exit(1)
}
