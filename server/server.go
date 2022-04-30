package server

import (
	"log"
	"net/http"

	"instrumented_web_server/api"

	"github.com/gorilla/mux"
)

type server struct {
	port   string
	router *mux.Router
}

func New(port string) *server {
	s := &server{
		port:   port,
		router: mux.NewRouter(),
	}

	s.router.HandleFunc("/api/v1/test", api.ExampleHandler)
	return s
}

func (s *server) Run() {
	log.Fatal(http.ListenAndServe(s.port, s.router))
}
