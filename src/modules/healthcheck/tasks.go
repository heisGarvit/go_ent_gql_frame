package healthcheck

import (
	"context"
	"log/slog"
	"project/src/infra/taskqueue"
)

func init() {
	err := taskqueue.Server.RegisterTask("Add", Add)
	if err != nil {
		slog.Error("Error registering task", "error", err)
	}
}

func Add(ctx context.Context, args ...int64) (int64, error) {
	var sum int64
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}
