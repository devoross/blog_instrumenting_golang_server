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
	store  *api.Store
}

func New(port string) *server {
	s := &server{
		port:   port,
		router: mux.NewRouter(),
		store:  api.NewStore(),
	}

	// populate the store in a go routine
	go s.store.PopulateStore(1)

	// every endpoint registered again this router will execute this middleware
	s.router.Use(api.ExampleMiddleware)
	s.router.HandleFunc("/api/v1/advice", s.store.AdviceHandler)
	return s
}

func (s *server) Run() {
	log.Fatal(http.ListenAndServe(s.port, s.router))
}
