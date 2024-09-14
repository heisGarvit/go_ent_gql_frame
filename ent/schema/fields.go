package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"project/utils/entFields"
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

		entFields.NullableDateOnly("date_of_birth", true),
		entFields.NullableDateOnly("date_of_join", false),

		entFields.NullableTimeOnly("timing_start", true),

		field.String("employee_id").Optional(),
		field.UUID("role_id", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}
