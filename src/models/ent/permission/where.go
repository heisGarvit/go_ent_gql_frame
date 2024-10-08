// Code generated by ent, DO NOT EDIT.

package permission

import (
	"project/src/models/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedAt, v))
}

// CompanyID applies equality check predicate on the "company_id" field. It's identical to CompanyIDEQ.
func CompanyID(v uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCompanyID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldDescription, v))
}

// Read applies equality check predicate on the "read" field. It's identical to ReadEQ.
func Read(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldRead, v))
}

// Write applies equality check predicate on the "write" field. It's identical to WriteEQ.
func Write(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldWrite, v))
}

// Patch applies equality check predicate on the "patch" field. It's identical to PatchEQ.
func Patch(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldPatch, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldDelete, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldUpdatedAt, v))
}

// CompanyIDEQ applies the EQ predicate on the "company_id" field.
func CompanyIDEQ(v uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCompanyID, v))
}

// CompanyIDNEQ applies the NEQ predicate on the "company_id" field.
func CompanyIDNEQ(v uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCompanyID, v))
}

// CompanyIDIn applies the In predicate on the "company_id" field.
func CompanyIDIn(vs ...uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldCompanyID, vs...))
}

// CompanyIDNotIn applies the NotIn predicate on the "company_id" field.
func CompanyIDNotIn(vs ...uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldCompanyID, vs...))
}

// CompanyIDGT applies the GT predicate on the "company_id" field.
func CompanyIDGT(v uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldCompanyID, v))
}

// CompanyIDGTE applies the GTE predicate on the "company_id" field.
func CompanyIDGTE(v uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldCompanyID, v))
}

// CompanyIDLT applies the LT predicate on the "company_id" field.
func CompanyIDLT(v uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldCompanyID, v))
}

// CompanyIDLTE applies the LTE predicate on the "company_id" field.
func CompanyIDLTE(v uuid.UUID) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldCompanyID, v))
}

// CompanyIDIsNil applies the IsNil predicate on the "company_id" field.
func CompanyIDIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldCompanyID))
}

// CompanyIDNotNil applies the NotNil predicate on the "company_id" field.
func CompanyIDNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldCompanyID))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldDescription, v))
}

// ReadEQ applies the EQ predicate on the "read" field.
func ReadEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldRead, v))
}

// ReadNEQ applies the NEQ predicate on the "read" field.
func ReadNEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldRead, v))
}

// WriteEQ applies the EQ predicate on the "write" field.
func WriteEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldWrite, v))
}

// WriteNEQ applies the NEQ predicate on the "write" field.
func WriteNEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldWrite, v))
}

// PatchEQ applies the EQ predicate on the "patch" field.
func PatchEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldPatch, v))
}

// PatchNEQ applies the NEQ predicate on the "patch" field.
func PatchNEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldPatch, v))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldDelete, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.NotPredicates(p))
}
