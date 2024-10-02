package healthcheck

import (
	"context"
	"log/slog"
)

type Resolver struct{}

type Ping struct {
	Success bool
}

func (r *Resolver) Ping(ctx context.Context) (*Ping, error) {
	slog.InfoContext(ctx, "Ping")
	return &Ping{Success: true}, nil
}
