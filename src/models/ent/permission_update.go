// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project/src/models/ent/permission"
	"project/src/models/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks     []Hook
	mutation  *PermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetCompanyID sets the "company_id" field.
func (pu *PermissionUpdate) SetCompanyID(u uuid.UUID) *PermissionUpdate {
	pu.mutation.SetCompanyID(u)
	return pu
}

// SetNillableCompanyID sets the "company_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCompanyID(u *uuid.UUID) *PermissionUpdate {
	if u != nil {
		pu.SetCompanyID(*u)
	}
	return pu
}

// ClearCompanyID clears the value of the "company_id" field.
func (pu *PermissionUpdate) ClearCompanyID() *PermissionUpdate {
	pu.mutation.ClearCompanyID()
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableName(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *PermissionUpdate) SetDescription(s string) *PermissionUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDescription(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetRead sets the "read" field.
func (pu *PermissionUpdate) SetRead(b bool) *PermissionUpdate {
	pu.mutation.SetRead(b)
	return pu
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableRead(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetRead(*b)
	}
	return pu
}

// SetWrite sets the "write" field.
func (pu *PermissionUpdate) SetWrite(b bool) *PermissionUpdate {
	pu.mutation.SetWrite(b)
	return pu
}

// SetNillableWrite sets the "write" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableWrite(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetWrite(*b)
	}
	return pu
}

// SetPatch sets the "patch" field.
func (pu *PermissionUpdate) SetPatch(b bool) *PermissionUpdate {
	pu.mutation.SetPatch(b)
	return pu
}

// SetNillablePatch sets the "patch" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillablePatch(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetPatch(*b)
	}
	return pu
}

// SetDelete sets the "delete" field.
func (pu *PermissionUpdate) SetDelete(b bool) *PermissionUpdate {
	pu.mutation.SetDelete(b)
	return pu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDelete(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetDelete(*b)
	}
	return pu
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermissionUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PermissionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.CompanyID(); ok {
		_spec.SetField(permission.FieldCompanyID, field.TypeUUID, value)
	}
	if pu.mutation.CompanyIDCleared() {
		_spec.ClearField(permission.FieldCompanyID, field.TypeUUID)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Read(); ok {
		_spec.SetField(permission.FieldRead, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Write(); ok {
		_spec.SetField(permission.FieldWrite, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Patch(); ok {
		_spec.SetField(permission.FieldPatch, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Delete(); ok {
		_spec.SetField(permission.FieldDelete, field.TypeBool, value)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetCompanyID sets the "company_id" field.
func (puo *PermissionUpdateOne) SetCompanyID(u uuid.UUID) *PermissionUpdateOne {
	puo.mutation.SetCompanyID(u)
	return puo
}

// SetNillableCompanyID sets the "company_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCompanyID(u *uuid.UUID) *PermissionUpdateOne {
	if u != nil {
		puo.SetCompanyID(*u)
	}
	return puo
}

// ClearCompanyID clears the value of the "company_id" field.
func (puo *PermissionUpdateOne) ClearCompanyID() *PermissionUpdateOne {
	puo.mutation.ClearCompanyID()
	return puo
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableName(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *PermissionUpdateOne) SetDescription(s string) *PermissionUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDescription(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetRead sets the "read" field.
func (puo *PermissionUpdateOne) SetRead(b bool) *PermissionUpdateOne {
	puo.mutation.SetRead(b)
	return puo
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableRead(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetRead(*b)
	}
	return puo
}

// SetWrite sets the "write" field.
func (puo *PermissionUpdateOne) SetWrite(b bool) *PermissionUpdateOne {
	puo.mutation.SetWrite(b)
	return puo
}

// SetNillableWrite sets the "write" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableWrite(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetWrite(*b)
	}
	return puo
}

// SetPatch sets the "patch" field.
func (puo *PermissionUpdateOne) SetPatch(b bool) *PermissionUpdateOne {
	puo.mutation.SetPatch(b)
	return puo
}

// SetNillablePatch sets the "patch" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillablePatch(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetPatch(*b)
	}
	return puo
}

// SetDelete sets the "delete" field.
func (puo *PermissionUpdateOne) SetDelete(b bool) *PermissionUpdateOne {
	puo.mutation.SetDelete(b)
	return puo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDelete(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetDelete(*b)
	}
	return puo
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermissionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PermissionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.CompanyID(); ok {
		_spec.SetField(permission.FieldCompanyID, field.TypeUUID, value)
	}
	if puo.mutation.CompanyIDCleared() {
		_spec.ClearField(permission.FieldCompanyID, field.TypeUUID)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Read(); ok {
		_spec.SetField(permission.FieldRead, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Write(); ok {
		_spec.SetField(permission.FieldWrite, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Patch(); ok {
		_spec.SetField(permission.FieldPatch, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Delete(); ok {
		_spec.SetField(permission.FieldDelete, field.TypeBool, value)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
