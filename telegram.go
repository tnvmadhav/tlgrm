package tlgrm

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)


// TelegramClient is a blueprint to generate a (reusable) client to ping users as a telegram bot
// from a golang program.
type TelegramClient struct {
	BotToken   string
	ApiBaseUrl string
}


func (t *TelegramClient) apiRunner(url string, data *bytes.Buffer) error {

	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}


// SendPhoto simply sends an image along with text to the given user. The image 
// has to be hosted elsewhere on the internet and it's URL has to be sent as an
// argument. This method returns an error if encountered.
func (t *TelegramClient) SendPhoto(msg, photoURL, userChatID string) error {

	urlStr := fmt.Sprintf("%s/bot%s/sendPhoto", t.ApiBaseUrl, t.BotToken)

	data := url.Values{}
	data.Set("chat_id", userChatID)
	data.Set("photo", photoURL)
	data.Set("caption", msg)

	return t.apiRunner(urlStr, bytes.NewBufferString(data.Encode()))
}


// SendMessage simply sends a given text to the given user.
// This method returns an error if encountered.
func (t *TelegramClient) SendMessage(msg, userChatID string) error {

	urlStr := fmt.Sprintf("%s/bot%s/sendMessage", t.ApiBaseUrl, t.BotToken)

	data := url.Values{}
	data.Set("chat_id", userChatID)
	data.Set("text", msg)

	return t.apiRunner(urlStr, bytes.NewBufferString(data.Encode()))
}
