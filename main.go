package main

import (
	"flag"
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

	os.Args = os.Args[1:]
	pvalues := []string{}

	for i, arg := range os.Args {
		if arg[0] == '-' {
			pvalues = os.Args[0:i]
			os.Args = os.Args[i:]
			break
		}
	}

	if len(pvalues) == 0 {
		showHelp()
		return
	}

	cfg, err := loadConfig(CONFIG_FILENAME)

	if err != nil {
		exitWithError(err.Error())
	}

	botToken, ok := cfg.Bots[pvalues[0]]

	if !ok {
		exitWithError(fmt.Sprintf("bot not found with the name: \"%s\"", pvalues[0]))
	}

	var (
		getChatID  bool
		saveChatID bool
	)

	args := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	args.BoolVar(&getChatID, "chid", false, "Get chat ID")
	args.BoolVar(&saveChatID, "s", false, "Save the chat ID retrieved to config file")

	if err := args.Parse(os.Args); err != nil {
		exitByInvalidArguments(err.Error())
	} else if getChatID {
		getLastChatIDFromMessage(pvalues[0], botToken, args.Arg(0))
	}
}

func showHelp() {
	fmt.Print(`Very simple example to send messages using Telegram bots.

USES:

telegramsm bot-name chat-id|username message

  Send message to chat ID or username (config file).

telegramsm bot-name -chid [-s] [username]

  Get a chat ID from last message recieved for the bot.
`)
}

func exitWithError(error string) {
	colors.clrError.Fprintf(os.Stderr, "Error: %s\n", error)
	os.Exit(1)
}

func exitByInvalidArguments(error string) {
	fmt.Fprintf(os.Stderr, "%s\nType \"telegramsm\" without parameters for help.\n", colors.clrError.Sprintf("Error: %s", error))
	os.Exit(2)
}
