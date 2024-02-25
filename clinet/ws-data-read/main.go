package main

import (
    "log"
    "golang.org/x/net/websocket"
)

func main() {
    ws, err := websocket.Dial("ws://localhost:8083/websocket", "", "http://test.com/")
    if err != nil {
        log.Fatal("WebSocket dial error:", err)
    }

    defer ws.Close()

    var response = make([]byte, 512)
    for {
        n, err := ws.Read(response)
        if err != nil {
            log.Fatal("Read error:", err)
        }
        log.Println("Received:", string(response[:n]))
    }
}