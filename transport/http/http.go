package http

import (
    "context"
)
func Listen(ctx context.Context, opts ...Option) error {
    cfg := config{}

    for _, opt := range opts {
        if err := opt(&cfg); err != nil {
            return err
        }
    }

    return nil
}
