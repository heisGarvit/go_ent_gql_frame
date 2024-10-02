package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/index"
)

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("company_id", "email").Unique(),
		index.Fields("company_id", "mobile").Unique(),
	}
}
