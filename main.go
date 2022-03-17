package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
)

func contextWithLogger() context.Context {
	subLogger := log.With().Logger()
	return subLogger.WithContext(context.Background())
}

func main() {
	ctx := contextWithLogger()

	configuration := system.ConfigFromEnv()
	server := system.ServerFromConfig(configuration)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	go func() {
		s := <-signals
		log.Ctx(ctx).Info().Str("signal", s.String()).Msg("signal received. Stop server")
		err := server.Stop(ctx)
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("stopping server failed")
		}
	}()

	if err := server.Run(ctx); err != nil {
		log.Ctx(ctx).Err(err).Msg("server stopped")
	} else {
		log.Ctx(ctx).Info().Msg("server stopped")
	}
}
