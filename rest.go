package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RestController struct {
	Telega ITelega
}

func (service *RestController) Start() {
	http.HandleFunc("/send", service.SendHandler)
	fmt.Println("Starting server on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (service *RestController) SendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post requests are allowed", 400)
		return
	}

	var message Request
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	TOKEN := GetEnv("TOKEN", "")
	CHAT_ID := GetEnv("CHAT_ID", "")

	resp, err := service.Telega.SendTelegramMessage(TOKEN, CHAT_ID, message.Text)
	if err != nil {
		http.Error(w, "Message delivery failed", 500)
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Fatal(err.Error())
	}
}
