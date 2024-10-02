package schema

import (
	"project/src/utils/schema"
)

type Company struct{ schema.Base }
type Permission struct{ schema.Base }

type Role struct{ schema.Base }
type RolePermission struct{ schema.Base }

type User struct{ schema.Base }
