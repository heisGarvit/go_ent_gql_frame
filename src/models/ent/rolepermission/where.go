// Code generated by ent, DO NOT EDIT.

package rolepermission

import (
	"project/src/models/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldUpdatedAt, v))
}

// CompanyID applies equality check predicate on the "company_id" field. It's identical to CompanyIDEQ.
func CompanyID(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldCompanyID, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldRoleID, v))
}

// PermissionID applies equality check predicate on the "permission_id" field. It's identical to PermissionIDEQ.
func PermissionID(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldPermissionID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldUpdatedAt, v))
}

// CompanyIDEQ applies the EQ predicate on the "company_id" field.
func CompanyIDEQ(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldCompanyID, v))
}

// CompanyIDNEQ applies the NEQ predicate on the "company_id" field.
func CompanyIDNEQ(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldCompanyID, v))
}

// CompanyIDIn applies the In predicate on the "company_id" field.
func CompanyIDIn(vs ...uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldCompanyID, vs...))
}

// CompanyIDNotIn applies the NotIn predicate on the "company_id" field.
func CompanyIDNotIn(vs ...uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldCompanyID, vs...))
}

// CompanyIDIsNil applies the IsNil predicate on the "company_id" field.
func CompanyIDIsNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIsNull(FieldCompanyID))
}

// CompanyIDNotNil applies the NotNil predicate on the "company_id" field.
func CompanyIDNotNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotNull(FieldCompanyID))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDIsNil applies the IsNil predicate on the "role_id" field.
func RoleIDIsNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIsNull(FieldRoleID))
}

// RoleIDNotNil applies the NotNil predicate on the "role_id" field.
func RoleIDNotNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotNull(FieldRoleID))
}

// PermissionIDEQ applies the EQ predicate on the "permission_id" field.
func PermissionIDEQ(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldPermissionID, v))
}

// PermissionIDNEQ applies the NEQ predicate on the "permission_id" field.
func PermissionIDNEQ(v uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldPermissionID, v))
}

// PermissionIDIn applies the In predicate on the "permission_id" field.
func PermissionIDIn(vs ...uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldPermissionID, vs...))
}

// PermissionIDNotIn applies the NotIn predicate on the "permission_id" field.
func PermissionIDNotIn(vs ...uuid.UUID) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldPermissionID, vs...))
}

// PermissionIDIsNil applies the IsNil predicate on the "permission_id" field.
func PermissionIDIsNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIsNull(FieldPermissionID))
}

// PermissionIDNotNil applies the NotNil predicate on the "permission_id" field.
func PermissionIDNotNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotNull(FieldPermissionID))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermission applies the HasEdge predicate on the "permission" edge.
func HasPermission() predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PermissionTable, PermissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionWith applies the HasEdge predicate on the "permission" edge with a given conditions (other predicates).
func HasPermissionWith(preds ...predicate.Permission) predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := newPermissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.NotPredicates(p))
}
