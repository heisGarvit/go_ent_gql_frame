package auth

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
	"project/src/infra/db"
	"project/src/models/ent"
	"project/src/models/ent/user"
	"project/src/utils/jwt"
)

type AuthGqlResolver struct{}

type LoginResponse struct {
	Token string
	User  ent.User
}

func (r *AuthGqlResolver) Login(ctx context.Context, email string, password string) (*LoginResponse, error) {

	user, err := db.Client.User.Query().Where(
		user.EmailEQ(email),
		user.IsDisabled(false),
	).First(ctx)

	if err != nil {
		slog.Debug(err.Error())
		return nil, errors.New("username Or Password Invalid or is Disabled")
	}

	userCompany, err := user.Company(ctx)
	if err != nil {
		return nil, err
	}

	if userCompany.IsDisabled {
		return nil, errors.New("username Or Password Invalid or is Disabled")
	}

	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	//	return nil, errors.New("username Or Password Invalid or is Disabled")
	//}

	_, tokenString, err := jwt.GenerateToken(user.ID.String(), true)

	loginResponse := LoginResponse{
		Token: tokenString,
		User:  *user,
	}

	return &loginResponse, err
}

func (r *AuthGqlResolver) RefreshToken(ctx context.Context) (*LoginResponse, error) {

	userId, _, err := jwt.RefreshTokenValidate(ctx)
	if err != nil {
		return nil, err
	}

	userUUID, err := uuid.Parse(*userId)
	if err != nil {
		return nil, err
	}

	user, err := db.Client.User.Get(ctx, userUUID)
	if err != nil {
		return nil, errors.New("user is disabled")
	}

	userCompany, err := user.Company(ctx)
	if err != nil {
		return nil, err
	}

	if userCompany.IsDisabled {
		return nil, errors.New("company is disabled")
	}

	_, tokenString, err := jwt.GenerateToken(user.ID.String(), false)

	loginResponse := LoginResponse{
		Token: tokenString,
		User:  *user,
	}

	return &loginResponse, err
}
