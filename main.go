package main

import (
	"flag"
	"fmt"
	"os"
)

//  //  //

const CONFIG_FILENAME string = "telegramsm.conf"

// telegramsm bot-name chat-id|username message [...]
// telegramsm bot-name -chid [-s] [username]
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

	cfg, err := loadConfig(CONFIG_FILENAME)

	if err != nil {
		exitWithError(err.Error())
	}

	if len(pvalues) == 0 {
		if len(os.Args) >= 3 {
			if err := sendMessage(*cfg, os.Args[0], os.Args[1], os.Args[2:]...); err != nil {
				exitWithError(err.Error())
			}
		} else {
			showHelp()
		}

		return
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
		configUpdated, err := getLastChatIDFromMessage(cfg, pvalues[0], args.Arg(0))

		if err != nil {
			exitWithError(err.Error())
		}

		if saveChatID && configUpdated {
			if err := saveConfig(cfg, CONFIG_FILENAME); err != nil {
				exitWithError(err.Error())
			}
		}

		return
	}

	fmt.Println(pvalues, os.Args)
}

func showHelp() {
	fmt.Print(`Very simple example to send messages using Telegram bots.

USES:

telegramsm bot-name chat-id|username message [...]

  Send message to chat ID or username (stored in config file).

telegramsm -chid bot-name [-s] [username]

  Get a chat ID from last message recieved for the bot. Use "-s" to save in
  config file.
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
