package schema

import (
	"context"
	"errors"
	"project/src/models/ent"
	"project/src/models/ent/company"
	"project/src/models/ent/user"
)

func (Company) Resource() *ent.CompanyResource {
	return &ent.CompanyResource{
		HasReadPermission: func(ctx context.Context, jwtUser *ent.User, client *ent.CompanyQuery) error {
			client.Where(company.ID(jwtUser.CompanyID)) // only allow the company members to read the company
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
			client.Where(user.CompanyID(jwtUser.CompanyID))
			return nil
		},
		HasCreatePermission: func(ctx context.Context, jwtUser *ent.User, client []*ent.CreateUserInput) error {
			for _, member := range client {
				member.CompanyID = &jwtUser.CompanyID
			}
			return nil
		},
		HasDeletePermission: func(ctx context.Context, jwtUser *ent.User, client *ent.User) error {
			return errors.New("you cannot delete users")
		},
		HasUpdatePermission: func(ctx context.Context, jwtUser *ent.User, member *ent.User, update *ent.UpdateUserInput) error {
			if jwtUser.ID == member.ID {
				return nil
			}
			return errors.New("you cannot update other users")
		},
	}
}
