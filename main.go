package main

import (
	"fmt"
	"os"
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
}

func showHelp() {
	fmt.Print(`telegramsm alias-bot chat-id message

Very simple example to send messages using Telegram bots.

`)
}

func exitWithError(error string) {
	fmt.Fprintf(os.Stderr, "Error: %s.\n", error)
	os.Exit(1)
}
