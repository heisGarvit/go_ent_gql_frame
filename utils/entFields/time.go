package entFields

import (
	"database/sql"
	"database/sql/driver"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"time"
)

type TimeOnly sql.NullTime

func NullableTimeOnly(name string, nullable bool) ent.Field {
	f := field.Other(name, &TimeOnly{}).SchemaType(
		map[string]string{
			dialect.MySQL:    "time",
			dialect.Postgres: "time",
		}).Annotations(
		entgql.OrderField(gen.Funcs["camel"].(func(string) string)(name)),
		entgql.Type("TimeOnly"),
	)
	if nullable {
		return f.Optional()
	}
	return f
}

func (j *TimeOnly) Scan(value interface{}) error {
	if value == nil {
		j = nil
		return nil
	}
	d, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal Date value:", value))
	}
	t, err := time.Parse(time.TimeOnly, string(d))
	if err != nil {
		return errors.New(fmt.Sprint("Failed to unmarshal Date value:", value))
	}
	*j = TimeOnly(sql.NullTime{Time: t, Valid: true})
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *TimeOnly) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	z := sql.NullTime(*j)
	return []byte(z.Time.Format(time.TimeOnly)), nil
}

func MarshalTimeOnly(t TimeOnly) graphql.Marshaler {

	if t.Valid {
		return graphql.MarshalString(t.Time.Format(time.TimeOnly))
	} else {
		return graphql.Null
	}

}

func UnmarshalTimeOnly(v interface{}) (TimeOnly, error) {
	dateAsString, ok := v.(string)
	if !ok {
		return TimeOnly(sql.NullTime{}), errors.New(" should be a valid date like 15:04:05")
	}
	t, err := time.Parse(time.TimeOnly, dateAsString)
	if err != nil {
		return TimeOnly(sql.NullTime{}), err
	}
	return TimeOnly(sql.NullTime{Time: t, Valid: true}), err
}
