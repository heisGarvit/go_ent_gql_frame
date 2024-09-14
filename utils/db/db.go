package db

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
	"log/slog"
	"project/ent"
	"project/utils/apm"
	"project/utils/env"
)

var Client *ent.Client

func init() {
	var err error
	slog.Debug(env.Get("DB_URI"))
	db, err := otelsql.Open(
		"postgres",
		env.Get("DB_URI"),
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
	)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	_, err = db.Query("SELECT 1")
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
		return
	}

	Client = ent.NewClient(ent.Driver(sql.OpenDB(dialect.Postgres, db)))

	if apm.LogLevel.String() == "DEBUG" {
		Client = Client.Debug()
	}

}
