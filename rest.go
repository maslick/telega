package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type RestController struct {
	Telega ITelega
}

func (service *RestController) Start() {
	_, ok1 := os.LookupEnv("USERNAME")
	_, ok2 := os.LookupEnv("PASSWORD")
	if !ok1 && !ok2 {
		http.HandleFunc("/send", service.SendHandler)
	} else {
		http.HandleFunc("/send", basicAuth(service.SendHandler))
	}

	http.HandleFunc("/health", service.HealthHandler)

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func getPort() string {
	var port = getEnv("PORT", "8080")
	return ":" + port
}

func (service *RestController) SendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", 400)
		return
	}

	var message Request
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	TOKEN := getEnv("TOKEN", "")
	CHAT_ID := getEnv("CHAT_ID", "")

	resp, err := service.Telega.SendTelegramMessage(TOKEN, CHAT_ID, message.Text)
	if err != nil {
		http.Error(w, "Message delivery failed: "+err.Error(), 500)
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (service *RestController) HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", 400)
		return
	}
	_, _ = w.Write([]byte("UP"))
}
