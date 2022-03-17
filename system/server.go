package system

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Server struct{}

func ServerFromConfig(cfg *Config) *Server {
	return &Server{}
}

// Stop stops the server. Blocks until the server is stopped
// or the context times out.
func (s *Server) Stop(ctx context.Context) error {
	return nil
}

// Run the server. This blocks until the server is stopped.
func (s *Server) Run(ctx context.Context) error {
	server := http.Server{}
	err := server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Ctx(ctx).Info().Msg("server stopped serving with ErrServerClosed error")
			return nil
		}
		log.Ctx(ctx).Err(err).Msg("server stopped serving with error")
		return err
	}
	log.Ctx(ctx).Info().Msg("server stopped serving without error")
	return nil
}
