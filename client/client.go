package main

import (
    "fmt"
    "log"
    "golang.org/x/net/websocket"
)

func main() {
    // Connect to WebSocket server
    ws, err := websocket.Dial("ws://localhost:6969/ws", "", "http://localhost/")
    if err != nil {
        log.Fatal("WebSocket dial error:", err)
    }
    defer ws.Close()

    // Send a message
    message := "Hello, WebSocket Server!"
    _, err = ws.Write([]byte(message))
    if err != nil {
        log.Fatal("Write error:", err)
    }
    fmt.Println("Sent:", message)

    // Read the response from the server
    var response = make([]byte, 512)
    n, err := ws.Read(response)
    if err != nil {
        log.Fatal("Read error:", err)
    }
    fmt.Println("Received:", string(response[:n]))
}
