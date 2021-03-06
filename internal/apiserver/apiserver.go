package apiserver

import (
	"net/http"

	"project/internal/config"
	"project/internal/logger"
	"project/internal/routes"

	"github.com/gorilla/mux"
)

var (
	Instance *APIServer
)

// APIServer struct
type APIServer struct {
	router *mux.Router
}

// New creates new api server
func New(config *config.Config) *APIServer {
	router := mux.NewRouter()

	routes.ApplyRoutes(router)

	//http.Handle("/", router)

	Instance = &APIServer{
		router: router,
	}

	return Instance
}

// Start api server
func (s *APIServer) Start() error {
	logger.Instance.LogInfo("Server starting...")

	http.ListenAndServe(config.Instance.BindAddr, s.router)

	return nil
}
