package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"os"
	"strings"
)

func receiveMessages(ws *websocket.Conn) {
	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Fatal("Error receiving message:", err)
		}
		fmt.Println("Received:", msg)
	}
}

func sendMessages(ws *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message to send (or type 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			break
		}
		if err := websocket.Message.Send(ws, text); err != nil {
			log.Fatal("Error sending message:", err)
		}
	}
}

func main() {
	ws, err := websocket.Dial("ws://localhost:6969/ws", "", "http://ujstor.dev")
	if err != nil {
		log.Fatal("WebSocket dial error:", err)
	}
	defer ws.Close()

	go receiveMessages(ws)
	sendMessages(ws)
}