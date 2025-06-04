package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/2pizzzza/IskenderBackend/pkg/logger"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

type Server struct {
	App *http.Server
	Log *zap.Logger
	Mux *http.ServeMux
}

func New(log *zap.Logger, host, port string) *Server {
	mux := http.NewServeMux()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowedHeaders:   []string{},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(mux)
	loggerMux := logger.LoggingMiddleware(log)(handler)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: loggerMux,
	}

	return &Server{
		App: httpServer,
		Log: log,
		Mux: mux,
	}
}

func (s *Server) Run() error {
	if err := s.App.ListenAndServe(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.App.Shutdown(ctx); err != nil {
		s.Log.Error(fmt.Sprintf("error while closing HTTP server: %v", err))
	}
}
