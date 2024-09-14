package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"project/utils/resource"
	"time"
)

type Base struct {
	ent.Schema
}

func (Base) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
	}
}

type Mixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Annotations(
				entgql.OrderField("id"),
				entsql.DefaultExpr("uuid_generate_v4()"),
			).Unique(),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(
				entgql.OrderField("createdAt"),
				entsql.Default("now()"),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("updatedAt"),
				entsql.Default("now()"),
			),
		field.UUID("company_id", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}

func (Mixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("company_id"),
	}
}

func (b Base) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.Mutations(),
		entgql.MultiOrder(),
	}
}

func (Base) Resource() *resource.Base[any, any, any, any, any, any] {
	return &resource.Base[any, any, any, any, any, any]{}
}
