package http

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/meshenka/pin"
	"github.com/rs/zerolog/log"
)

func Listen(ctx context.Context, opts ...Option) error {
	cfg := config{}

	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			return err
		}
	}

	r := chi.NewRouter()
	r.Use(LoggerMiddleware(&log.Logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.NoCache)
	r.Use(middleware.Recoverer)

	generator := pin.NewGenerator()
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(cfg.Timeout))
	r.Post("/pin/_get", get(generator))
	r.Post("/pin/_verify", heartbeat())
	r.Get("/__internal__/heartbeat", heartbeat())
	return serve(ctx, cfg.Address, r)
}

// serve routes HTTP requests to handler.
func serve(ctx context.Context, addr string, handler http.Handler) error {
	log.Debug().Str("address", addr).Msg("starting HTTP server")

	srv := new(http.Server)
	srv.Addr = addr
	srv.Handler = handler

	sink := make(chan error, 1)

	go func() {
		defer close(sink)
		sink <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		return shutdown(srv)
	case err := <-sink:
		return err
	}
}

func shutdown(srv *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}
