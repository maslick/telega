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
	err := json.NewEncoder(body).Encode(Message{chat, message})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json; charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	text, _ := ioutil.ReadAll(resp.Body)
	return text, err
}
