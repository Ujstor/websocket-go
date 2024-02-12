package main

import (
    "fmt"
    "log"
    "golang.org/x/net/websocket"
)

func main() {
    ws, err := websocket.Dial("ws://localhost:6969/orderbook", "", "http://localhost/")
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
        fmt.Println("Received:", string(response[:n]))
    }
}
