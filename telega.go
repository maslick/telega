package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ITelega interface {
	SendTelegramMessage(token, chat, message string) ([]byte, error)
}

type Telega struct{}

func (t *Telega) SendTelegramMessage(token, chat, message string) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(Message{chat, message})
	resp, err := http.Post(url, "application/json; charset=utf-8", body)
	text, _ := ioutil.ReadAll(resp.Body)
	return text, err
}
