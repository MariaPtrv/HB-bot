package bot

import (
	"log"
	"net/url"
	"strconv"
)

func (b *Bot) SendMessage(chatId int64, text string) error {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatUint(uint64(chatId), 10))
	data.Add("text", text)
	err := b.httpClient.Post("sendMessage", data, nil)
	if err != nil {
		log.Println("Error occurred while sending message", err)
		return err
	}
	return nil
}

func (b *Bot) SendPhoto(chatId int64, text string) error {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatUint(uint64(chatId), 10))
	data.Add("caption", text)
	data.Add("photo", "https://i.ibb.co/4YNMtgm/present.png")
	err := b.httpClient.Post("sendPhoto", data, nil)
	if err != nil {
		log.Println("Error occurred while sending photo", err)
		return err
	}
	return nil
}
