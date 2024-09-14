package schema

import (
	"context"
	"project/ent"
	"project/ent/company"
)

func (Company) Resource() *ent.CompanyResource {
	return &ent.CompanyResource{
		HasReadPermission: func(ctx context.Context, jwtUser *ent.User, client *ent.CompanyQuery) error {
			client.Where(company.ID(jwtUser.CompanyID))
			return nil
		},
		HasCreatePermission: func(ctx context.Context, jwtUser *ent.User, client []*ent.CreateCompanyInput) error {
			return nil
		},
	}
}

func (User) Resource() *ent.UserResource {
	return &ent.UserResource{
		HasReadPermission: func(ctx context.Context, jwtUser *ent.User, client *ent.UserQuery) error {
			return nil
		},
		HasCreatePermission: func(ctx context.Context, jwtUser *ent.User, client []*ent.CreateUserInput) error {
			for _, member := range client {
				member.CompanyID = &jwtUser.CompanyID
			}
			return nil
		},
		HasDeletePermission: func(ctx context.Context, jwtUser *ent.User, client *ent.User) error {
			return nil
		},
		HasUpdatePermission: func(ctx context.Context, jwtUser *ent.User, member *ent.User, update *ent.UpdateUserInput) error {
			return nil
		},
	}
}
