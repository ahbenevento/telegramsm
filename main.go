package main

import (
	"flag"
	"fmt"
	"os"
)

//  //  //

const CONFIG_FILENAME string = "telegramsm.conf"

// telegramsm nombre-bot id-chat|usuario mensaje [mensaje ...]
// telegramsm -chid nombre-bot [-s] [usuario]
func main() {
	if len(os.Args) == 1 {
		showHelp()
		return
	}

	cfg, err := loadConfig(CONFIG_FILENAME)

	if err != nil {
		exitWithError(err.Error())
	}

	os.Args = os.Args[1:]

	if os.Args[0][0] != '-' {
		// Enviar mensaje
		if len(os.Args) >= 3 {
			if err := sendMessage(*cfg, os.Args[0], os.Args[1], os.Args[2:]...); err != nil {
				exitWithError(err.Error())
			}
		} else {
			showHelp()
		}

		return
	}

	// Obtener ID de chat
	var (
		getChatID  string
		saveChatID bool
	)

	args := flag.NewFlagSet("telegramsm", flag.ContinueOnError)

	args.StringVar(&getChatID, "chid", "", "Obtener el ID del chat.")
	args.BoolVar(&saveChatID, "s", false, "Guardar los datos en el archivo de configuración.")

	if err := args.Parse(os.Args); err != nil {
		exitByInvalidArguments(err.Error())
	} else if getChatID != "" {
		configUpdated, err := getLastChatIDFromMessage(cfg, getChatID, args.Arg(0))

		if err != nil {
			exitWithError(err.Error())
		}

		if saveChatID && configUpdated {
			if err := saveConfig(cfg, CONFIG_FILENAME); err != nil {
				exitWithError(err.Error())
			}
		}
	}
}

func showHelp() {
	fmt.Print(`Ejemplo muy simple para el envío de mensajes utilizando un Bot de Telegram.

USOS:

telegramsm nombre-bot id-chat|usuario mensaje [mensaje ...]

  Envía un mensaje utiliando el ID del chat o nombre de usuario (guardado en el
  archivo de configuración).

telegramsm -chid nombre-bot [-s] [usuario]

  Obtiene el ID del último mensaje enviado al bot. Usa el parámetro "-s" para
  guardar en el archivo de configuración la combinación de nombre de usuario e
  ID del chat.

  Si se establece un valor para "usuario" este reemplazará al definido en
  Telegram.
`)
}

func exitWithError(error string) {
	colors.clrError.Fprintf(os.Stderr, "Error: %s\n", error)
	os.Exit(1)
}

func exitByInvalidArguments(error string) {
	fmt.Fprintf(os.Stderr, "%s\nEscribe \"telegramsm\" sin parámetros para obtener ayuda.\n", colors.clrError.Sprintf("Error: %s", error))
	os.Exit(2)
}
