// Code generated by ent, DO NOT EDIT.

package user

import (
	"project/ent/predicate"
	"project/utils/entFields"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// CompanyID applies equality check predicate on the "company_id" field. It's identical to CompanyIDEQ.
func CompanyID(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCompanyID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// IsDisabled applies equality check predicate on the "is_disabled" field. It's identical to IsDisabledEQ.
func IsDisabled(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsDisabled, v))
}

// DateOfBirth applies equality check predicate on the "date_of_birth" field. It's identical to DateOfBirthEQ.
func DateOfBirth(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDateOfBirth, v))
}

// DateOfJoin applies equality check predicate on the "date_of_join" field. It's identical to DateOfJoinEQ.
func DateOfJoin(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDateOfJoin, v))
}

// TimingStart applies equality check predicate on the "timing_start" field. It's identical to TimingStartEQ.
func TimingStart(v *entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldEQ(FieldTimingStart, v))
}

// EmployeeID applies equality check predicate on the "employee_id" field. It's identical to EmployeeIDEQ.
func EmployeeID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmployeeID, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRoleID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// CompanyIDEQ applies the EQ predicate on the "company_id" field.
func CompanyIDEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCompanyID, v))
}

// CompanyIDNEQ applies the NEQ predicate on the "company_id" field.
func CompanyIDNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCompanyID, v))
}

// CompanyIDIn applies the In predicate on the "company_id" field.
func CompanyIDIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldCompanyID, vs...))
}

// CompanyIDNotIn applies the NotIn predicate on the "company_id" field.
func CompanyIDNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCompanyID, vs...))
}

// CompanyIDIsNil applies the IsNil predicate on the "company_id" field.
func CompanyIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCompanyID))
}

// CompanyIDNotNil applies the NotNil predicate on the "company_id" field.
func CompanyIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCompanyID))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldMobile, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// IsDisabledEQ applies the EQ predicate on the "is_disabled" field.
func IsDisabledEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsDisabled, v))
}

// IsDisabledNEQ applies the NEQ predicate on the "is_disabled" field.
func IsDisabledNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsDisabled, v))
}

// DateOfBirthEQ applies the EQ predicate on the "date_of_birth" field.
func DateOfBirthEQ(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDateOfBirth, v))
}

// DateOfBirthNEQ applies the NEQ predicate on the "date_of_birth" field.
func DateOfBirthNEQ(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDateOfBirth, v))
}

// DateOfBirthIn applies the In predicate on the "date_of_birth" field.
func DateOfBirthIn(vs ...*entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldIn(FieldDateOfBirth, vs...))
}

// DateOfBirthNotIn applies the NotIn predicate on the "date_of_birth" field.
func DateOfBirthNotIn(vs ...*entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDateOfBirth, vs...))
}

// DateOfBirthGT applies the GT predicate on the "date_of_birth" field.
func DateOfBirthGT(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldGT(FieldDateOfBirth, v))
}

// DateOfBirthGTE applies the GTE predicate on the "date_of_birth" field.
func DateOfBirthGTE(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDateOfBirth, v))
}

// DateOfBirthLT applies the LT predicate on the "date_of_birth" field.
func DateOfBirthLT(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldLT(FieldDateOfBirth, v))
}

// DateOfBirthLTE applies the LTE predicate on the "date_of_birth" field.
func DateOfBirthLTE(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDateOfBirth, v))
}

// DateOfBirthIsNil applies the IsNil predicate on the "date_of_birth" field.
func DateOfBirthIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDateOfBirth))
}

// DateOfBirthNotNil applies the NotNil predicate on the "date_of_birth" field.
func DateOfBirthNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDateOfBirth))
}

// DateOfJoinEQ applies the EQ predicate on the "date_of_join" field.
func DateOfJoinEQ(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDateOfJoin, v))
}

// DateOfJoinNEQ applies the NEQ predicate on the "date_of_join" field.
func DateOfJoinNEQ(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDateOfJoin, v))
}

// DateOfJoinIn applies the In predicate on the "date_of_join" field.
func DateOfJoinIn(vs ...*entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldIn(FieldDateOfJoin, vs...))
}

// DateOfJoinNotIn applies the NotIn predicate on the "date_of_join" field.
func DateOfJoinNotIn(vs ...*entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDateOfJoin, vs...))
}

// DateOfJoinGT applies the GT predicate on the "date_of_join" field.
func DateOfJoinGT(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldGT(FieldDateOfJoin, v))
}

// DateOfJoinGTE applies the GTE predicate on the "date_of_join" field.
func DateOfJoinGTE(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDateOfJoin, v))
}

// DateOfJoinLT applies the LT predicate on the "date_of_join" field.
func DateOfJoinLT(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldLT(FieldDateOfJoin, v))
}

// DateOfJoinLTE applies the LTE predicate on the "date_of_join" field.
func DateOfJoinLTE(v *entFields.DateOnly) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDateOfJoin, v))
}

// TimingStartEQ applies the EQ predicate on the "timing_start" field.
func TimingStartEQ(v *entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldEQ(FieldTimingStart, v))
}

// TimingStartNEQ applies the NEQ predicate on the "timing_start" field.
func TimingStartNEQ(v *entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldTimingStart, v))
}

// TimingStartIn applies the In predicate on the "timing_start" field.
func TimingStartIn(vs ...*entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldIn(FieldTimingStart, vs...))
}

// TimingStartNotIn applies the NotIn predicate on the "timing_start" field.
func TimingStartNotIn(vs ...*entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldTimingStart, vs...))
}

// TimingStartGT applies the GT predicate on the "timing_start" field.
func TimingStartGT(v *entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldGT(FieldTimingStart, v))
}

// TimingStartGTE applies the GTE predicate on the "timing_start" field.
func TimingStartGTE(v *entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldGTE(FieldTimingStart, v))
}

// TimingStartLT applies the LT predicate on the "timing_start" field.
func TimingStartLT(v *entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldLT(FieldTimingStart, v))
}

// TimingStartLTE applies the LTE predicate on the "timing_start" field.
func TimingStartLTE(v *entFields.TimeOnly) predicate.User {
	return predicate.User(sql.FieldLTE(FieldTimingStart, v))
}

// TimingStartIsNil applies the IsNil predicate on the "timing_start" field.
func TimingStartIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldTimingStart))
}

// TimingStartNotNil applies the NotNil predicate on the "timing_start" field.
func TimingStartNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldTimingStart))
}

// EmployeeIDEQ applies the EQ predicate on the "employee_id" field.
func EmployeeIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmployeeID, v))
}

// EmployeeIDNEQ applies the NEQ predicate on the "employee_id" field.
func EmployeeIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmployeeID, v))
}

// EmployeeIDIn applies the In predicate on the "employee_id" field.
func EmployeeIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmployeeID, vs...))
}

// EmployeeIDNotIn applies the NotIn predicate on the "employee_id" field.
func EmployeeIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmployeeID, vs...))
}

// EmployeeIDGT applies the GT predicate on the "employee_id" field.
func EmployeeIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmployeeID, v))
}

// EmployeeIDGTE applies the GTE predicate on the "employee_id" field.
func EmployeeIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmployeeID, v))
}

// EmployeeIDLT applies the LT predicate on the "employee_id" field.
func EmployeeIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmployeeID, v))
}

// EmployeeIDLTE applies the LTE predicate on the "employee_id" field.
func EmployeeIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmployeeID, v))
}

// EmployeeIDContains applies the Contains predicate on the "employee_id" field.
func EmployeeIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmployeeID, v))
}

// EmployeeIDHasPrefix applies the HasPrefix predicate on the "employee_id" field.
func EmployeeIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmployeeID, v))
}

// EmployeeIDHasSuffix applies the HasSuffix predicate on the "employee_id" field.
func EmployeeIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmployeeID, v))
}

// EmployeeIDIsNil applies the IsNil predicate on the "employee_id" field.
func EmployeeIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldEmployeeID))
}

// EmployeeIDNotNil applies the NotNil predicate on the "employee_id" field.
func EmployeeIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldEmployeeID))
}

// EmployeeIDEqualFold applies the EqualFold predicate on the "employee_id" field.
func EmployeeIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmployeeID, v))
}

// EmployeeIDContainsFold applies the ContainsFold predicate on the "employee_id" field.
func EmployeeIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmployeeID, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDIsNil applies the IsNil predicate on the "role_id" field.
func RoleIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldRoleID))
}

// RoleIDNotNil applies the NotNil predicate on the "role_id" field.
func RoleIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldRoleID))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
