package main

import (
	"context"
	"github.com/meshenka/pin/cmd"
	"github.com/meshenka/pin/transport/http"
	"os"
)

func main() {
	if err := cmd.Run(run); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	return http.Listen(
		ctx,
		http.WithEndpoint(
			cmd.Env("APPLICATION_ADDR", ":8080"),
			cmd.Env("HTTP_TIMEOUT", "5s"),
		),
		http.WithLogLevel(
			cmd.Env("LOGLEVEL", "debug"),
		),
	)
}
