package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebSocketServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HandleClients))
	defer server.Close()

	u := "ws" + strings.TrimPrefix(server.URL, "http")

	socket, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer socket.Close()

	m := Message{Message: "hello"}

	if err := socket.WriteJSON(&m); err != nil {
		t.Fatalf("%v", err)
	}

	var message Message
	err = socket.ReadJSON(&message)
	if err != nil {
		t.Fatalf("%v", err)
	}
	assert.Equal(t, "hello", message.Message, "they should be equal")
}
