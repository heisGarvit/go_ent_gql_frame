package customfields

import (
	"database/sql/driver"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"time"
)

type DateOnly sql.NullTime

func NullableDateOnly(name string, nullable bool) ent.Field {
	f := field.Other(name, &DateOnly{}).SchemaType(
		map[string]string{
			dialect.MySQL:    "date",
			dialect.Postgres: "date",
		}).Annotations(
		entgql.OrderField(gen.Funcs["camel"].(func(string) string)(name)),
		entgql.Type("DateOnly"),
	)
	if nullable {
		return f.Optional().Nillable()
	}
	return f
}

func (j *DateOnly) Scan(value interface{}) error {
	if value == nil {
		j = nil
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal Date value:", value))
	}
	*j = DateOnly(sql.NullTime{Time: t, Valid: true})
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *DateOnly) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	z := sql.NullTime(*j)
	return []byte(z.Time.Format(time.DateOnly)), nil
}

func MarshalDateOnly(t DateOnly) graphql.Marshaler {

	if t.Valid {
		return graphql.MarshalString(t.Time.Format(time.DateOnly))
	} else {
		return graphql.Null
	}

}

func UnmarshalDateOnly(v interface{}) (DateOnly, error) {
	dateAsString, ok := v.(string)
	if !ok {
		return DateOnly(sql.NullTime{}), errors.New(" should be a valid date like 2006-01-02")
	}
	t, err := time.Parse(time.DateOnly, dateAsString)
	if err != nil {
		return DateOnly(sql.NullTime{}), err
	}
	return DateOnly(sql.NullTime{Time: t, Valid: true}), err
}
