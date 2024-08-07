package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

//  //  //

func getLastChatIDFromMessage(botName, token, username string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	var result error

	funcPrintMessage := func() {
		fmt.Printf("Please, send a message to the bot \"%s\"...\n", colors.clrHighlighted.Sprint(botName))
	}
	opts := []bot.Option{
		bot.WithDefaultHandler(func(ctx context.Context, b *bot.Bot, update *models.Update) {
			if update.Message != nil {
				if username == "" && update.Message.From.Username != "" {
					username = update.Message.From.Username
				}

				if username != "" {
					fmt.Printf(
						"User %s (%s): %s\n",
						colors.clrBold.Sprint(strings.TrimSpace(strings.Join([]string{update.Message.From.FirstName, update.Message.From.LastName}, " "))),
						colors.clrHighlighted.Sprint(username),
						colors.clrHighlighted.Sprintf("%d", update.Message.From.ID),
					)
				} else {
					result = errors.New("must be define username")
				}

				cancel()
			} else {
				funcPrintMessage()
			}
		}),
	}
	b, result := bot.New(token, opts...)

	if result == nil {
		funcPrintMessage()
		b.Start(ctx)
	}

	return result
}
