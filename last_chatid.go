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

func getLastChatIDFromMessage(cfg *appConfig, botName, username string) (configUpdated bool, err error) {
	token, err := getBotTokenByName(cfg.Bots, botName)

	if err != nil {
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	funcPrintMessage := func() {
		fmt.Printf("Por favor, envía un mensaje al bot \"%s\"...\n", colors.clrHighlighted.Sprint(botName))
	}
	opts := []bot.Option{
		bot.WithDefaultHandler(func(ctx context.Context, b *bot.Bot, update *models.Update) {
			if update.Message != nil {
				if username == "" && update.Message.From.Username != "" {
					username = update.Message.From.Username
				}

				if username != "" {
					fmt.Printf(
						"Usuario %s (%s): %s\n",
						colors.clrBold.Sprint(strings.TrimSpace(strings.Join([]string{update.Message.From.FirstName, update.Message.From.LastName}, " "))),
						colors.clrHighlighted.Sprint(username),
						colors.clrHighlighted.Sprintf("%d", update.Message.From.ID),
					)

					// Buscar el usuario en la lista configurada
					if cfgUsername, ok := cfg.Users[update.Message.From.ID]; !ok || cfgUsername != username {
						cfg.Users[update.Message.From.ID] = username
						configUpdated = true
					}
				} else {
					err = errors.New("must be define username")
				}

				cancel()
			} else {
				funcPrintMessage()
			}
		}),
	}
	b, err := bot.New(token, opts...)

	if err == nil {
		funcPrintMessage()
		b.Start(ctx)
	}

	return
}
