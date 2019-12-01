package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	server := RestController{&Telega{}}
	server.Start()
}

///////////////////////
// Model
//////////////////////

type Request struct {
	Text string `json:"text"`
}

type Message struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

///////////////////////
// Telegram send
//////////////////////

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

///////////////////////
// Controller
//////////////////////

type RestController struct {
	service ITelega
}

func (this *RestController) Start() {
	if useAuth() {
		http.HandleFunc("/send", basicAuth(this.SendHandler))
	} else {
		http.HandleFunc("/send", this.SendHandler)
	}

	http.HandleFunc("/health", this.HealthHandler)

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func (this *RestController) SendHandler(w http.ResponseWriter, r *http.Request) {
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

	token := getEnv("BOT_TOKEN", "")
	chatId := getEnv("CHAT_ID", "")

	resp, err := this.service.SendTelegramMessage(token, chatId, message.Text)
	if err != nil {
		http.Error(w, "Message delivery failed: "+err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}

func (_ *RestController) HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", 400)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("UP"))
}

///////////////////////
// Helper functions
//////////////////////

func useAuth() bool {
	_, usernameOk := os.LookupEnv("USERNAME")
	_, passwordOk := os.LookupEnv("PASSWORD")
	return usernameOk && passwordOk
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getPort() string {
	var port = getEnv("PORT", "8080")
	return ":" + port
}

type handler func(w http.ResponseWriter, r *http.Request)

func basicAuth(pass handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validateUsernamePassword(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		pass(w, r)
	}
}

func validateUsernamePassword(username, password string) bool {
	if username == getEnv("USERNAME", "") && password == getEnv("PASSWORD", "") {
		return true
	}
	return false
}
