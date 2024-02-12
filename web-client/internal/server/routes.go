package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./cmd/web/static"))))

	mux.HandleFunc("/", s.Root)

	return mux
}

func (s *Server) Root(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./cmd/web/templates/index.html")
}
