# Ejemplo muy simple para el envío de mensajes utilizando un Bot de Telegram

Paquetes utilizados:

- [fatih/color](https://github.com/fatih/color)
- [go-telegram/bot](https://github.com/go-telegram/bot)

## Usos

### Enviar mensaje

```console
telegramsm nombre-bot id-chat|usuario mensaje [mensaje ...]
```

- **nombre-bot** El nombre del bot guardado en el archivo de configuración.
- **id-chat** Identificador del chat.
- **usuario** El nombre de usuario o alias guardado en el archivo de configuración.
- **message** Uno o más mensajes a enviar.

### Obtener ID de chat

Imprime y guarda el identificador de chat del último mensaje enviado al bot de
Telegram.

```console
telegramsm -chid nombre-bot [-s] [usuario]
```

- **nombre-bot** El nombre del bot guardado en el archivo de configuración.
- **-s** Guarda en el archivo de configuración la combinación de nombre de usuario e identificador.
- **usuario** El nombre de usuario a guardar en el archivo de configuración (reemplaza al definido en Telegram).
