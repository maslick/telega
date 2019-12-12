package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Mock struct{}

func (m *Mock) SendTelegramMessage(message string) ([]byte, error) {
	return []byte(message), nil
}

func TestHealthEndpoint(t *testing.T) {
	server := RestController{&Mock{}}
	handler := http.HandlerFunc(server.HealthHandler)
	rr := performRequest(handler, "GET", "/health", nil)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "text/plain", rr.Header().Get("Content-Type"))
	assert.Equal(t, "UP", rr.Body.String())
}

func TestSendEndpoint(t *testing.T) {
	server := RestController{&Mock{}}
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(Message{"chat", "hello world"})

	handler := http.HandlerFunc(server.SendHandler)
	rr := performRequest(handler, "POST", "/send", body)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, "hello world", rr.Body.String())
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
