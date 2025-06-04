package httpserver

import (
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	App *http.Server
	Log *zap.Logger
	Mux *http.ServeMux
}

func New()
