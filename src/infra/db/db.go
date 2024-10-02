package db

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
	"project/src/infra/apm"
	"project/src/infra/env"
	"project/src/models/ent"
)

var Client *ent.Client

func init() {
	var err error
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
