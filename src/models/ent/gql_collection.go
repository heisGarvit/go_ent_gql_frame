// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"project/src/models/ent/company"
	"project/src/models/ent/permission"
	"project/src/models/ent/role"
	"project/src/models/ent/rolepermission"
	"project/src/models/ent/user"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CompanyQuery) CollectFields(ctx context.Context, satisfies ...string) (*CompanyQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CompanyQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(company.Columns))
		selectedFields = []string{company.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoleClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedRoles(alias, func(wq *RoleQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[company.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, company.FieldCreatedAt)
				fieldSeen[company.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[company.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, company.FieldUpdatedAt)
				fieldSeen[company.FieldUpdatedAt] = struct{}{}
			}
		case "companyID":
			if _, ok := fieldSeen[company.FieldCompanyID]; !ok {
				selectedFields = append(selectedFields, company.FieldCompanyID)
				fieldSeen[company.FieldCompanyID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[company.FieldName]; !ok {
				selectedFields = append(selectedFields, company.FieldName)
				fieldSeen[company.FieldName] = struct{}{}
			}
		case "isDisabled":
			if _, ok := fieldSeen[company.FieldIsDisabled]; !ok {
				selectedFields = append(selectedFields, company.FieldIsDisabled)
				fieldSeen[company.FieldIsDisabled] = struct{}{}
			}
		case "ownerID":
			if _, ok := fieldSeen[company.FieldOwnerID]; !ok {
				selectedFields = append(selectedFields, company.FieldOwnerID)
				fieldSeen[company.FieldOwnerID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type companyPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CompanyPaginateOption
}

func newCompanyPaginateArgs(rv map[string]any) *companyPaginateArgs {
	args := &companyPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*CompanyOrder:
			args.opts = append(args.opts, WithCompanyOrder(v))
		case []any:
			var orders []*CompanyOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &CompanyOrder{Field: &CompanyOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithCompanyOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*CompanyWhereInput); ok {
		args.opts = append(args.opts, WithCompanyFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pe *PermissionQuery) CollectFields(ctx context.Context, satisfies ...string) (*PermissionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pe, nil
	}
	if err := pe.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pe *PermissionQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(permission.Columns))
		selectedFields = []string{permission.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdAt":
			if _, ok := fieldSeen[permission.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, permission.FieldCreatedAt)
				fieldSeen[permission.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[permission.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, permission.FieldUpdatedAt)
				fieldSeen[permission.FieldUpdatedAt] = struct{}{}
			}
		case "companyID":
			if _, ok := fieldSeen[permission.FieldCompanyID]; !ok {
				selectedFields = append(selectedFields, permission.FieldCompanyID)
				fieldSeen[permission.FieldCompanyID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[permission.FieldName]; !ok {
				selectedFields = append(selectedFields, permission.FieldName)
				fieldSeen[permission.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[permission.FieldDescription]; !ok {
				selectedFields = append(selectedFields, permission.FieldDescription)
				fieldSeen[permission.FieldDescription] = struct{}{}
			}
		case "read":
			if _, ok := fieldSeen[permission.FieldRead]; !ok {
				selectedFields = append(selectedFields, permission.FieldRead)
				fieldSeen[permission.FieldRead] = struct{}{}
			}
		case "write":
			if _, ok := fieldSeen[permission.FieldWrite]; !ok {
				selectedFields = append(selectedFields, permission.FieldWrite)
				fieldSeen[permission.FieldWrite] = struct{}{}
			}
		case "patch":
			if _, ok := fieldSeen[permission.FieldPatch]; !ok {
				selectedFields = append(selectedFields, permission.FieldPatch)
				fieldSeen[permission.FieldPatch] = struct{}{}
			}
		case "delete":
			if _, ok := fieldSeen[permission.FieldDelete]; !ok {
				selectedFields = append(selectedFields, permission.FieldDelete)
				fieldSeen[permission.FieldDelete] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pe.Select(selectedFields...)
	}
	return nil
}

type permissionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PermissionPaginateOption
}

func newPermissionPaginateArgs(rv map[string]any) *permissionPaginateArgs {
	args := &permissionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*PermissionOrder:
			args.opts = append(args.opts, WithPermissionOrder(v))
		case []any:
			var orders []*PermissionOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &PermissionOrder{Field: &PermissionOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithPermissionOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*PermissionWhereInput); ok {
		args.opts = append(args.opts, WithPermissionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RoleQuery) CollectFields(ctx context.Context, satisfies ...string) (*RoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RoleQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(role.Columns))
		selectedFields = []string{role.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "company":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CompanyClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			r.withCompany = query
			if _, ok := fieldSeen[role.FieldCompanyID]; !ok {
				selectedFields = append(selectedFields, role.FieldCompanyID)
				fieldSeen[role.FieldCompanyID] = struct{}{}
			}
		case "permissions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RolePermissionClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			r.WithNamedPermissions(alias, func(wq *RolePermissionQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[role.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldCreatedAt)
				fieldSeen[role.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[role.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldUpdatedAt)
				fieldSeen[role.FieldUpdatedAt] = struct{}{}
			}
		case "companyID":
			if _, ok := fieldSeen[role.FieldCompanyID]; !ok {
				selectedFields = append(selectedFields, role.FieldCompanyID)
				fieldSeen[role.FieldCompanyID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[role.FieldName]; !ok {
				selectedFields = append(selectedFields, role.FieldName)
				fieldSeen[role.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[role.FieldDescription]; !ok {
				selectedFields = append(selectedFields, role.FieldDescription)
				fieldSeen[role.FieldDescription] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type rolePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RolePaginateOption
}

func newRolePaginateArgs(rv map[string]any) *rolePaginateArgs {
	args := &rolePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*RoleOrder:
			args.opts = append(args.opts, WithRoleOrder(v))
		case []any:
			var orders []*RoleOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &RoleOrder{Field: &RoleOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithRoleOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*RoleWhereInput); ok {
		args.opts = append(args.opts, WithRoleFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (rp *RolePermissionQuery) CollectFields(ctx context.Context, satisfies ...string) (*RolePermissionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return rp, nil
	}
	if err := rp.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return rp, nil
}

func (rp *RolePermissionQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(rolepermission.Columns))
		selectedFields = []string{rolepermission.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "company":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CompanyClient{config: rp.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			rp.withCompany = query
			if _, ok := fieldSeen[rolepermission.FieldCompanyID]; !ok {
				selectedFields = append(selectedFields, rolepermission.FieldCompanyID)
				fieldSeen[rolepermission.FieldCompanyID] = struct{}{}
			}
		case "role":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoleClient{config: rp.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			rp.withRole = query
			if _, ok := fieldSeen[rolepermission.FieldRoleID]; !ok {
				selectedFields = append(selectedFields, rolepermission.FieldRoleID)
				fieldSeen[rolepermission.FieldRoleID] = struct{}{}
			}
		case "permission":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PermissionClient{config: rp.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			rp.withPermission = query
			if _, ok := fieldSeen[rolepermission.FieldPermissionID]; !ok {
				selectedFields = append(selectedFields, rolepermission.FieldPermissionID)
				fieldSeen[rolepermission.FieldPermissionID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[rolepermission.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, rolepermission.FieldCreatedAt)
				fieldSeen[rolepermission.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[rolepermission.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, rolepermission.FieldUpdatedAt)
				fieldSeen[rolepermission.FieldUpdatedAt] = struct{}{}
			}
		case "companyID":
			if _, ok := fieldSeen[rolepermission.FieldCompanyID]; !ok {
				selectedFields = append(selectedFields, rolepermission.FieldCompanyID)
				fieldSeen[rolepermission.FieldCompanyID] = struct{}{}
			}
		case "roleID":
			if _, ok := fieldSeen[rolepermission.FieldRoleID]; !ok {
				selectedFields = append(selectedFields, rolepermission.FieldRoleID)
				fieldSeen[rolepermission.FieldRoleID] = struct{}{}
			}
		case "permissionID":
			if _, ok := fieldSeen[rolepermission.FieldPermissionID]; !ok {
				selectedFields = append(selectedFields, rolepermission.FieldPermissionID)
				fieldSeen[rolepermission.FieldPermissionID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		rp.Select(selectedFields...)
	}
	return nil
}

type rolepermissionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RolePermissionPaginateOption
}

func newRolePermissionPaginateArgs(rv map[string]any) *rolepermissionPaginateArgs {
	args := &rolepermissionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*RolePermissionOrder:
			args.opts = append(args.opts, WithRolePermissionOrder(v))
		case []any:
			var orders []*RolePermissionOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &RolePermissionOrder{Field: &RolePermissionOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithRolePermissionOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*RolePermissionWhereInput); ok {
		args.opts = append(args.opts, WithRolePermissionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "company":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CompanyClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.withCompany = query
			if _, ok := fieldSeen[user.FieldCompanyID]; !ok {
				selectedFields = append(selectedFields, user.FieldCompanyID)
				fieldSeen[user.FieldCompanyID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[user.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldUpdatedAt)
				fieldSeen[user.FieldUpdatedAt] = struct{}{}
			}
		case "companyID":
			if _, ok := fieldSeen[user.FieldCompanyID]; !ok {
				selectedFields = append(selectedFields, user.FieldCompanyID)
				fieldSeen[user.FieldCompanyID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "mobile":
			if _, ok := fieldSeen[user.FieldMobile]; !ok {
				selectedFields = append(selectedFields, user.FieldMobile)
				fieldSeen[user.FieldMobile] = struct{}{}
			}
		case "isDisabled":
			if _, ok := fieldSeen[user.FieldIsDisabled]; !ok {
				selectedFields = append(selectedFields, user.FieldIsDisabled)
				fieldSeen[user.FieldIsDisabled] = struct{}{}
			}
		case "roleID":
			if _, ok := fieldSeen[user.FieldRoleID]; !ok {
				selectedFields = append(selectedFields, user.FieldRoleID)
				fieldSeen[user.FieldRoleID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*UserOrder:
			args.opts = append(args.opts, WithUserOrder(v))
		case []any:
			var orders []*UserOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &UserOrder{Field: &UserOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithUserOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
