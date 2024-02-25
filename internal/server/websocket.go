package server 

import (
	"log"
	"io"
	"golang.org/x/net/websocket"
	
)


func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err!= nil {
			if err == io.EOF {
				log.Println("Client disconnected")
				break
			}
			log.Println("Error reading from client:", err)
			continue
		}
		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for  ws := range s.conn {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err!= nil {
				log.Println("Write error:", err)
			}
		} (ws)
	}
}