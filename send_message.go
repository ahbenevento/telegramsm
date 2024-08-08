package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

//  //  //

func sendMessage(cfg appConfig, botName, username string, messages ...string) error {
	token, err := getBotTokenByName(cfg.Bots, botName)

	if err != nil {
		return err
	}

	userID, err := getUserID(cfg.Users, username)

	if err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	b, err := bot.New(token)

	if err != nil {
		return err
	}

	for _, msg := range messages {
		smParams := &bot.SendMessageParams{
			ChatID: userID,
			Text:   msg,
		}

		if _, err = b.SendMessage(ctx, smParams); err != nil {
			return err
		}
	}

	return nil
}
