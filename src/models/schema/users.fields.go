package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("is_disabled").Default(false),
		field.UUID("owner_id", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}

func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Bool("read").Default(false),
		field.Bool("write").Default(false),
		field.Bool("patch").Default(false),
		field.Bool("delete").Default(false),
	}
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
	}
}

func (RolePermission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("role_id", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
		field.UUID("permission_id", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").Optional().Nillable(),

		field.String("email"),
		field.String("mobile"),

		field.String("password").Sensitive(),
		field.Bool("is_disabled").Default(false),

		field.UUID("role_id", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}
