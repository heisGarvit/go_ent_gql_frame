package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"project/src/infra/env"
	"project/src/utils/tracer"
	"time"
)

var secret = []byte(env.Get("JWT_SECRET"))

func GenerateToken(userId string, fresh bool) (*jwt.Token, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":   uuid.NewString(),
		"sub":   userId,
		"fresh": fresh,
		"exp":   time.Now().Add(time.Hour * time.Duration(10000)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	return token, tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})

	return token, err
}

func RefreshTokenValidate(ctx context.Context) (*string, *jwt.Token, error) {
	jwtTokenString, ok := ctx.Value("Authorization").(string)
	if !ok || jwtTokenString == "" {
		return nil, nil, errors.New("invalid token")
	}

	jwtToken, err := ParseToken(jwtTokenString)
	if err != nil && err.Error() != "token has invalid claims: token is expired" {
		return nil, jwtToken, errors.New("invalid token")
	}

	expirationTime, err := jwtToken.Claims.GetExpirationTime()
	if err != nil {
		return nil, jwtToken, err
	}

	if expirationTime.Time.Add(time.Hour * time.Duration(720)).Before(time.Now()) {
		return nil, jwtToken, errors.New("invalid token expiry")
	}

	userId, err := jwtToken.Claims.GetSubject()
	if err != nil {
		return nil, jwtToken, err
	}
	return &userId, jwtToken, nil
}

func GetCurrentJwtUserID(ctx context.Context) (*string, error) {
	_, span := tracer.Tracer.Start(ctx, "GetCurrentJwtUser")
	defer span.End()

	jwtTokenString, ok := ctx.Value("Authorization").(string)
	if !ok || jwtTokenString == "" {
		return nil, errors.New("invalid token")
	}

	jwtToken, err := ParseToken(jwtTokenString)
	if err != nil {
		return nil, err
	}
	userId, err := jwtToken.Claims.GetSubject()
	if err != nil || userId == "" {
		return nil, err
	}
	return &userId, nil
}
