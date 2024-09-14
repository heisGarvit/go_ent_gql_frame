package apm

import (
	"go.elastic.co/apm/module/apmotel/v2"
	"go.opentelemetry.io/otel"
	"log/slog"
	"project/utils/apm/apmslog"
	"project/utils/env"
)

type ResolverSpanNames struct {
	Names []string
}

var LogLevel slog.Level

func init() {
	InitLogger()
	InitApm()
}

func InitLogger() {

	err := LogLevel.UnmarshalText([]byte(env.Get("LOG_LEVEL")))
	if err != nil {
		LogLevel = slog.LevelInfo
	}

	slog.SetDefault(apmslog.InitLogger(LogLevel, true))

	slog.Info("Logger initialized", "level", LogLevel)

}

func InitApm() {
	provider, err := apmotel.NewTracerProvider()
	if err != nil {
		slog.Error("Apm NewTracerProvider", "error", err)
	} else {
		otel.SetTracerProvider(provider)
	}
}
