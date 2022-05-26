package http

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

// Ensure logging is configured at least once.
func init() {
	if err := configureLogging("debug"); err != nil {
		panic(err)
	}
}

var logOnce sync.Once

// See: https://pkg.go.dev/github.com/rs/zerolog#pkg-variables
func configureLogging(lvl string) error {
	// Enable compatibility with StackDriver.
	logOnce.Do(func() {
		zerolog.LevelFieldName = "severity"
	})

	switch strings.ToLower(lvl) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		return fmt.Errorf("unknown loglevel: %s", lvl)
	}
	return nil
}

// Option is an application setting.
type Option func(*config) error

type config struct {
	Address string
	Timeout time.Duration
}

// WithLogLevel configures the log level.
func WithLogLevel(lvl string) Option {
	return func(cfg *config) error {
		return configureLogging(lvl)
	}
}

// WithEndpoint configures the search endpoint.
func WithEndpoint(addr, timeout string) Option {
	return func(cfg *config) error {
		dur, err := time.ParseDuration(timeout)
		if err != nil {
			return err
		}
		cfg.Address = addr
		cfg.Timeout = dur
		return nil
	}
}
