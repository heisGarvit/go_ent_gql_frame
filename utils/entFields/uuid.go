package entFields

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"io"
	"strconv"
)

func NullableUUID(name string, nullable bool) ent.Field {
	f := field.UUID(name, uuid.UUID{}).Annotations(entgql.Type("ID"))
	if nullable {
		return f.Optional().Nillable()
	}
	return f
}

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(u.String()))
	})
}

func UnmarshalUUID(v interface{}) (u uuid.UUID, err error) {
	s, ok := v.(string)
	if !ok {
		return u, fmt.Errorf("invalid type %T, expect string", v)
	}
	return uuid.Parse(s)
}
