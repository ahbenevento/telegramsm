# Send messages to Telegram using bots

Very simple example to send messages using Telegram bots.

Packages used:

- [fatih/color](https://github.com/fatih/color)
- [go-telegram/bot](https://github.com/go-telegram/bot)

## Uses

### Send message

```console
telegramsm bot-name chat-id|username message [...]
```

- **bot-name** The name of the bot stored in config file.
- **chat-id** Number of chat identify.
- **username** Nickname or alias stored in config file.
- **message** One o more messages to send.

### Get chat ID

```console
telegramsm -chid bot-name [-s] [username]
```

- **bot-name** The name of the bot stored in config file.
- **-s** Save data in config file.
- **username** User's nickname to save in config file (replaces the username defined in Telegram).
