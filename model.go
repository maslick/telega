package main

type Request struct {
	Text string `json:"text"`
}

type Message struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}
