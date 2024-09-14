package bgTasks

import (
	"context"
	"log/slog"
	"project/utils/taskQueue"
)

var ImportWorkAround = ""

func init() {
	err := taskQueue.Server.RegisterTask("Add", Add)
	if err != nil {
		slog.Error("Error registering task", "error", err)
	}
}

func Add(x context.Context, args ...int64) (int64, int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, sum, nil
}
