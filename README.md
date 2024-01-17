# tlgrm

A simple go package to send telegram bot pings from your go programs. zero dependency.


## Usage

To send a message as a telegram bot you need two main things:

1. The bot's token
2. The recipient user's chat ID

If you're not sure what these are, you can learn on how to get them from this [guide](https://core.telegram.org/bots/tutorial#obtain-your-bot-token)

Once you get these, all you have to do is plug them in your go program like the one below:

```go
package main

import (
	"os"
	"github.com/tnvmadhav/tlgrm"
)


func main(){

	// init a client
	tlgrmClient := tlgrm.TelegramClient{
		ApiBaseUrl: "https://api.telegram.org",
		BotToken:   os.Getenv("TELEGRAM_BOT_TOKEN"),
	}

	// send a simple text ping
	tlgrmClient.SendMessage("Hey!", os.Getenv("TELEGRAM_CHAT_ID"))

	// send an image with text caption
	tlgrmClient.SendPhoto("This is a caption", "https://tnvmadhav.me/pfp.png", os.Getenv("TELEGRAM_CHAT_ID"))

}
```