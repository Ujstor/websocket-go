package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"golang.org/x/net/websocket"
	"time"
	"fmt"	
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Handle("/websocket", websocket.Handler(s.handleWsOrderbook))
	r.Handle("/ws", websocket.Handler(s.handleWs))

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) handleWsOrderbook(ws *websocket.Conn) {
	log.Println("New incoming connection from client:", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("orderbook data ----> %d", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}


func (s *Server) handleWs(ws *websocket.Conn) {
	log.Println("New incoming connection from client:", ws.RemoteAddr())

	s.conn[ws] = true

	s.readLoop(ws)
}