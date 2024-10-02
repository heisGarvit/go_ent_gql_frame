package taskqueue

import (
	"errors"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/RichardKnop/machinery/v2/tasks"
	"log/slog"
	"os"
	"project/src/infra/env"
)

var Server *machinery.Server

func init() {

	cnf := &config.Config{
		DefaultQueue:    "bg_tasks",
		ResultsExpireIn: 3600,
		Broker:          env.Get("BROKER"),
		ResultBackend:   env.Get("RESULT_BACKEND"),
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	broker := redisbroker.NewGR(cnf, []string{env.Get("BROKER")}, 1)
	backend := redisbackend.NewGR(cnf, []string{env.Get("RESULT_BACKEND")}, 2)
	lock := eagerlock.New()
	Server = machinery.NewServer(cnf, broker, backend, lock)

	//go launchWorker()
}

func LaunchWorker() {

	worker := Server.NewWorker("bg_tasks_worker", 10)

	worker.SetErrorHandler(func(err error) {
		slog.Error("BgTasks worker error", "error", err)
	})
	worker.SetPreTaskHandler(func(signature *tasks.Signature) {
		slog.Debug("BgTasks worker task started", "task", signature.Name)
	})
	worker.SetPostTaskHandler(func(signature *tasks.Signature) {
		slog.Debug("BgTasks worker task finished", "task", signature.Name)
	})

	err := worker.Launch()
	if err != nil {
		if errors.Is(err, machinery.ErrWorkerQuitGracefully) {
			slog.Debug("BgTasks worker gracefully stopped")
			os.Exit(1)
		} else {
			slog.Error("worker.Launch()", "error", err)
		}
	}
}
