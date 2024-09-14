package resource

import "context"

type Base[
	Model any,
	ModelList any,
	ModelQuery any,
	CreateModelInputs any,
	UpdateModelInput any,
	JwtUser any,
] struct {
	DefaultLimit int

	SkipJWTCheck bool

	HasReadPermission   func(context.Context, JwtUser, ModelQuery) error
	HasCreatePermission func(context.Context, JwtUser, CreateModelInputs) error
	HasUpdatePermission func(context.Context, JwtUser, Model, UpdateModelInput) error
	HasDeletePermission func(context.Context, JwtUser, Model) error

	ExportHeaders []string
	ExportPreload func(context.Context, JwtUser, ModelQuery) error
	ExportRow     func(context.Context, JwtUser, Model) []interface{}

	AfterObjectSave   func(context.Context, ModelList, JwtUser)
	AfterObjectUpdate func(context.Context, Model, JwtUser)
}
