package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"os"
	"strings"
)

type Client struct {
	WS       *websocket.Conn
	ClientID string 
}

func NewClient(ws *websocket.Conn, clientID string) *Client {
	return &Client{
		WS:       ws,
		ClientID: clientID,
	}
}

func (c *Client) receiveMessages() {
	for {
		var msg string
		if err := websocket.Message.Receive(c.WS, &msg); err != nil {
			log.Fatal("Error receiving message:", err)
		}

		if !strings.HasPrefix(msg, c.ClientID+": ") {
			log.Println(msg)
		}
	}
}

func (c *Client) sendMessages() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter message to send (or type 'exit' to quit): ")
	for {
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			break
		}

		msg := c.ClientID + ": " + text

		if err := websocket.Message.Send(c.WS, msg); err != nil {
			log.Fatal("Error sending message:", err)
		}
	}
}

func main() {
	fmt.Println("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	ws, err := websocket.Dial("ws://localhost:8082/ws", "", "https://test.net")
	if err != nil {
		log.Fatal("WebSocket dial error:", err)
	}
	defer ws.Close()

	client := NewClient(ws, name)

	go client.receiveMessages()
	client.sendMessages()

	select {}
}
