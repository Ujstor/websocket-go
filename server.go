package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)


type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err!= nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
				break
			}
			fmt.Println("Error reading from client:", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("Thank you for the message!"))
	}
}

func (s *Server) handleWs(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func main () {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWs))
	http.ListenAndServe(":6969", nil)
}