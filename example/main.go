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