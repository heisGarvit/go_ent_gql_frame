package tracer

import "go.opentelemetry.io/otel"

var Tracer = otel.Tracer("gorm.io/plugin/opentelemetry")
