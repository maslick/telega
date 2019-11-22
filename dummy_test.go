package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Mock struct{}

func (m *Mock) SendTelegramMessage(token, chat, message string) ([]byte, error) {
	return []byte(message), nil
}

func TestHealthEndpoint(t *testing.T) {
	controller := RestController{&Mock{}}
	req, _ := http.NewRequest("GET", "/health", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.HealthHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "text/plain", rr.Header().Get("Content-Type"))
	assert.Equal(t, "UP", rr.Body.String())
}

func TestSendEndpoint(t *testing.T) {
	controller := RestController{&Mock{}}
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(Message{"chat", "hello world"})
	req, _ := http.NewRequest("POST", "/send", body)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SendHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, "hello world", rr.Body.String())
}
