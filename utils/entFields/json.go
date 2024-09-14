package entFields

import (
	"database/sql/driver"
	"encoding/json"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"errors"
	"fmt"
	"io"
)

type JSON json.RawMessage

func NullableJSON(name string, nullable bool) ent.Field {
	f := field.Other(name, &JSON{}).SchemaType(
		map[string]string{
			dialect.MySQL:    "JSON",
			dialect.Postgres: "JSON",
		}).Annotations(
		entgql.Type("JSON"),
	)
	if nullable {
		return f.Nillable().Optional()
	}
	return f
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (j *JSON) UnmarshalGQL(v interface{}) error {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("must be valid json")
	}
	jsonRawMessage := json.RawMessage(jsonBytes)

	*j = JSON(jsonRawMessage)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (j JSON) MarshalGQL(w io.Writer) {
	marshalled, marshalError := json.RawMessage(j).MarshalJSON()
	if marshalError == nil {
		w.Write(marshalled)
	}
}
