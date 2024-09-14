package entFields

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/99designs/gqlgen/graphql"
)

func NullableString(name string, nullable bool) ent.Field {
	if !nullable {
		return field.String(name).Annotations(entgql.OrderField(gen.Funcs["camel"].(func(string) string)(name)))
	}
	return field.Other(name, &sql.NullString{}).SchemaType(
		map[string]string{
			dialect.MySQL:    "text",
			dialect.Postgres: "text",
		}).Annotations(
		entgql.OrderField(gen.Funcs["camel"].(func(string) string)(name)),
		entgql.Type("NullString"),
	).Optional()
}

func MarshalNullString(ns sql.NullString) graphql.Marshaler {
	if !ns.Valid {
		// this is also important, so we can detect if this scalar is used in a not null context and return an appropriate error
		return graphql.Null
	}
	return graphql.MarshalString(ns.String)
}

func UnmarshalNullString(v interface{}) (sql.NullString, error) {

	if v == nil {
		return sql.NullString{Valid: false}, nil
	}
	// again you can delegate to the default implementation to save yourself some work.
	s, err := graphql.UnmarshalString(v)
	return sql.NullString{String: s, Valid: true}, err
}
