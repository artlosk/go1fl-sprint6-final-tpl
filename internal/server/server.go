package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	server := &Server{
		logger:     logger,
		httpServer: httpServer,
	}

	return server
}

func (s *Server) Start() error {
	s.logger.Printf("Сервер запущен на порту %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) GetHTTPServer() *http.Server {
	return s.httpServer
}
