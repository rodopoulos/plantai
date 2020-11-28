package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server is one of the access layers to plantai. It exposes the following
// endpoints and methods:
//
// 	GET  /healthcheck
//
type Server struct {
	addr       string
	httpServer *http.Server
	logger     logrus.FieldLogger
}

func NewServer(addr string, logger logrus.FieldLogger) *Server {
	router := buildRouter()
	httpServer := &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	return &Server{
		addr:       addr,
		logger:     logger,
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen HTTP: %w", err)
	}
	defer listener.Close()

	s.logger.Infof("Server starting at %v", s.addr)
	err = s.httpServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to listen HTTP: %w", err)
	}

	return nil
}

func buildRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", healthcheck).Methods("GET")

	return router
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("WORKING"))
}
