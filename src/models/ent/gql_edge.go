// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Company) Users(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.UsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryUsers().All(ctx)
	}
	return result, err
}

func (c *Company) Roles(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryRoles().All(ctx)
	}
	return result, err
}

func (r *Role) Company(ctx context.Context) (*Company, error) {
	result, err := r.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Role) Permissions(ctx context.Context) (result []*RolePermission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedPermissions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.PermissionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryPermissions().All(ctx)
	}
	return result, err
}

func (rp *RolePermission) Company(ctx context.Context) (*Company, error) {
	result, err := rp.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = rp.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (rp *RolePermission) Role(ctx context.Context) (*Role, error) {
	result, err := rp.Edges.RoleOrErr()
	if IsNotLoaded(err) {
		result, err = rp.QueryRole().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (rp *RolePermission) Permission(ctx context.Context) (*Permission, error) {
	result, err := rp.Edges.PermissionOrErr()
	if IsNotLoaded(err) {
		result, err = rp.QueryPermission().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Company(ctx context.Context) (*Company, error) {
	result, err := u.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}
