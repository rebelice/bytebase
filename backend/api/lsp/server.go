package lsp

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/bytebase/bytebase/backend/common/log"
	"github.com/bytebase/bytebase/backend/store"
)

const (
	readTimeout  = 75 * time.Second
	writeTimeout = 60 * time.Second
)

// Server is the Language Server Protocol service.
type Server struct {
	connectionCount atomic.Uint64
	server          *http.Server

	store *store.Store
}

// NewServer creates a Language Server Protocol service.
func NewServer(
	store *store.Store,
) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) Serve(lis net.Listener) error {
	return s.server.Serve(lis)
}

func (s *Server) GracefulStop(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown LSP server", log.BBError(err))
	}
}

func (s *Server) Stop() {
	if err := s.server.Close(); err != nil {
		slog.Error("Failed to close LSP server", log.BBError(err))
	}
}
