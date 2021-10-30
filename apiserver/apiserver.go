package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Config *Config
	Router *mux.Router
}

// New - creates a new APIServer, using config params from Config and mux.
func New(config *Config) *APIServer {
	return &APIServer{
		Config: config,
		Router: mux.NewRouter(),
	}
}

// Start - starts our server up.
func (s *APIServer) Start() error {
	s.configureRouter()
	return http.ListenAndServe(s.Config.BindAddr, s.Router)
}

// configureRouter - local func to set up our routing.
func (s *APIServer) configureRouter() {
	s.Router.HandleFunc("/", s.HandleHello())
}

// HandleHello - sends back an http.HandlerFunc to send hello .
func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}
}
