package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("company"),
		edge.From("roles", Role.Type).Ref("company"),
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).Unique().Field("company_id"),
		edge.From("permissions", RolePermission.Type).Ref("role"),
	}
}

func (RolePermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).Unique().Field("company_id"),
		edge.To("role", Role.Type).Unique().Field("role_id"),
		edge.To("permission", Permission.Type).Unique().Field("permission_id"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).Unique().Field("company_id"),
	}

}
